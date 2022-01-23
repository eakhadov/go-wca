package wca

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IAudioSessionManager struct {
	ole.IUnknown
}

type IAudioSessionManagerVtbl struct {
	ole.IUnknownVtbl
	GetAudioSessionControl uintptr
	GetSimpleAudioVolume   uintptr
}

func (v *IAudioSessionManager) VTable() *IAudioSessionManagerVtbl {
	return (*IAudioSessionManagerVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IAudioSessionManager) GetAudioSessionControl(AudioSessionGuid *ole.GUID, StreamFlags uint32, SessionControl **IAudioSessionControl) (err error) {
	hr, _, _ := syscall.Syscall6(
		v.VTable().GetAudioSessionControl,
		4,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(AudioSessionGuid)),
		uintptr(StreamFlags),
		uintptr(unsafe.Pointer(SessionControl)),
		0,
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func (v *IAudioSessionManager) GetSimpleAudioVolume(AudioSessionGuid *ole.GUID, StreamFlags uint32, AudioVolume **ISimpleAudioVolume) (err error) {
	hr, _, _ := syscall.Syscall6(
		v.VTable().GetSimpleAudioVolume,
		4,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(AudioSessionGuid)),
		uintptr(StreamFlags),
		uintptr(unsafe.Pointer(AudioVolume)),
		0,
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
