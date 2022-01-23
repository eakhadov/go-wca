package wca

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IAudioRenderClient struct {
	ole.IUnknown
}

type IAudioRenderClientVtbl struct {
	ole.IUnknownVtbl
	GetBuffer     uintptr
	ReleaseBuffer uintptr
}

func (v *IAudioRenderClient) VTable() *IAudioRenderClientVtbl {
	return (*IAudioRenderClientVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IAudioRenderClient) GetBuffer(NumFramesRequested uint32, ppData **byte) (err error) {
	hr, _, _ := syscall.Syscall(
		v.VTable().GetBuffer,
		3,
		uintptr(unsafe.Pointer(v)),
		uintptr(NumFramesRequested),
		uintptr(unsafe.Pointer(ppData)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func (v *IAudioRenderClient) ReleaseBuffer(NumFramesWritten, dwFlags uint32) (err error) {
	hr, _, _ := syscall.Syscall(
		v.VTable().ReleaseBuffer,
		3,
		uintptr(unsafe.Pointer(v)),
		uintptr(NumFramesWritten),
		uintptr(dwFlags))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
