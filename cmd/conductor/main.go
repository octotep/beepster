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
	port := flag.String("p", "8888", "Specifies the port")

	flag.Parse()

	var wg sync.WaitGroup

	track := make(chan *beepster.Note, 14)
	wg.Add(1)
	go func() {
		defer wg.Done()
		track <- &beepster.Note{261.0, 500, 5}
		track <- &beepster.Note{261.0, 500, 5}
		track <- &beepster.Note{392.0, 500, 5}
		track <- &beepster.Note{392.0, 500, 5}
		track <- &beepster.Note{440.0, 500, 5}
		track <- &beepster.Note{440.0, 500, 5}
		track <- &beepster.Note{392.0, 1000, 5}

		track <- &beepster.Note{349.23, 500, 5}
		track <- &beepster.Note{349.23, 500, 5}
		track <- &beepster.Note{329.63, 500, 5}
		track <- &beepster.Note{329.63, 500, 5}
		track <- &beepster.Note{293.67, 500, 5}
		track <- &beepster.Note{293.67, 500, 5}
		track <- &beepster.Note{261.63, 1000, 5}

		track <- &beepster.Note{392.0, 500, 5}
		track <- &beepster.Note{392.0, 500, 5}
		track <- &beepster.Note{349.23, 500, 5}
		track <- &beepster.Note{349.23, 500, 5}
		track <- &beepster.Note{329.63, 500, 5}
		track <- &beepster.Note{329.63, 500, 5}
		track <- &beepster.Note{293.67, 1000, 5}

		track <- &beepster.Note{392.0, 500, 5}
		track <- &beepster.Note{392.0, 500, 5}
		track <- &beepster.Note{349.23, 500, 5}
		track <- &beepster.Note{349.23, 500, 5}
		track <- &beepster.Note{329.63, 500, 5}
		track <- &beepster.Note{329.63, 500, 5}
		track <- &beepster.Note{293.67, 1000, 5}

		track <- &beepster.Note{261.0, 500, 5}
		track <- &beepster.Note{261.0, 500, 5}
		track <- &beepster.Note{392.0, 500, 5}
		track <- &beepster.Note{392.0, 500, 5}
		track <- &beepster.Note{440.0, 500, 5}
		track <- &beepster.Note{440.0, 500, 5}
		track <- &beepster.Note{392.0, 1000, 5}

		track <- &beepster.Note{349.23, 500, 5}
		track <- &beepster.Note{349.23, 500, 5}
		track <- &beepster.Note{329.63, 500, 5}
		track <- &beepster.Note{329.63, 500, 5}
		track <- &beepster.Note{293.67, 500, 5}
		track <- &beepster.Note{293.67, 500, 5}
		track <- &beepster.Note{261.63, 1000, 5}
		close(track)
	}()

	listener, err := net.Listen("tcp", ":"+(*port))
	if err != nil {
		fmt.Print("Error listening on port " + *port)
		os.Exit(1)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Print("Error accepting connection")
			os.Exit(1)
		}
		wg.Add(1)
		fmt.Println("New incomming connection")
		go handleConnection(conn, track, wg)
	}
	wg.Wait()
}

func handleConnection(conn net.Conn, track chan *beepster.Note, wg sync.WaitGroup) {
	defer wg.Done()
	enc := gob.NewEncoder(conn)
	for note := range track {
		fmt.Println("Encoded another note")
		enc.Encode(note)
	}
	conn.Close()
	fmt.Println("Connection closed")
}
