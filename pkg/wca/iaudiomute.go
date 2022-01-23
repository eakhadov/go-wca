package wca

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IAudioMute struct {
	ole.IUnknown
}

type IAudioMuteVtbl struct {
	ole.IUnknownVtbl
	SetMute uintptr
	GetMute uintptr
}

func (v *IAudioMute) VTable() *IAudioMuteVtbl {
	return (*IAudioMuteVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IAudioMute) SetMute(bMuted bool, pguidEventContext *ole.GUID) (err error) {
	var muteValue uint32
	if bMuted {
		muteValue = 1
	}
	hr, _, _ := syscall.Syscall(
		v.VTable().GetMute,
		3,
		uintptr(unsafe.Pointer(v)),
		uintptr(muteValue),
		uintptr(unsafe.Pointer(pguidEventContext)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func (v *IAudioMute) GetMute(pbMuted *bool) (err error) {
	hr, _, _ := syscall.Syscall(
		v.VTable().GetMute,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(pbMuted)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
