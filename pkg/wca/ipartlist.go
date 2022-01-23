package wca

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IPartsList struct {
	ole.IUnknown
}

type IPartsListVtbl struct {
	ole.IUnknownVtbl
	GetCount uintptr
	GetPart  uintptr
}

func (v *IPartsList) VTable() *IPartsListVtbl {
	return (*IPartsListVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IPartsList) GetCount(pCount *uint32) (err error) {
	hr, _, _ := syscall.Syscall(
		v.VTable().GetCount,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(pCount)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func (v *IPartsList) GetPart(nIndex uint32, ppPart **IPart) (err error) {
	hr, _, _ := syscall.Syscall(
		v.VTable().GetPart,
		3,
		uintptr(unsafe.Pointer(v)),
		uintptr(nIndex),
		uintptr(unsafe.Pointer(ppPart)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
