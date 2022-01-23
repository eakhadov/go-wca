package wca

import (
	"reflect"
	"syscall"
	"unsafe"

	"github.com/eakhadov/go-wca/internal/wca"
	"github.com/go-ole/go-ole"
)

type IPart struct {
	ole.IUnknown
}

type IPartVtbl struct {
	ole.IUnknownVtbl
	GetName                         uintptr
	GetLocalId                      uintptr
	GetGlobalId                     uintptr
	GetPartType                     uintptr
	GetSubType                      uintptr
	GetControlInterfaceCount        uintptr
	GetControlInterface             uintptr
	EnumPartsIncoming               uintptr
	EnumPartsOutgoing               uintptr
	GetTopologyObject               uintptr
	Activate                        uintptr
	RegisterControlChangeCallback   uintptr
	UnregisterControlChangeCallback uintptr
}

func (v *IPart) VTable() *IPartVtbl {
	return (*IPartVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IPart) GetName(retVal *string) (err error) {
	var retValPtr uint64
	hr, _, _ := syscall.Syscall(
		v.VTable().GetName,
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

func (v *IPart) GetLocalId(pnId *uint32) (err error) {
	hr, _, _ := syscall.Syscall(
		v.VTable().GetLocalId,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(pnId)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func (v *IPart) GetGlobalId() (err error) {
	return
}

func (v *IPart) GetPartType(pPartType *PartType) (err error) {
	hr, _, _ := syscall.Syscall(
		v.VTable().GetPartType,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(pPartType)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func (v *IPart) GetSubType() (err error) {
	return
}

func (v *IPart) GetControlInterfaceCount() (err error) {
	return
}

func (v *IPart) GetControlInterface() (err error) {
	return
}

func (v *IPart) EnumPartsIncoming() (err error) {
	return
}

func (v *IPart) EnumPartsOutgoing(ppParts **IPartsList) (err error) {
	hr, _, _ := syscall.Syscall(
		v.VTable().EnumPartsOutgoing,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(ppParts)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func (v *IPart) GetTopologyObject() (err error) {
	return
}

func (v *IPart) Activate(dwClsContext uint32, refiid *ole.GUID, ppvObject interface{}) (err error) {
	objValue := reflect.ValueOf(ppvObject).Elem()
	hr, _, _ := syscall.Syscall6(
		v.VTable().Activate,
		4,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&dwClsContext)),
		uintptr(unsafe.Pointer(refiid)),
		objValue.Addr().Pointer(),
		0,
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func (v *IPart) RegisterControlChangeCallback() (err error) {
	return
}

func (v *IPart) UnregisterControlChangeCallback() (err error) {
	return
}
