package wca

import (
	"math"
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

type ISimpleAudioVolume struct {
	ole.IUnknown
}

type ISimpleAudioVolumeVtbl struct {
	ole.IUnknownVtbl
	SetMasterVolume uintptr
	GetMasterVolume uintptr
	SetMute         uintptr
	GetMute         uintptr
}

func (v *ISimpleAudioVolume) VTable() *ISimpleAudioVolumeVtbl {
	return (*ISimpleAudioVolumeVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *ISimpleAudioVolume) SetMasterVolume(fLevel float32, EventContext *ole.GUID) (err error) {
	levelValue := math.Float32bits(fLevel)
	hr, _, _ := syscall.Syscall(
		v.VTable().SetMasterVolume,
		3,
		uintptr(unsafe.Pointer(v)),
		uintptr(levelValue),
		uintptr(unsafe.Pointer(EventContext)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func (v *ISimpleAudioVolume) GetMasterVolume(pfLevel *float32) (err error) {
	hr, _, _ := syscall.Syscall(
		v.VTable().GetMasterVolume,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(pfLevel)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func (v *ISimpleAudioVolume) SetMute(bMute bool, EventContext *ole.GUID) (err error) {
	var muteValue uint32
	if bMute {
		muteValue = 1
	}
	hr, _, _ := syscall.Syscall(
		v.VTable().SetMute,
		3,
		uintptr(unsafe.Pointer(v)),
		uintptr(muteValue),
		uintptr(unsafe.Pointer(EventContext)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func (v *ISimpleAudioVolume) GetMute(pbMute *bool) (err error) {
	hr, _, _ := syscall.Syscall(
		v.VTable().GetMute,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(pbMute)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
