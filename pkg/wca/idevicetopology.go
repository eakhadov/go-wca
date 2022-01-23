package wca

import (
	"syscall"
	"unsafe"

	"github.com/eakhadov/go-wca/internal/wca"
	"github.com/go-ole/go-ole"
)

type IDeviceTopology struct {
	ole.IUnknown
}

type IDeviceTopologyVtbl struct {
	ole.IUnknownVtbl
	GetConnectorCount uintptr
	GetConnector      uintptr
	GetSubunitCount   uintptr
	GetSubunit        uintptr
	GetPartById       uintptr
	GetDeviceId       uintptr
	GetSignalPath     uintptr
}

func (v *IDeviceTopology) VTable() *IDeviceTopologyVtbl {
	return (*IDeviceTopologyVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IDeviceTopology) GetConnectorCount(pCount *uint32) (err error) {
	hr, _, _ := syscall.Syscall(
		v.VTable().GetConnectorCount,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(pCount)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func (v *IDeviceTopology) GetConnector(nIndex uint32, ppConnector **IConnector) (err error) {
	hr, _, _ := syscall.Syscall(
		v.VTable().GetConnector,
		3,
		uintptr(unsafe.Pointer(v)),
		uintptr(nIndex),
		uintptr(unsafe.Pointer(ppConnector)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func (v *IDeviceTopology) GetDeviceId(retVal *string) (err error) {
	var retValPtr uint64
	hr, _, _ := syscall.Syscall(
		v.VTable().GetDeviceId,
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
