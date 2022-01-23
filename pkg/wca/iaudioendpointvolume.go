package wca

import (
	"math"
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IAudioEndpointVolume struct {
	ole.IUnknown
}

type IAudioEndpointVolumeVtbl struct {
	ole.IUnknownVtbl
	RegisterControlChangeNotify   uintptr
	UnregisterControlChangeNotify uintptr
	GetChannelCount               uintptr
	SetMasterVolumeLevel          uintptr
	SetMasterVolumeLevelScalar    uintptr
	GetMasterVolumeLevel          uintptr
	GetMasterVolumeLevelScalar    uintptr
	SetChannelVolumeLevel         uintptr
	SetChannelVolumeLevelScalar   uintptr
	GetChannelVolumeLevel         uintptr
	GetChannelVolumeLevelScalar   uintptr
	SetMute                       uintptr
	GetMute                       uintptr
	GetVolumeStepInfo             uintptr
	VolumeStepUp                  uintptr
	VolumeStepDown                uintptr
	QueryHardwareSupport          uintptr
	GetVolumeRange                uintptr
}

func (v *IAudioEndpointVolume) VTable() *IAudioEndpointVolumeVtbl {
	return (*IAudioEndpointVolumeVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IAudioEndpointVolume) RegisterControlChangeNotify() (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IAudioEndpointVolume) UnregisterControlChangeNotify() (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IAudioEndpointVolume) GetChannelCount(pnChannelCount *uint32) (err error) {
	hr, _, _ := syscall.Syscall(
		v.VTable().GetChannelCount,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(pnChannelCount)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func (v *IAudioEndpointVolume) SetMasterVolumeLevel(fLevelDB float32, pguidEventContext *ole.GUID) (err error) {
	levelDBValue := math.Float32bits(fLevelDB)
	hr, _, _ := syscall.Syscall(
		v.VTable().SetMasterVolumeLevel,
		3,
		uintptr(unsafe.Pointer(v)),
		uintptr(levelDBValue),
		uintptr(unsafe.Pointer(pguidEventContext)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func (v *IAudioEndpointVolume) SetMasterVolumeLevelScalar(fLevel float32, pguidEventContext *ole.GUID) (err error) {
	levelValue := math.Float32bits(fLevel)
	hr, _, _ := syscall.Syscall(
		v.VTable().SetMasterVolumeLevelScalar,
		3,
		uintptr(unsafe.Pointer(v)),
		uintptr(levelValue),
		uintptr(unsafe.Pointer(pguidEventContext)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func (v *IAudioEndpointVolume) GetMasterVolumeLevel(pfLevelDB *float32) (err error) {
	hr, _, _ := syscall.Syscall(
		v.VTable().GetMasterVolumeLevel,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(pfLevelDB)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func (v *IAudioEndpointVolume) GetMasterVolumeLevelScalar(pfLevel *float32) (err error) {
	hr, _, _ := syscall.Syscall(
		v.VTable().GetMasterVolumeLevelScalar,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(pfLevel)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func (v *IAudioEndpointVolume) SetChannelVolumeLevel(nChannel uint32, fLevelDB float32, pguidEventContext *ole.GUID) (err error) {
	levelDBValue := math.Float32bits(fLevelDB)
	hr, _, _ := syscall.Syscall6(
		v.VTable().SetChannelVolumeLevel,
		4,
		uintptr(unsafe.Pointer(v)),
		uintptr(nChannel),
		uintptr(levelDBValue),
		uintptr(unsafe.Pointer(pguidEventContext)),
		0,
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func (v *IAudioEndpointVolume) SetChannelVolumeLevelScalar(nChannel uint32, fLevel float32, pguidEventContext *ole.GUID) (err error) {
	levelValue := math.Float32bits(fLevel)
	hr, _, _ := syscall.Syscall6(
		v.VTable().SetChannelVolumeLevelScalar,
		4,
		uintptr(unsafe.Pointer(v)),
		uintptr(nChannel),
		uintptr(levelValue),
		uintptr(unsafe.Pointer(pguidEventContext)),
		0,
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func (v *IAudioEndpointVolume) GetChannelVolumeLevel(nChannel uint32, pfLevelDB *float32) (err error) {
	hr, _, _ := syscall.Syscall(
		v.VTable().GetChannelVolumeLevel,
		3,
		uintptr(unsafe.Pointer(v)),
		uintptr(nChannel),
		uintptr(unsafe.Pointer(pfLevelDB)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func (v *IAudioEndpointVolume) GetChannelVolumeLevelScalar(nChannel uint32, pfLevel *float32) (err error) {
	hr, _, _ := syscall.Syscall(
		v.VTable().GetChannelVolumeLevelScalar,
		3,
		uintptr(unsafe.Pointer(v)),
		uintptr(nChannel),
		uintptr(unsafe.Pointer(pfLevel)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func (v *IAudioEndpointVolume) SetMute(bMute bool, pguidEventContext *ole.GUID) (err error) {
	var muteValue uint32
	if bMute {
		muteValue = 1
	}
	hr, _, _ := syscall.Syscall(
		v.VTable().SetMute,
		3,
		uintptr(unsafe.Pointer(v)),
		uintptr(muteValue),
		uintptr(unsafe.Pointer(pguidEventContext)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func (v *IAudioEndpointVolume) GetMute(pbMute *bool) (err error) {
	hr, _, _ := syscall.Syscall(
		v.VTable().GetMute,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(pbMute)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func (v *IAudioEndpointVolume) GetVolumeStepInfo(pnStep, pnStepCount *uint32) (err error) {
	hr, _, _ := syscall.Syscall(
		v.VTable().GetVolumeStepInfo,
		3,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(pnStep)),
		uintptr(unsafe.Pointer(pnStepCount)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func (v *IAudioEndpointVolume) VolumeStepUp(pguidEventContext *ole.GUID) (err error) {
	hr, _, _ := syscall.Syscall(
		v.VTable().VolumeStepUp,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(pguidEventContext)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func (v *IAudioEndpointVolume) VolumeStepDown(pguidEventContext *ole.GUID) (err error) {
	hr, _, _ := syscall.Syscall(
		v.VTable().VolumeStepDown,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(pguidEventContext)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func (v *IAudioEndpointVolume) QueryHardwareSupport(pdwHardwareSupportMask *uint32) (err error) {
	hr, _, _ := syscall.Syscall(
		v.VTable().QueryHardwareSupport,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(pdwHardwareSupportMask)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func (v *IAudioEndpointVolume) GetVolumeRange(pflVolumeMindB, pflVolumeMaxdB, pflVolumeIncrementdB *float32) (err error) {
	hr, _, _ := syscall.Syscall6(
		v.VTable().GetVolumeRange,
		4,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(pflVolumeMindB)),
		uintptr(unsafe.Pointer(pflVolumeMaxdB)),
		uintptr(unsafe.Pointer(pflVolumeIncrementdB)),
		0,
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
