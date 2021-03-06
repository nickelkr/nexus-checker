package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"time"
	"runtime"
)

var (
	longColor string
	size      = flag.Int("size", 64, "The storage capacity")
	color     = flag.String("color", "white", "The color of phone")
	duration  = flag.Int("duration", 10, "The sleep time between checks")
)
/* Struct originally created to utilize receiving just the headers and checking
the byte length to see if the page had changed. I don't know if this was
actually worth the trouble. */
type Page struct {
	Url        string
}

// Contains checks to see if a given string exists in a webpage.
func (p *Page) contains(exp string) bool {
	regex := regexp.MustCompile(exp)
	if regex.Find(p.request()) == nil {
		return false
	} else {
		return true
	}
}

// Request issues a GET request and returns the body of the response.
func (p *Page) request() []byte {
	resp, err := http.Get(p.Url)
	defer resp.Body.Close()
	if err != nil {
		die(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		die(err)
	}
	return body
}

// CheckParams looks at the provided flags and sets appropriately.
func checkParams() {
	if *color == "white" {
		longColor = "Cloud_White"
	} else {
		*color = "blue"
		longColor = "Midnight_Blue"
	}
	if *size != 64 && *size != 32 {
		*size = 64
	}
}

// Die prints the error to STDERR and exits with '1' status code.
func die(err error) {
	fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	os.Exit(1)
}

func main() {
	flag.Parse()
	checkParams()
	page := &Page{
		Url: fmt.Sprintf("https://play.google.com/store/devices/details/Nexus_6_%dGB_%s?id=nexus_6_%s_%dgb", *size, longColor, *color, *size),
	}
	for {
		if page.contains("We are out of inventory") {
			fmt.Fprintf(os.Stderr, "Out of stock, still...\n")
		} else {
			fmt.Println("In stock, go get nexus 6!!!")
			if runtime.GOOS == "darwin" {
				chrome := exec.Command(`/Applications/Google Chrome.app/Contents/MacOS/Google Chrome`, page.Url)
				err := chrome.Run()
				if err != nil {
					die(err)
				}
			}
		}

		time.Sleep(time.Duration(*duration) * time.Second)
	}
}
