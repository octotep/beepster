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
- [First order functions in golang](https://golang.org/doc/codewalk/functions/)

# TODO
- [X] Learn go
- [X] Play a beep with go
- [X] Play different pitched beeps
- [X] Create library for beeping
- [X] Design test program which uses the library
- [X] Learn goroutines
- [X] Create a Player client
- [X] Create a Conductor server
	- [X] Simple test to play across network connection
	- [X] Sync multiple computers to play together
- [X] Play Row, Row, Row Your Boat in a round
	- Two clients synced with one server
- [X] Polish server and client
	- [X] Handle when the songs over intelligently
	- [X] Move song definition and creation out of main, too much clutter
	- [X] Remove debug output
- [ ] Play ~~midi~~ MusicXML?!
	- [X] Can create two tracks for two computers to play in sync
	- [X] Handle chords without going out of sync
	- [ ] Handle ties correctly (no note delay)
	- [ ] Handle articulaton (change delay)
	- [ ] Handle transposition
	- [ ] Test for robustness
- [ ] General Polish
	- [ ] Error handling in library (pass to actual program)
	- [ ] Document everything, add program usage examples, also limitiations
	- [ ] Let a conductor take less computers than there are parts (ignore the rest)
	- [ ] Handle compressed mxl documents, because XML is very verbose
	- [ ] Make gobeep a better alternative to the beep program, handle all original
