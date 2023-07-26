//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type ICoreWebView2Environment7Vtbl struct {
	IUnknownVtbl
	GetUserDataFolder ComProc
}

type ICoreWebView2Environment7 struct {
	Vtbl *ICoreWebView2Environment7Vtbl
}

func (i *ICoreWebView2Environment7) AddRef() uintptr {
	refCounter, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter
}

func (i *ICoreWebView2) GetICoreWebView2Environment7() *ICoreWebView2Environment7 {
	var result *ICoreWebView2Environment7

	iidICoreWebView2Environment7 := NewGUID("{43C22296-3BBD-43A4-9C00-5C0DF6DD29A2}")
	_, _, _ = i.Vtbl.QueryInterface.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(iidICoreWebView2Environment7)),
		uintptr(unsafe.Pointer(&result)))

	return result
}

func (i *ICoreWebView2Environment7) GetUserDataFolder() (*string, error) {
	// Create *uint16 to hold result
	var _value *uint16

	hr, _, err := i.Vtbl.GetUserDataFolder.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_value)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	}
	// Get result and cleanup
	value := ptr(UTF16PtrToString(_value))
	CoTaskMemFree(unsafe.Pointer(_value))
	return value, err
}
