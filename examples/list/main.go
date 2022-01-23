package main

import (
	"fmt"

	"github.com/eakhadov/go-wca/pkg/wca"
	"github.com/go-ole/go-ole"
	"github.com/mitchellh/go-ps"
)

var pEpVol *wca.IAudioEndpointVolume
var pAudioSessionEnumerator *wca.IAudioSessionEnumerator

func init() {
	var (
		pDevice               *wca.IMMDevice
		pDeviceEnumerator     *wca.IMMDeviceEnumerator
		pAudioSessionManager2 *wca.IAudioSessionManager2
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

	pDevice.Activate(wca.IID_IAudioSessionManager2, wca.CLSCTX_ALL, nil, &pAudioSessionManager2)
	defer pAudioSessionManager2.Release()

	pAudioSessionManager2.GetSessionEnumerator(&pAudioSessionEnumerator)
}

func main() {
	defer pEpVol.Release()
	defer pAudioSessionEnumerator.Release()

	printInfo()
}

func printInfo() {
	var (
		sessionCount int
		masterMuted  bool
		masterVolume float32
	)

	pEpVol.GetMute(&masterMuted)
	pEpVol.GetMasterVolumeLevelScalar(&masterVolume)

	pAudioSessionEnumerator.GetCount(&sessionCount)

	fmt.Printf("\n\nAudio Sessions - %d | Muted - %t | Volume = %d\n", sessionCount, masterMuted, int(masterVolume*100))
	fmt.Println("---------------------------------------------------------------------------------")
	fmt.Printf("[id] PID  - %-21s - %-30s - %-4s\n", "Process", "Window Title", "Volume")
	fmt.Println("---------------------------------------------------------------------------------")

	var (
		pSessionControl    *wca.IAudioSessionControl
		pSessionControl2   *wca.IAudioSessionControl2
		pSimpleAudioVolume *wca.ISimpleAudioVolume
	)

	for i := 0; i < sessionCount; i++ {

		if err := pAudioSessionEnumerator.GetSession(i, &pSessionControl); err != nil {
			fmt.Println(err)
			continue
		}

		pSessionControl.PutQueryInterface(wca.IID_IAudioSessionControl2, &pSessionControl2)
		pSessionControl.PutQueryInterface(wca.IID_ISimpleAudioVolume, &pSimpleAudioVolume)

		var (
			title  string
			procid uint32
			vol    float32
		)
		pSessionControl2.GetProcessId(&procid)
		pSimpleAudioVolume.GetMasterVolume(&vol)

		p, err := ps.FindProcess(int(procid))

		if err != nil {
			fmt.Println("Error: ", err)
		}

		fmt.Printf("[% 2d]% 6d - %-20s - %-30s - %-4d\n",
			i,
			procid,
			p.Executable(),
			title,
			int(vol*masterVolume*100),
		)

	}
}
