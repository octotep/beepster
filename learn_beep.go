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

func main() {
	// Open terminal
	fd, openerr := os.Create("/dev/tty0")
	if openerr != nil {
		panic(openerr)
	}
	// Close fd when done
	defer func() {
		if err := fd.Close(); err != nil {
			panic(err)
		}
	}()
	// Set the pitch
	_, _, _ = syscall.Syscall(syscall.SYS_IOCTL, fd.Fd(), KIOCSOUND, CLOCK_TICK_RATE/440)
	bel := []byte{7}
	_, _ = fd.Write(bel)
	time.Sleep(2 * time.Second)
	_, _, _ = syscall.Syscall(syscall.SYS_IOCTL, fd.Fd(), KIOCSOUND, 0)
}
