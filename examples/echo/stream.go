package main

import (
	"unsafe"

	"github.com/eakhadov/go-wca/pkg/wca"
	"github.com/go-ole/go-ole"
)

func initStream() (wfx *wca.WAVEFORMATEX, acc *wca.IAudioCaptureClient, arc *wca.IAudioRenderClient) {
	g_ACapture.GetMixFormat(&wfx)

	wfx.WFormatTag = 1
	wfx.NBlockAlign = (wfx.WBitsPerSample / 8) * wfx.NChannels
	wfx.NAvgBytesPerSec = wfx.NSamplesPerSec * uint32(wfx.NBlockAlign)
	wfx.CbSize = 0

	var dep wca.REFERENCE_TIME
	g_ACapture.GetDevicePeriod(&dep, nil)

	g_ACapture.Initialize(wca.AUDCLNT_SHAREMODE_SHARED, 0, dep, 0, wfx, nil)
	g_ARender.Initialize(wca.AUDCLNT_SHAREMODE_SHARED, 0, dep, 0, wfx, nil)

	g_ACapture.GetService(wca.IID_IAudioCaptureClient, &acc)
	g_ARender.GetService(wca.IID_IAudioRenderClient, &arc)

	return
}

func StartStream() {
	wfx, acc, arc := initStream()

	defer ole.CoTaskMemFree(uintptr(unsafe.Pointer(wfx)))
	defer acc.Release()
	defer arc.Release()

	g_ACapture.Start()
	g_ARender.Start()
	defer g_ACapture.Stop()
	defer g_ARender.Stop()

	c := struct {
		data         *byte
		aFS, flags   uint32
		dPos, qcpPos uint64
	}{}

	r := struct {
		data              *byte
		bFS, aFS, padding uint32
	}{}

	for {
		select {
		case <-g_ctx.Done():
			return
		default:
			if err := acc.GetBuffer(
				&c.data,
				&c.aFS,
				&c.flags,
				&c.dPos,
				&c.qcpPos,
			); err != nil || c.aFS == 0 {
				continue
			}

			if err := arc.GetBuffer(c.aFS, &r.data); err != nil {
				continue
			}

			lim := int(c.aFS) * int(wfx.NBlockAlign)

			startC := unsafe.Pointer(c.data)
			startR := unsafe.Pointer(r.data)

			for n := 0; n < lim; n++ {
				bC := (*byte)(unsafe.Pointer(uintptr(startC) + uintptr(n)))
				bR := (*byte)(unsafe.Pointer(uintptr(startR) + uintptr(n)))
				*bR = *bC
			}

			if err := acc.ReleaseBuffer(c.aFS); err != nil {
				return
			}
			if err := arc.ReleaseBuffer(c.aFS, 0); err != nil {
				return
			}
		}
	}
}
