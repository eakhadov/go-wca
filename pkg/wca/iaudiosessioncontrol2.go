package wca

import (
	"syscall"
	"unsafe"

	"github.com/eakhadov/go-wca/internal/wca"
	"github.com/go-ole/go-ole"
)

type IAudioSessionControl2 struct {
	IAudioSessionControl
}

type IAudioSessionControl2Vtbl struct {
	IAudioSessionControlVtbl
	GetSessionIdentifier         uintptr
	GetSessionInstanceIdentifier uintptr
	GetProcessId                 uintptr
	IsSystemSoundsSession        uintptr
	SetDuckingPreference         uintptr
}

func (v *IAudioSessionControl2) VTable() *IAudioSessionControl2Vtbl {
	return (*IAudioSessionControl2Vtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IAudioSessionControl2) GetSessionIdentifier(retVal *string) (err error) {
	var retValPtr uint64
	hr, _, _ := syscall.Syscall(
		v.VTable().GetSessionIdentifier,
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

func (v *IAudioSessionControl2) GetSessionInstanceIdentifier(retVal *string) (err error) {
	var retValPtr uint64
	hr, _, _ := syscall.Syscall(
		v.VTable().GetSessionInstanceIdentifier,
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

func (v *IAudioSessionControl2) GetProcessId(pRetVal *uint32) (err error) {
	hr, _, _ := syscall.Syscall(
		v.VTable().GetProcessId,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(pRetVal)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func (v *IAudioSessionControl2) IsSystemSoundsSession() (err error) {
	hr, _, _ := syscall.Syscall(
		v.VTable().IsSystemSoundsSession,
		1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func (v *IAudioSessionControl2) SetDuckingPreference(optOut bool) (err error) {
	var optOutValue uint32
	if optOut {
		optOutValue = 1
	}
	hr, _, _ := syscall.Syscall(
		v.VTable().SetDuckingPreference,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(optOutValue),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
