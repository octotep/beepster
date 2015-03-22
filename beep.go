package beepster

import (
	"os"
	"syscall"
	"time"
)

const (
	KIOCSOUND       = 0x4B2F
	CLOCK_TICK_RATE = 1193180
)

type Note struct {
	Freq float32   // Frequency (Hz)
	Length uint32  // Length of note (ms)
	Delay uint32   // Delay after note (ms)
}

type Speaker struct {
	// File for speaker access
	file *os.File
}

// Create a new speaker object ready for use
func Open() *Speaker {
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
func (spkr *Speaker) Beep(freq float32) {
	if freq != 0 {
		// Beep at the given pitch
		_, _, _ = syscall.Syscall(syscall.SYS_IOCTL, spkr.file.Fd(), uintptr(KIOCSOUND), uintptr(CLOCK_TICK_RATE/freq))
	} else {
		// Cease the beeping
		_, _, _ = syscall.Syscall(syscall.SYS_IOCTL, spkr.file.Fd(), uintptr(KIOCSOUND), uintptr(0))
	}
}

// Plays an entire Note object
func (spkr *Speaker) PlayNote(note *Note) {
	spkr.Beep(note.Freq)
	time.Sleep(time.Duration(note.Length) * time.Millisecond)
	spkr.Beep(0)
	time.Sleep(time.Duration(note.Delay) * time.Millisecond)
}

// Properly shutdowns the speaker
func (spkr *Speaker) Close() {
	// Stop all beeping
	spkr.Beep(0)
	// Close the speaker
	err := spkr.file.Close()
	if err != nil {
		panic(err)
	}
}
