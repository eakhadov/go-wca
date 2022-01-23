package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/eakhadov/go-wca/pkg/wca"
	"github.com/go-ole/go-ole"
)

var g_ACapture *wca.IAudioClient
var g_ARender *wca.IAudioClient

var g_ctx context.Context

func init() {
	var (
		mmdc, mmdr *wca.IMMDevice
		mmdentr    *wca.IMMDeviceEnumerator
	)

	ole.CoInitializeEx(0, ole.COINIT_MULTITHREADED)
	defer ole.CoUninitialize()

	wca.CoCreateInstance(
		wca.CLSID_MMDeviceEnumerator,
		0,
		wca.CLSCTX_ALL,
		wca.IID_IMMDeviceEnumerator,
		&mmdentr,
	)
	defer mmdentr.Release()

	mmdentr.GetDefaultAudioEndpoint(wca.ECapture, wca.EConsole, &mmdc)
	mmdentr.GetDefaultAudioEndpoint(wca.ERender, wca.EConsole, &mmdr)

	mmdc.Activate(wca.IID_IAudioClient, wca.CLSCTX_ALL, nil, &g_ACapture)
	mmdr.Activate(wca.IID_IAudioClient, wca.CLSCTX_ALL, nil, &g_ARender)

	mmdc.Release()
	mmdr.Release()
}

func main() {

	createExitChan()

	fmt.Println("Start capturing with shared timer driven mode")
	fmt.Println("Press Ctrl-C to stop capturing")

	StartStream()

	for {
		select {
		case <-g_ctx.Done():
			break
		}
		break
	}

	fmt.Println("\n\nStop capturing")

	g_ACapture.Release()
	g_ARender.Release()
}

func createExitChan() {
	sCh := make(chan os.Signal, 1)
	signal.Notify(sCh, os.Interrupt)

	var cancel context.CancelFunc
	g_ctx, cancel = context.WithCancel(context.Background())

	go func() {
		select {
		case <-sCh:
			cancel()
		}
	}()
}
