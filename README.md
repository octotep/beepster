# beepster

This is my final project for Parallel and Distributed computing.

It plays music using multiple networked computers and their motherboard
speakers. My first (and hopefully only attempt) at this will be in golang.

This will only work on linux. It uses kernel level calls to get at the 
speaker.

# Resources
- [Original beep program](https://github.com/johnath/beep/blob/master/beep.c)
- [ioctl implementation in go](https://github.com/edsrzf/fineline/blob/master/ioctl.go)
- [Programming the internal speaker](http://www.tldp.org/LDP/lpg/node83.html)
- [https://tour.golang.org/](Tour of golang)

# TODO
- [X] Learn go
- [X] Play a beep with go
- [X] Play different pitched beeps
- [X] Create library for beeping
- [ ] Create a Player client
- [ ] Create a Conductor server
- [ ] Play midi?!
