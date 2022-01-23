package wca

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IConnector struct {
	ole.IUnknown
}

type IConnectorVtbl struct {
	ole.IUnknownVtbl
	GetType                   uintptr
	GetDataFlow               uintptr
	ConnectTo                 uintptr
	Disconnect                uintptr
	IsConnected               uintptr
	GetConnectedTo            uintptr
	GetConnectorIdConnectedTo uintptr
	GetDeviceIdConnectedTo    uintptr
}

func (v *IConnector) VTable() *IConnectorVtbl {
	return (*IConnectorVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IConnector) GetType(pType *ConnectorType) (err error) {
	hr, _, _ := syscall.Syscall(
		v.VTable().GetType,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(pType)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func (v *IConnector) GetDataFlow(pFlow *DataFlow) (err error) {
	hr, _, _ := syscall.Syscall(
		v.VTable().GetDataFlow,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(pFlow)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func (v *IConnector) ConnectTo(pFlow *DataFlow) (err error) {
	return
}

func (v *IConnector) Disconnect(pFlow *DataFlow) (err error) {
	return
}

func (v *IConnector) IsConnected(pbConnected *bool) (err error) {
	hr, _, _ := syscall.Syscall(
		v.VTable().IsConnected,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(pbConnected)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func (v *IConnector) GetConnectedTo(ppConTo **IConnector) (err error) {
	hr, _, _ := syscall.Syscall(
		v.VTable().GetConnectedTo,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(ppConTo)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
