// Copyright 2026 Matthew Wilson and Synesis Information Systems. All
// rights reserved. Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux

/*
 * Created: 16th November 2026
 * Updated: 17th November 2026
 */

package woad_test

import (
	"os"
	"strconv"
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
	var number uint32

	err := ioctl(master.Fd(), syscall.TIOCGPTN, uintptr(unsafe.Pointer(&number)))

	if err != nil {
		return "", err
	} else {
		return "/dev/pts/" + strconv.FormatUint(uint64(number), 10), nil
	}
}

func grantAndUnlock(master *os.File) error {
	var unlock int32

	return ioctl(master.Fd(), syscall.TIOCSPTLCK, uintptr(unsafe.Pointer(&unlock)))
}
