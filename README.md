# beepster

This is my final project for Parallel and Distributed computing.

It plays music using multiple networked computers and their motherboard
speakers. My first (and hopefully only attempt) at this will be in golang.

This will only work on linux. It uses kernel level calls to get at the 
speaker.

## Programs

In the `cmd` directory, there are three programs.

1. **gobeep** - this is a simple beeper program to make sure I know how to use
the library
2. **conductor** - responsible for gathering players and feeding them notes
3. **beepster** - receive note commands from the conductor and plays them

# Resources
- [Original beep program](https://github.com/johnath/beep/blob/master/beep.c)
- [ioctl implementation in go](https://github.com/edsrzf/fineline/blob/master/ioctl.go)
- [Programming the internal speaker](http://www.tldp.org/LDP/lpg/node83.html)
- [Tour of golang](https://tour.golang.org/)
- [Simple chat server](http://www.badgerr.co.uk/2011/06/20/golang-away-tcp-chat-server/)
- [Read gob data from a network connection](http://stackoverflow.com/a/11202252)

# TODO
- [X] Learn go
- [X] Play a beep with go
- [X] Play different pitched beeps
- [X] Create library for beeping
- [X] Design test program which uses the library
- [X] Learn goroutines
- [X] Create a Player client
- [ ] Create a Conductor server
	- [X] Simple test to play across network connection
	- [ ] Sync multiple computers to play together
- [ ] Play Row, Row, Row Your Boat in a round
	- Two clients synced with one server
- [ ] Play midi?!
