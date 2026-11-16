// Copyright 2026 Matthew Wilson and Synesis Information Systems. All
// rights reserved. Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin

/*
 * Created: 16th November 2026
 * Updated: 17th November 2026
 */

package woad_test

import (
	"os"
	"syscall"
	"unsafe"
)

func ioctl(
	fd uintptr,
	request uintptr,
	arg uintptr,
) error {
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, fd, request, arg)

	if errno != 0 {
		return errno
	} else {
		return nil
	}
}

func openPty() (*os.File, *os.File, error) {
	master, err := os.OpenFile("/dev/ptmx", os.O_RDWR, 0)
	if err != nil {
		return nil, nil, err
	}

	err = grantAndUnlock(master)
	if err != nil {
		master.Close()

		return nil, nil, err
	}

	name, err := ptsname(master)
	if err != nil {
		master.Close()

		return nil, nil, err
	}

	slave, err := os.OpenFile(name, os.O_RDWR, 0)
	if err != nil {
		master.Close()

		return nil, nil, err
	}

	return master, slave, nil
}

func ptsname(master *os.File) (string, error) {
	name := make([]byte, 128)

	err := ioctl(master.Fd(), syscall.TIOCPTYGNAME, uintptr(unsafe.Pointer(&name[0])))
	if err != nil {
		return "", err
	}

	end := 0

	for end < len(name) && name[end] != 0 {
		end++
	}

	return string(name[:end]), nil
}

func grantAndUnlock(master *os.File) error {
	err := ioctl(master.Fd(), syscall.TIOCPTYGRANT, 0)

	if err != nil {
		return err
	} else {
		return ioctl(master.Fd(), syscall.TIOCPTYUNLK, 0)
	}
}
