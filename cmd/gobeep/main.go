package main

import (
	"sync"
	"fmt"
	"github.com/octotep/beepster"
)

func main() {
	// Use a wait group to keep track of the goroutines
	var wg sync.WaitGroup

	// Open speaker
	speaker := beepster.Open()
	// Close speaker when done
	defer speaker.Close()

	channel := make(chan beepster.Note, 7)

	// goroutine to play any notes from a channel
	wg.Add(1)
	go func() {
		defer wg.Done()
		for note := range channel {
			speaker.PlayNote(&note)
		}
	}()

	// Create array to store notes in
	track := make([]beepster.Note, 42)
	track[0]  = beepster.Note{261.0,   500, 5}
	track[1]  = beepster.Note{261.0,   500, 5}
	track[2]  = beepster.Note{392.0,   500, 5}
	track[3]  = beepster.Note{392.0,   500, 5}
	track[4]  = beepster.Note{440.0,   500, 5}
	track[5]  = beepster.Note{440.0,   500, 5}
	track[6]  = beepster.Note{392.0,  1000, 5}

	track[7]  = beepster.Note{349.23,  500, 5}
	track[8]  = beepster.Note{349.23,  500, 5}
	track[9]  = beepster.Note{329.63,  500, 5}
	track[10] = beepster.Note{329.63,  500, 5}
	track[11] = beepster.Note{293.67,  500, 5}
	track[12] = beepster.Note{293.67,  500, 5}
	track[13] = beepster.Note{261.63, 1000, 5}

	track[14] = beepster.Note{392.0,   500, 5}
	track[15] = beepster.Note{392.0,   500, 5}
	track[16] = beepster.Note{349.23,  500, 5}
	track[17] = beepster.Note{349.23,  500, 5}
	track[18] = beepster.Note{329.63,  500, 5}
	track[19] = beepster.Note{329.63,  500, 5}
	track[20] = beepster.Note{293.67, 1000, 5}

	track[21] = beepster.Note{392.0,   500, 5}
	track[22] = beepster.Note{392.0,   500, 5}
	track[23] = beepster.Note{349.23,  500, 5}
	track[24] = beepster.Note{349.23,  500, 5}
	track[25] = beepster.Note{329.63,  500, 5}
	track[26] = beepster.Note{329.63,  500, 5}
	track[27] = beepster.Note{293.67, 1000, 5}

	track[28] = beepster.Note{261.0,   500, 5}
	track[29] = beepster.Note{261.0,   500, 5}
	track[30] = beepster.Note{392.0,   500, 5}
	track[31] = beepster.Note{392.0,   500, 5}
	track[32] = beepster.Note{440.0,   500, 5}
	track[33] = beepster.Note{440.0,   500, 5}
	track[34] = beepster.Note{392.0,  1000, 5}

	track[35] = beepster.Note{349.23,  500, 5}
	track[36] = beepster.Note{349.23,  500, 5}
	track[37] = beepster.Note{329.63,  500, 5}
	track[38] = beepster.Note{329.63,  500, 5}
	track[39] = beepster.Note{293.67,  500, 5}
	track[40] = beepster.Note{293.67,  500, 5}
	track[41] = beepster.Note{261.63, 1000, 5}

	// goroutine to take notes from array and put them into 
	// the channel
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i, val := range track {
			fmt.Print(i, " ")
			channel <- val
		}
		close(channel)
	}()

	// Wait until all the goroutines are finished before terminating
	wg.Wait()
}
