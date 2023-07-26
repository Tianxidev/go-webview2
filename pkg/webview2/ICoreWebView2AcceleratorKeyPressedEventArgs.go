//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type ICoreWebView2AcceleratorKeyPressedEventArgsVtbl struct {
	IUnknownVtbl
	GetKeyEventKind      ComProc
	GetVirtualKey        ComProc
	GetKeyEventLParam    ComProc
	GetPhysicalKeyStatus ComProc
	GetHandled           ComProc
	PutHandled           ComProc
}

type ICoreWebView2AcceleratorKeyPressedEventArgs struct {
	Vtbl *ICoreWebView2AcceleratorKeyPressedEventArgsVtbl
}

func (i *ICoreWebView2AcceleratorKeyPressedEventArgs) AddRef() uintptr {
	refCounter, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter
}

func (i *ICoreWebView2AcceleratorKeyPressedEventArgs) GetKeyEventKind() (*COREWEBVIEW2_KEY_EVENT_KIND, error) {

	var keyEventKind *COREWEBVIEW2_KEY_EVENT_KIND

	hr, _, err := i.Vtbl.GetKeyEventKind.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&keyEventKind)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	}
	return keyEventKind, err
}

func (i *ICoreWebView2AcceleratorKeyPressedEventArgs) GetVirtualKey() (*uint, error) {

	var virtualKey *uint

	hr, _, err := i.Vtbl.GetVirtualKey.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&virtualKey)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	}
	return virtualKey, err
}

func (i *ICoreWebView2AcceleratorKeyPressedEventArgs) GetKeyEventLParam() (*int, error) {

	var lParam *int

	hr, _, err := i.Vtbl.GetKeyEventLParam.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&lParam)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	}
	return lParam, err
}

func (i *ICoreWebView2AcceleratorKeyPressedEventArgs) GetPhysicalKeyStatus() (*COREWEBVIEW2_PHYSICAL_KEY_STATUS, error) {

	var physicalKeyStatus *COREWEBVIEW2_PHYSICAL_KEY_STATUS

	hr, _, err := i.Vtbl.GetPhysicalKeyStatus.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&physicalKeyStatus)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	}
	return physicalKeyStatus, err
}

func (i *ICoreWebView2AcceleratorKeyPressedEventArgs) GetHandled() (*bool, error) {
	// Create int32 to hold bool result
	var _handled int32

	hr, _, err := i.Vtbl.GetHandled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_handled)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	}
	// Get result and cleanup
	handled := ptr(_handled != 0)
	return handled, err
}

func (i *ICoreWebView2AcceleratorKeyPressedEventArgs) PutHandled(handled bool) error {

	hr, _, err := i.Vtbl.PutHandled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&handled)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return syscall.Errno(hr)
	}
	return err
}
