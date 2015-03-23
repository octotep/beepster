package main

import (
	"encoding/gob"
	"flag"
	"fmt"
	"github.com/octotep/beepster"
	"net"
	"os"
	"sync"
)

type Song struct {
	numTracks uint8
	track     []chan *beepster.Note
}

// Set up a mutex to sync all the goroutines
type State struct {
	m     sync.Mutex
	c     *sync.Cond
	begin bool
}

func main() {
	// Parse command line flags
	port := flag.String("p", "8888", "Specifies the port")

	flag.Parse()

	// Init the conductor's internal state
	conductor := State{}
	conductor.c = sync.NewCond(&conductor.m)
	conductor.begin = false

	// Create the song and fill out all the information we know about it
	mysong := new(Song)
	mysong.numTracks = 2
	mysong.track = make([]chan *beepster.Note, mysong.numTracks)
	for i := range mysong.track {
		mysong.track[i] = make(chan *beepster.Note, 14)
	}

	var wg sync.WaitGroup

	// Add the notes to the first track as required
	wg.Add(1)
	go func() {
		defer wg.Done()
		mysong.track[0] <- &beepster.Note{261.0, 500, 5}
		mysong.track[0] <- &beepster.Note{261.0, 500, 5}
		mysong.track[0] <- &beepster.Note{261.0, 333, 5}
		mysong.track[0] <- &beepster.Note{293.0, 166, 5}
		mysong.track[0] <- &beepster.Note{329.0, 500, 5}

		mysong.track[0] <- &beepster.Note{329.0, 333, 5}
		mysong.track[0] <- &beepster.Note{293.0, 166, 5}
		mysong.track[0] <- &beepster.Note{329.0, 333, 5}
		mysong.track[0] <- &beepster.Note{349.0, 166, 5}
		mysong.track[0] <- &beepster.Note{392.0, 1000, 5}

		mysong.track[0] <- &beepster.Note{523.0, 166, 5}
		mysong.track[0] <- &beepster.Note{523.0, 166, 5}
		mysong.track[0] <- &beepster.Note{523.0, 166, 5}
		mysong.track[0] <- &beepster.Note{392.0, 166, 5}
		mysong.track[0] <- &beepster.Note{392.0, 166, 5}
		mysong.track[0] <- &beepster.Note{392.0, 166, 5}
		mysong.track[0] <- &beepster.Note{329.0, 166, 5}
		mysong.track[0] <- &beepster.Note{329.0, 166, 5}
		mysong.track[0] <- &beepster.Note{329.0, 166, 5}
		mysong.track[0] <- &beepster.Note{261.0, 166, 5}
		mysong.track[0] <- &beepster.Note{261.0, 166, 5}
		mysong.track[0] <- &beepster.Note{261.0, 166, 5}

		mysong.track[0] <- &beepster.Note{392.0, 333, 5}
		mysong.track[0] <- &beepster.Note{349.0, 166, 5}
		mysong.track[0] <- &beepster.Note{329.0, 333, 5}
		mysong.track[0] <- &beepster.Note{293.0, 166, 5}
		mysong.track[0] <- &beepster.Note{261.0, 1000, 5}

		// second time
		mysong.track[0] <- &beepster.Note{261.0, 500, 5}
		mysong.track[0] <- &beepster.Note{261.0, 500, 5}
		mysong.track[0] <- &beepster.Note{261.0, 333, 5}
		mysong.track[0] <- &beepster.Note{293.0, 166, 5}
		mysong.track[0] <- &beepster.Note{329.0, 500, 5}

		mysong.track[0] <- &beepster.Note{329.0, 333, 5}
		mysong.track[0] <- &beepster.Note{293.0, 166, 5}
		mysong.track[0] <- &beepster.Note{329.0, 333, 5}
		mysong.track[0] <- &beepster.Note{349.0, 166, 5}
		mysong.track[0] <- &beepster.Note{392.0, 1000, 5}

		mysong.track[0] <- &beepster.Note{523.0, 166, 5}
		mysong.track[0] <- &beepster.Note{523.0, 166, 5}
		mysong.track[0] <- &beepster.Note{523.0, 166, 5}
		mysong.track[0] <- &beepster.Note{392.0, 166, 5}
		mysong.track[0] <- &beepster.Note{392.0, 166, 5}
		mysong.track[0] <- &beepster.Note{392.0, 166, 5}
		mysong.track[0] <- &beepster.Note{329.0, 166, 5}
		mysong.track[0] <- &beepster.Note{329.0, 166, 5}
		mysong.track[0] <- &beepster.Note{329.0, 166, 5}
		mysong.track[0] <- &beepster.Note{261.0, 166, 5}
		mysong.track[0] <- &beepster.Note{261.0, 166, 5}
		mysong.track[0] <- &beepster.Note{261.0, 166, 5}

		mysong.track[0] <- &beepster.Note{392.0, 333, 5}
		mysong.track[0] <- &beepster.Note{349.0, 166, 5}
		mysong.track[0] <- &beepster.Note{329.0, 333, 5}
		mysong.track[0] <- &beepster.Note{293.0, 166, 5}
		mysong.track[0] <- &beepster.Note{261.0, 1000, 5}
		close(mysong.track[0])
	}()

	// Add the notes to the second track as required
	wg.Add(1)
	go func() {
		defer wg.Done()
		mysong.track[1] <- &beepster.Note{0, 500, 5}
		mysong.track[1] <- &beepster.Note{0, 500, 5}
		mysong.track[1] <- &beepster.Note{0, 333, 5}
		mysong.track[1] <- &beepster.Note{0, 166, 5}
		mysong.track[1] <- &beepster.Note{0, 500, 5}

		mysong.track[1] <- &beepster.Note{261.0, 500, 5}
		mysong.track[1] <- &beepster.Note{261.0, 500, 5}
		mysong.track[1] <- &beepster.Note{261.0, 333, 5}
		mysong.track[1] <- &beepster.Note{293.0, 166, 5}
		mysong.track[1] <- &beepster.Note{329.0, 500, 5}

		mysong.track[1] <- &beepster.Note{329.0, 333, 5}
		mysong.track[1] <- &beepster.Note{293.0, 166, 5}
		mysong.track[1] <- &beepster.Note{329.0, 333, 5}
		mysong.track[1] <- &beepster.Note{349.0, 166, 5}
		mysong.track[1] <- &beepster.Note{392.0, 1000, 5}

		mysong.track[1] <- &beepster.Note{523.0, 166, 5}
		mysong.track[1] <- &beepster.Note{523.0, 166, 5}
		mysong.track[1] <- &beepster.Note{523.0, 166, 5}
		mysong.track[1] <- &beepster.Note{392.0, 166, 5}
		mysong.track[1] <- &beepster.Note{392.0, 166, 5}
		mysong.track[1] <- &beepster.Note{392.0, 166, 5}
		mysong.track[1] <- &beepster.Note{329.0, 166, 5}
		mysong.track[1] <- &beepster.Note{329.0, 166, 5}
		mysong.track[1] <- &beepster.Note{329.0, 166, 5}
		mysong.track[1] <- &beepster.Note{261.0, 166, 5}
		mysong.track[1] <- &beepster.Note{261.0, 166, 5}
		mysong.track[1] <- &beepster.Note{261.0, 166, 5}

		mysong.track[1] <- &beepster.Note{392.0, 333, 5}
		mysong.track[1] <- &beepster.Note{349.0, 166, 5}
		mysong.track[1] <- &beepster.Note{329.0, 333, 5}
		mysong.track[1] <- &beepster.Note{293.0, 166, 5}
		mysong.track[1] <- &beepster.Note{261.0, 1000, 5}

		// Second time
		mysong.track[1] <- &beepster.Note{261.0, 500, 5}
		mysong.track[1] <- &beepster.Note{261.0, 500, 5}
		mysong.track[1] <- &beepster.Note{261.0, 333, 5}
		mysong.track[1] <- &beepster.Note{293.0, 166, 5}
		mysong.track[1] <- &beepster.Note{329.0, 500, 5}

		mysong.track[1] <- &beepster.Note{329.0, 333, 5}
		mysong.track[1] <- &beepster.Note{293.0, 166, 5}
		mysong.track[1] <- &beepster.Note{329.0, 333, 5}
		mysong.track[1] <- &beepster.Note{349.0, 166, 5}
		mysong.track[1] <- &beepster.Note{392.0, 1000, 5}

		mysong.track[1] <- &beepster.Note{523.0, 166, 5}
		mysong.track[1] <- &beepster.Note{523.0, 166, 5}
		mysong.track[1] <- &beepster.Note{523.0, 166, 5}
		mysong.track[1] <- &beepster.Note{392.0, 166, 5}
		mysong.track[1] <- &beepster.Note{392.0, 166, 5}
		mysong.track[1] <- &beepster.Note{392.0, 166, 5}
		mysong.track[1] <- &beepster.Note{329.0, 166, 5}
		mysong.track[1] <- &beepster.Note{329.0, 166, 5}
		mysong.track[1] <- &beepster.Note{329.0, 166, 5}
		mysong.track[1] <- &beepster.Note{261.0, 166, 5}
		mysong.track[1] <- &beepster.Note{261.0, 166, 5}
		mysong.track[1] <- &beepster.Note{261.0, 166, 5}

		mysong.track[1] <- &beepster.Note{392.0, 333, 5}
		mysong.track[1] <- &beepster.Note{349.0, 166, 5}
		mysong.track[1] <- &beepster.Note{329.0, 333, 5}
		mysong.track[1] <- &beepster.Note{293.0, 166, 5}
		mysong.track[1] <- &beepster.Note{261.0, 1000, 5}

		mysong.track[1] <- &beepster.Note{392.0, 333, 5}
		mysong.track[1] <- &beepster.Note{349.0, 166, 5}
		mysong.track[1] <- &beepster.Note{329.0, 333, 5}
		mysong.track[1] <- &beepster.Note{293.0, 166, 5}
		mysong.track[1] <- &beepster.Note{261.0, 1000, 5}
		close(mysong.track[1])
	}()

	listener, err := net.Listen("tcp", ":"+(*port))
	if err != nil {
		fmt.Print("Error listening on port " + *port)
		os.Exit(1)
	}

	for i := uint8(0); i < mysong.numTracks; i++ {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Print("Error accepting connection")
			os.Exit(1)
		}
		wg.Add(1)
		fmt.Println("New incomming connection")
		go handleConnection(conn, mysong.track[i], wg, &conductor)
	}
	fmt.Println("All parts filled: commencing playing")
	conductor.begin = true
	conductor.m.Lock()
	fmt.Println("Main locking it down")
	conductor.c.Broadcast()
	conductor.m.Unlock()

	wg.Wait()
}

func handleConnection(conn net.Conn, track chan *beepster.Note, wg sync.WaitGroup, state *State) {
	defer wg.Done()
	enc := gob.NewEncoder(conn)
	state.m.Lock()
	for !state.begin {
		state.c.Wait()
	}
	state.m.Unlock()
	for note := range track {
		fmt.Println("Encoded another note")
		enc.Encode(note)
	}
	conn.Close()
	fmt.Println("Connection closed")
}
