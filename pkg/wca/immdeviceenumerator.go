package wca

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IMMDeviceEnumerator struct {
	ole.IUnknown
}

type IMMDeviceEnumeratorVtbl struct {
	ole.IUnknownVtbl
	EnumAudioEndpoints                     uintptr
	GetDefaultAudioEndpoint                uintptr
	GetDevice                              uintptr
	RegisterEndpointNotificationCallback   uintptr
	UnregisterEndpointNotificationCallback uintptr
}

func (v *IMMDeviceEnumerator) VTable() *IMMDeviceEnumeratorVtbl {
	return (*IMMDeviceEnumeratorVtbl)(unsafe.Pointer(v.RawVTable))
}

// func (v *IMMDeviceEnumerator) EnumAudioEndpoints(eDataFlow, stateMask uint32, dc **IMMDeviceCollection) (err error) {
// 	err = mmdeEnumAudioEndpoints(v, eDataFlow, stateMask, dc)
// 	return
// }

func (v *IMMDeviceEnumerator) GetDefaultAudioEndpoint(dataFlow, role uint32, ppEndpoint **IMMDevice) (err error) {
	hr, _, _ := syscall.Syscall6(
		v.VTable().GetDefaultAudioEndpoint,
		4,
		uintptr(unsafe.Pointer(v)),
		uintptr(dataFlow),
		uintptr(role),
		uintptr(unsafe.Pointer(ppEndpoint)),
		0,
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func (v *IMMDeviceEnumerator) GetDevice(pwstrId *uint32, ppDevice **IMMDevice) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

// func (v *IMMDeviceEnumerator) RegisterEndpointNotificationCallback(pClient *IMMNotificationClient) (err error) {
// 	hr, _, _ := syscall.Syscall(
// 		v.VTable().RegisterEndpointNotificationCallback,
// 		2,
// 		uintptr(unsafe.Pointer(v)),
// 		uintptr(unsafe.Pointer(pClient)),
// 		0)
// 	if hr != 0 {
// 		err = ole.NewError(hr)
// 	}
// 	return
// }

// func (v *IMMDeviceEnumerator) UnregisterEndpointNotificationCallback(mmnc *IMMNotificationClient) (err error) {
// 	err = mmdeUnregisterEndpointNotificationCallback(v, mmnc)
// 	return
// }
