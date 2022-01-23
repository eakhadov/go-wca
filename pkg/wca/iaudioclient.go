package wca

import (
	"reflect"
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IAudioClient struct {
	ole.IUnknown
}

type IAudioClientVtbl struct {
	ole.IUnknownVtbl
	Initialize        uintptr
	GetBufferSize     uintptr
	GetStreamLatency  uintptr
	GetCurrentPadding uintptr
	IsFormatSupported uintptr
	GetMixFormat      uintptr
	GetDevicePeriod   uintptr
	Start             uintptr
	Stop              uintptr
	Reset             uintptr
	SetEventHandle    uintptr
	GetService        uintptr
}

func (v *IAudioClient) VTable() *IAudioClientVtbl {
	return (*IAudioClientVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IAudioClient) Initialize(ShareMode, StreamFlags uint32, hnsBufferDuration, hnsPeriodicity REFERENCE_TIME, pFormat *WAVEFORMATEX, AudioSessionGuid *ole.GUID) (err error) {
	hr, _, _ := syscall.Syscall9(
		v.VTable().Initialize,
		7,
		uintptr(unsafe.Pointer(v)),
		uintptr(ShareMode),
		uintptr(StreamFlags),
		uintptr(hnsBufferDuration),
		uintptr(hnsPeriodicity),
		uintptr(unsafe.Pointer(pFormat)),
		uintptr(unsafe.Pointer(AudioSessionGuid)),
		0,
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func (v *IAudioClient) GetBufferSize(pNumBufferFrames *uint32) (err error) {
	hr, _, _ := syscall.Syscall(
		v.VTable().GetBufferSize,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(pNumBufferFrames)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func (v *IAudioClient) GetStreamLatency(phnsLatency *REFERENCE_TIME) (err error) {
	hr, _, _ := syscall.Syscall(
		v.VTable().GetStreamLatency,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(phnsLatency)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func (v *IAudioClient) GetCurrentPadding(pNumPaddingFrames *uint32) (err error) {
	hr, _, _ := syscall.Syscall(
		v.VTable().GetCurrentPadding,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(pNumPaddingFrames)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func (v *IAudioClient) IsFormatSupported(ShareMode uint32, pFormat *WAVEFORMATEX, ppClosestMatch **WAVEFORMATEX) (err error) {
	hr, _, _ := syscall.Syscall6(
		v.VTable().IsFormatSupported,
		4,
		uintptr(unsafe.Pointer(v)),
		uintptr(ShareMode),
		uintptr(unsafe.Pointer(pFormat)),
		uintptr(unsafe.Pointer(ppClosestMatch)),
		0,
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func (v *IAudioClient) GetMixFormat(ppDeviceFormat **WAVEFORMATEX) (err error) {
	hr, _, _ := syscall.Syscall(
		v.VTable().GetMixFormat,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(ppDeviceFormat)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func (v *IAudioClient) GetDevicePeriod(phnsDefaultDevicePeriod, phnsMinimumDevicePeriod *REFERENCE_TIME) (err error) {
	hr, _, _ := syscall.Syscall(
		v.VTable().GetDevicePeriod,
		3,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(phnsDefaultDevicePeriod)),
		uintptr(unsafe.Pointer(phnsMinimumDevicePeriod)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func (v *IAudioClient) Start() (err error) {
	hr, _, _ := syscall.Syscall(
		v.VTable().Start,
		1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func (v *IAudioClient) Stop() (err error) {
	hr, _, _ := syscall.Syscall(
		v.VTable().Stop,
		1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func (v *IAudioClient) Reset() (err error) {
	hr, _, _ := syscall.Syscall(
		v.VTable().Reset,
		1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func (v *IAudioClient) SetEventHandle(eventHandle uintptr) (err error) {
	hr, _, _ := syscall.Syscall(
		v.VTable().SetEventHandle,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(eventHandle),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func (v *IAudioClient) GetService(riid *ole.GUID, ppv interface{}) (err error) {
	objValue := reflect.ValueOf(ppv).Elem()
	hr, _, _ := syscall.Syscall(
		v.VTable().GetService,
		3,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(riid)),
		objValue.Addr().Pointer())
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
