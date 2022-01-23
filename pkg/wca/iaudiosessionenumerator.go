package wca

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IAudioSessionEnumerator struct {
	ole.IUnknown
}

type IAudioSessionEnumeratorVtbl struct {
	ole.IUnknownVtbl
	GetCount   uintptr
	GetSession uintptr
}

func (v *IAudioSessionEnumerator) VTable() *IAudioSessionEnumeratorVtbl {
	return (*IAudioSessionEnumeratorVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IAudioSessionEnumerator) GetCount(SessionCount *int) (err error) {
	hr, _, _ := syscall.Syscall(
		v.VTable().GetCount,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(SessionCount)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return

}

func (v *IAudioSessionEnumerator) GetSession(SessionCount int, Session **IAudioSessionControl) (err error) {
	hr, _, _ := syscall.Syscall(
		v.VTable().GetSession,
		3,
		uintptr(unsafe.Pointer(v)),
		uintptr(SessionCount),
		uintptr(unsafe.Pointer(Session)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
