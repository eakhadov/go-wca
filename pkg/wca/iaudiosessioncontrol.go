package wca

import (
	"syscall"
	"unsafe"

	"github.com/eakhadov/go-wca/internal/wca"
	"github.com/go-ole/go-ole"
)

type IAudioSessionControl struct {
	ole.IUnknown
}

type IAudioSessionControlVtbl struct {
	ole.IUnknownVtbl
	GetState                           uintptr
	GetDisplayName                     uintptr
	SetDisplayName                     uintptr
	GetIconPath                        uintptr
	SetIconPath                        uintptr
	GetGroupingParam                   uintptr
	SetGroupingParam                   uintptr
	RegisterAudioSessionNotification   uintptr
	UnregisterAudioSessionNotification uintptr
}

func (v *IAudioSessionControl) VTable() *IAudioSessionControlVtbl {
	return (*IAudioSessionControlVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IAudioSessionControl) GetState(pRetVal *uint32) (err error) {
	hr, _, _ := syscall.Syscall(
		v.VTable().GetState,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(pRetVal)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return

}

func (v *IAudioSessionControl) GetDisplayName(retVal *string) (err error) {
	var retValPtr uint64
	hr, _, _ := syscall.Syscall(
		v.VTable().GetDisplayName,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retValPtr)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	*retVal = wca.LPWSTRToString(retValPtr)
	ole.CoTaskMemFree(uintptr(retValPtr))
	return
}

func (v *IAudioSessionControl) SetDisplayName(Value *string, EventContext *ole.GUID) (err error) {
	hr, _, _ := syscall.Syscall(
		v.VTable().SetDisplayName,
		3,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(Value)),
		uintptr(unsafe.Pointer(EventContext)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func (v *IAudioSessionControl) GetIconPath(retVal *string) (err error) {
	var retValPtr uint64
	hr, _, _ := syscall.Syscall(
		v.VTable().GetIconPath,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&retValPtr)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	*retVal = wca.LPWSTRToString(retValPtr)
	ole.CoTaskMemFree(uintptr(retValPtr))
	return
}

func (v *IAudioSessionControl) SetIconPath(Value *string, EventContext *ole.GUID) (err error) {
	hr, _, _ := syscall.Syscall(
		v.VTable().SetIconPath,
		3,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(Value)),
		uintptr(unsafe.Pointer(EventContext)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func (v *IAudioSessionControl) GetGroupingParam(pRetVal *ole.GUID) (err error) {
	hr, _, _ := syscall.Syscall(
		v.VTable().GetGroupingParam,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(pRetVal)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func (v *IAudioSessionControl) SetGroupingParam(Override *ole.GUID, EventContext *ole.GUID) (err error) {
	hr, _, _ := syscall.Syscall(
		v.VTable().SetGroupingParam,
		3,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(Override)),
		uintptr(unsafe.Pointer(EventContext)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func (v *IAudioSessionControl) RegisterAudioSessionNotification(NewNotifications *IAudioSessionEvents) (err error) {
	hr, _, _ := syscall.Syscall(
		v.VTable().RegisterAudioSessionNotification,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(NewNotifications)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func (v *IAudioSessionControl) UnregisterAudioSessionNotification(NewNotifications *IAudioSessionEvents) (err error) {
	hr, _, _ := syscall.Syscall(
		v.VTable().UnregisterAudioSessionNotification,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(NewNotifications)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
