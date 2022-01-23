package wca

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IAudioCaptureClient struct {
	ole.IUnknown
}

type IAudioCaptureClientVtbl struct {
	ole.IUnknownVtbl
	GetBuffer         uintptr
	ReleaseBuffer     uintptr
	GetNextPacketSize uintptr
}

func (v *IAudioCaptureClient) VTable() *IAudioCaptureClientVtbl {
	return (*IAudioCaptureClientVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IAudioCaptureClient) GetBuffer(ppData **byte, pNumFramesToRead, pdwFlags *uint32, pu64DevicePosition, pu64QPCPosition *uint64) (err error) {
	hr, _, _ := syscall.Syscall6(
		v.VTable().GetBuffer,
		6,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(ppData)),
		uintptr(unsafe.Pointer(pNumFramesToRead)),
		uintptr(unsafe.Pointer(pdwFlags)),
		uintptr(unsafe.Pointer(pu64DevicePosition)),
		uintptr(unsafe.Pointer(pu64QPCPosition)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func (v *IAudioCaptureClient) ReleaseBuffer(NumFramesRead uint32) (err error) {
	hr, _, _ := syscall.Syscall(
		v.VTable().ReleaseBuffer,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(NumFramesRead),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func (v *IAudioCaptureClient) GetNextPacketSize(pNumFramesInNextPacket *uint32) (err error) {
	hr, _, _ := syscall.Syscall(
		v.VTable().GetNextPacketSize,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(pNumFramesInNextPacket)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
