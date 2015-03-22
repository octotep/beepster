package main

import (
	"os"
	"syscall"
	"time"
)

const (
	KIOCSOUND       = 0x4B2F
	EVIOCGSND       = 0x4B2F
	CLOCK_TICK_RATE = 1193180
)

type Note struct {
	// Frequency (Hz)
	freq float32
	// Length of note (ms)
	length uint32
	// Delay after note (ms)
	delay uint32
}

type Speaker struct {
	// File for speaker access
	file *os.File
}

// Create a new speaker object ready for use
func speaker_init() *Speaker {
	// Open speaker
	fd, err := os.Create("/dev/tty0")
	if err != nil {
		panic(err)
	}
	speaker := new(Speaker)
	speaker.file = fd
	return speaker
}

// Makes the speaker play a certain tone.
// To stop it, call it with a frequency of 0
func (spkr *Speaker) beep(freq float32) {
	if freq != 0 {
		// Beep at the given pitch
		_, _, _ = syscall.Syscall(syscall.SYS_IOCTL, spkr.file.Fd(), uintptr(KIOCSOUND), uintptr(CLOCK_TICK_RATE/freq))
	} else {
		// Cease the beeping
		_, _, _ = syscall.Syscall(syscall.SYS_IOCTL, spkr.file.Fd(), uintptr(KIOCSOUND), uintptr(0))
	}
}

// Plays an entire Note object
func (spkr *Speaker) play_note(note *Note) {
	spkr.beep(note.freq)
	time.Sleep(time.Duration(note.length) * time.Millisecond)
	spkr.beep(0)
	time.Sleep(time.Duration(note.delay) * time.Millisecond)
}

// Properly shutdowns the speaker
func (spkr *Speaker) shutdown() {
	// Stop all beeping
	spkr.beep(0)
	// Close the speaker
	err := spkr.file.Close()
	if err != nil {
		panic(err)
	}
}

func main() {
	// Open speaker
	speaker := speaker_init()
	// Close speaker when done
	defer speaker.shutdown()

	note := Note{440.0, 1000, 5}
	note2 := Note{880.0, 1000, 5}

	speaker.play_note(&note)
	speaker.play_note(&note2)
}
