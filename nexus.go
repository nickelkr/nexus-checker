package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"time"
)

var (
	longColor string
	size      = flag.Int("size", 64, "The storage capacity")
	color     = flag.String("color", "white", "The color of phone")
	duration  = flag.Int("duration", 10, "The sleep time between checks")
)

type Page struct {
	Url        string
	LastLength string
}

func (p *Page) changed() (bool, error) {
	resp, err := http.Head(p.Url)
	if err != nil {
		return false, err
	}

	newLength := resp.Header["Content-Length"][0]

	if p.LastLength != newLength {
		p.LastLength = resp.Header["Content-Length"][0]
		return true, nil
	} else {
		return false, nil
	}
}

func (p *Page) update(length string) {
	p.LastLength = length
}

func (p *Page) request() ([]byte, error) {
	resp, err := http.Get(p.Url)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if result, exists := resp.Header["Content-Length"]; exists {
		p.update(result[0])
	}
	return ioutil.ReadAll(resp.Body)
}

func check(page []byte) bool {
	regex := regexp.MustCompile("out of inventory")
	unavailable := regex.Find(page)
	if unavailable == nil {
		return true
	} else {
		return false
	}
}

func checkParams() {
	if *color == "white" {
		longColor = "Cloud_White"
	} else {
		*color = "Midnight_Blue"
		longColor = "blue"
	}
	if *size != 64 || *size != 32 {
		*size = 64
	}
}

func die(error) {
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
		changed, err := page.changed()
		if err != nil {
			die(err)
		}
		if changed {
			body, err := page.request()
			if err != nil {
				die(err)
			}
			if check(body) {
				fmt.Println("In stock!")
			} else {
				fmt.Println("Out of stock... still...")
			}
		}
		time.Sleep(time.Duration(*duration) * time.Second)
	}
}
