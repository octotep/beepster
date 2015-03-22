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
	// Set the pitch and beep
	_, _, _ = syscall.Syscall(syscall.SYS_IOCTL, fd.Fd(), uintptr(KIOCSOUND), uintptr(CLOCK_TICK_RATE/440))
	time.Sleep(1 * time.Second)
	_, _, _ = syscall.Syscall(syscall.SYS_IOCTL, fd.Fd(), uintptr(KIOCSOUND), uintptr(CLOCK_TICK_RATE/880))
	time.Sleep(1 * time.Second)
	// Stop beeping
	_, _, _ = syscall.Syscall(syscall.SYS_IOCTL, fd.Fd(), KIOCSOUND, 0)
}
