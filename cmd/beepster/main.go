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

func main() {
	// Parse command line flags
	host := flag.String("h", "localhost", "Specifies the hostname")
	port := flag.String("p", "8888", "Specifies the port")

	flag.Parse()

	// Create a channel to send notes to the player goroutine
	track := make(chan *beepster.Note, 10)
	defer close(track)

	speaker := beepster.Open()
	defer speaker.Close()

	// Create a wait group to halt execution of main until both goroutines are finished
	var wg sync.WaitGroup

	// Open a connection to the server
	connection, err := net.Dial("tcp", (*host)+":"+(*port))
	if err != nil {
		fmt.Println("Error connecting: ", err.Error())
		os.Exit(1)
	}
	// Close the connection when we reach the end of main
	defer connection.Close()
	fmt.Println("Connected to: " + *host + ":" + *port)

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
	// TODO: fix this. Does not advance the note stream at all
	for {
		note := &beepster.Note{}
		err := dec.Decode(note)
		if err != nil {
			break
		}
		fmt.Println("adding note")
		output <- note
	}
	close(output)
}

func playStream(spkr *beepster.Speaker, output chan *beepster.Note, wg sync.WaitGroup) {
	defer wg.Done()
	for note := range output {
		fmt.Println("playing note")
		spkr.PlayNote(note)
	}
}
