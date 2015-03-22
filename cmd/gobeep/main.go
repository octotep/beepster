package main

import (
	"github.com/octotep/beepster"
)

func main() {
	// Open speaker
	speaker := beepster.Open()
	// Close speaker when done
	defer speaker.Close()

	note := beepster.Note{440.0, 1000, 5}
	note2 := beepster.Note{880.0, 1000, 5}

	speaker.PlayNote(&note)
	speaker.PlayNote(&note2)
}
