package wca

import (
	"syscall"
	"unsafe"
)

func LPWSTRToString(pointer uint64) string {
	var us []uint16
	var i uint32

	var start = unsafe.Pointer(uintptr(pointer))
	for {
		u := *(*uint16)(unsafe.Pointer(uintptr(start) + 2*uintptr(i)))
		if u == 0 {
			break
		}
		us = append(us, u)
		i++
	}

	return syscall.UTF16ToString(us)
}
