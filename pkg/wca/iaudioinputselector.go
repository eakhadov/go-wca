package wca

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IAudioInputSelector struct {
	ole.IUnknown
}

type IAudioInputSelectorVtbl struct {
	ole.IUnknownVtbl
	GetSelection uintptr
	SetSelection uintptr
}

func (v *IAudioInputSelector) VTable() *IAudioInputSelectorVtbl {
	return (*IAudioInputSelectorVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IAudioInputSelector) GetSelection(pnIdSelected *uint32) (err error) {
	hr, _, _ := syscall.Syscall(
		v.VTable().GetSelection,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(pnIdSelected)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func (v *IAudioInputSelector) SetSelection(nIdSelect uint32, pguidEventContext *ole.GUID) (err error) {
	hr, _, _ := syscall.Syscall(
		v.VTable().SetSelection,
		3,
		uintptr(unsafe.Pointer(v)),
		uintptr(nIdSelect),
		uintptr(unsafe.Pointer(pguidEventContext)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
