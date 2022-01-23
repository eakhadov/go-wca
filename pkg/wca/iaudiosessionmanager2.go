package wca

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IAudioSessionManager2 struct {
	IAudioSessionManager
}

type IAudioSessionManager2Vtbl struct {
	IAudioSessionManagerVtbl
	GetSessionEnumerator          uintptr
	RegisterSessionNotification   uintptr
	UnregisterSessionNotification uintptr
	RegisterDuckNotification      uintptr
	UnregisterDuckNotification    uintptr
}

func (v *IAudioSessionManager2) VTable() *IAudioSessionManager2Vtbl {
	return (*IAudioSessionManager2Vtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IAudioSessionManager2) GetSessionEnumerator(SessionEnum **IAudioSessionEnumerator) (err error) {
	hr, _, _ := syscall.Syscall(
		v.VTable().GetSessionEnumerator,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(SessionEnum)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func (v *IAudioSessionManager2) RegisterSessionNotification(SessionNotification *IAudioSessionNotification) (err error) {
	hr, _, _ := syscall.Syscall(
		v.VTable().RegisterSessionNotification,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(SessionNotification)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func (v *IAudioSessionManager2) UnregisterSessionNotification(SessionNotification *IAudioSessionNotification) (err error) {
	hr, _, _ := syscall.Syscall(
		v.VTable().UnregisterSessionNotification,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(SessionNotification)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func (v *IAudioSessionManager2) RegisterDuckNotification(sessionID *string, duckNotification *IAudioVolumeDuckNotification) (err error) {
	hr, _, _ := syscall.Syscall(
		v.VTable().RegisterDuckNotification,
		3,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(sessionID)),
		uintptr(unsafe.Pointer(duckNotification)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func (v *IAudioSessionManager2) UnregisterDuckNotification(duckNotification *IAudioVolumeDuckNotification) (err error) {
	hr, _, _ := syscall.Syscall(
		v.VTable().UnregisterDuckNotification,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(duckNotification)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
