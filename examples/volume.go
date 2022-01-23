package main

import (
	"fmt"

	"github.com/eakhadov/go-wca/pkg/wca"
	"github.com/go-ole/go-ole"
)

var pEpVol *wca.IAudioEndpointVolume

func init() {
	var (
		pDevice           *wca.IMMDevice
		pDeviceEnumerator *wca.IMMDeviceEnumerator
	)

	ole.CoInitializeEx(0, ole.COINIT_MULTITHREADED)
	defer ole.CoUninitialize()

	wca.CoCreateInstance(
		wca.CLSID_MMDeviceEnumerator,
		0,
		wca.CLSCTX_ALL,
		wca.IID_IMMDeviceEnumerator,
		&pDeviceEnumerator,
	)
	defer pDeviceEnumerator.Release()

	pDeviceEnumerator.GetDefaultAudioEndpoint(wca.ERender, wca.EMultimedia, &pDevice)
	defer pDevice.Release()

	pDevice.Activate(wca.IID_IAudioEndpointVolume, wca.CLSCTX_ALL, nil, &pEpVol)
}

func main() {
	defer pEpVol.Release()

	var (
		masterMuted  bool
		masterVolume float32
		chount       uint32
	)

	pEpVol.GetMute(&masterMuted)
	pEpVol.GetMasterVolumeLevelScalar(&masterVolume)
	pEpVol.GetChannelCount(&chount)

	fmt.Printf("%t | %d | %d\n\n", masterMuted, int(masterVolume*100), chount)

	// fmt.Scanln()
}
