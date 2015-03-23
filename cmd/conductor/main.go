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

// Set up a mutex to sync all the goroutines
type State struct {
	m     sync.Mutex
	c     *sync.Cond
	begin bool
}

func main() {
	// Parse command line flags
	port := flag.String("p", "8888", "Specifies the port")
	reps := flag.Int("r", 1, "Specifies the number of repetitions")

	flag.Parse()

	// Init the conductor's internal state
	conductor := State{}
	conductor.c = sync.NewCond(&conductor.m)
	conductor.begin = false

	mysong := beepster.CreateSong(2, *reps)

	var wg sync.WaitGroup

	// Create goroutines to compose all the parts to create a track
	wg.Add(1)
	go mysong.CreateTrackFiller(0, wg.Done, &melody, &rest, &rest)()
	wg.Add(1)
	go mysong.CreateTrackFiller(1, wg.Done, &rest, &melody, &tag)()

	// Start listening for client connections
	listener, err := net.Listen("tcp", ":"+(*port))
	if err != nil {
		fmt.Print("Error listening on port " + *port)
		os.Exit(1)
	}

	fmt.Println("We need", mysong.NumTracks, " computers")
	// Only accept mysong.numTracks clients, one for each part
	for i := uint8(0); i < mysong.NumTracks; i++ {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Print("Error accepting connection")
			os.Exit(1)
		}
		fmt.Println("New incomming connection")

		wg.Add(1)
		go handleConnection(conn, &mysong.Track[i], &wg, &conductor)
	}

	fmt.Println("All parts filled: commencing playing")

	// Tell all goroutines to start sending data to their prospective clients
	conductor.m.Lock()
	conductor.begin = true
	conductor.c.Broadcast()
	conductor.m.Unlock()

	// Wait until all goroutines are finished before halting execution
	wg.Wait()
}

// Send data from a given track accros a given connection by encoding it with gob
func handleConnection(conn net.Conn, track *chan beepster.Note, wg *sync.WaitGroup, state *State) {
	defer wg.Done()
	enc := gob.NewEncoder(conn)

	// Wait until all the connections are ready
	state.m.Lock()
	for !state.begin {
		state.c.Wait()
	}
	state.m.Unlock()

	// Send notes over the wire
	for note := range *track {
		enc.Encode(note)
	}
	conn.Close()
	fmt.Println("Connection closed")
}
