package wca

import (
	"reflect"
	"syscall"
	"unsafe"

	"github.com/eakhadov/go-wca/internal/wca"
	"github.com/go-ole/go-ole"
)

//-
type IMMDevice struct {
	ole.IUnknown
}

type IMMDeviceVtbl struct {
	ole.IUnknownVtbl
	Activate          uintptr
	OpenPropertyStore uintptr
	GetId             uintptr
	GetState          uintptr
}

func (v *IMMDevice) VTable() *IMMDeviceVtbl {
	return (*IMMDeviceVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IMMDevice) Activate(iid *ole.GUID, dwClsCtx uint32, pActivationParams, ppInterface interface{}) (err error) {
	objValue := reflect.ValueOf(ppInterface).Elem()
	hr, _, _ := syscall.Syscall6(
		v.VTable().Activate,
		5,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(iid)),
		uintptr(unsafe.Pointer(&dwClsCtx)),
		0,
		objValue.Addr().Pointer(),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

// func (v *IMMDevice) OpenPropertyStore(stgmAccess uint32, ppProperties **IPropertyStore) (err error) {
// 	hr, _, _ := syscall.Syscall(
// 		v.VTable().OpenPropertyStore,
// 		3,
// 		uintptr(unsafe.Pointer(v)),
// 		uintptr(stgmAccess),
// 		uintptr(unsafe.Pointer(ppProperties)))
// 	if hr != 0 {
// 		err = ole.NewError(hr)
// 	}
// 	return
// }

func (v *IMMDevice) GetId(strId *string) (err error) {
	var strIdPtr uint64
	hr, _, _ := syscall.Syscall(
		v.VTable().GetId,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&strIdPtr)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	*strId = wca.LPWSTRToString(strIdPtr)
	ole.CoTaskMemFree(uintptr(strIdPtr))
	return
}

func (v *IMMDevice) GetState(pdwState *uint32) (err error) {
	hr, _, _ := syscall.Syscall(
		v.VTable().GetState,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(pdwState)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
