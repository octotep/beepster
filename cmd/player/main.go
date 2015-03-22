package main

import (
	"encoding/gob"
	"fmt"
	"github.com/octotep/beepster"
	"net"
	"os"
	"sync"
)

const (
	HOST = "localhost"
	PORT = "8888"
)

func main() {
	// Create a channel to send notes to the player goroutine
	track := make(chan *beepster.Note, 10)
	defer close(track)

	speaker := beepster.Open()
	defer speaker.Close()

	// Create a wait group to halt execution of main until both goroutines are finished
	var wg sync.WaitGroup

	// Open a connection to the server
	connection, err := net.Dial("tcp", HOST+":"+PORT)
	if err != nil {
		fmt.Println("Error connecting: ", err.Error())
		os.Exit(1)
	}
	// Close the connection when we reach the end of main
	defer connection.Close()
	fmt.Println("Connected to: " + HOST + ":" + PORT)

	// Create a new decoder for receiving notes
	decoder := gob.NewDecoder(connection)

	wg.Add(1)
	go parseStream(decoder, track, wg)
	wg.Add(1)
	go playStream(speaker, track, wg)

	wg.Wait()
	speaker.Close()
}

func parseStream(dec *gob.Decoder, output chan *beepster.Note, wg sync.WaitGroup) {
	defer wg.Done()
	note := &beepster.Note{}
	for err := dec.Decode(note); err == nil; {
		output <- note
	}
	close(output)
}

func playStream(spkr *beepster.Speaker, output chan *beepster.Note, wg sync.WaitGroup) {
	defer wg.Done()
	for note := range output {
		spkr.PlayNote(note)
	}
}
