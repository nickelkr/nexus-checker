# nexus-checker
A simple program to check if the Google Nexus 6 is available.

### Purpose
I was tired of hitting refresh on Wednesdays to see if the Nexus 6 64 GB was
available.

### How to run
First build the project:
`go build`

Then run:
`./nexus-checker`

Options:
- size: 32GB or 64GB
- color: white or blue
- duration: how often to check the page (in seconds)

Without any options provided, the checker defaults:
- size     ~> 64 GB
- color    ~> white
- duration ~> 10 seconds

Example with options that would search for a 32GB Midnight Blue Nexus 6 every
3 seconds:
`./nexus-checker -size 32 -color blue -duration 3`

### OSX and Chrome
If you are on OSX and have Chrome installed it will open the purchase page for
the phone version it was monitoring. This is very basic and if you're on OSX
but don't have Chrome installed it will return an error and die.

### Notes
Originally, the nexus-checker would issue a http.Head() request and utilize the
Content-Length attribute to determine if the page had changed. If the page had
changed, it would issue a http.Get request and then run the contains() on the
received page. This functionality has been removed and also explains the still
exists Page struct.
