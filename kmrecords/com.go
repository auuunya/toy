package kmrecords

import (
	"fmt"
	"log"
	"syscall"
	"unsafe"
)

var (
	user32 = syscall.NewLazyDLL("User32.dll")

	procGetCursorPos        = user32.NewProc("GetCursorPos")
	procSetWindowsHookExA   = user32.NewProc("SetWindowsHookExA")
	procCallNextHookEx      = user32.NewProc("CallNextHookEx")
	procGetMessageA         = user32.NewProc("GetMessageA")
	procRegisterHotKey      = user32.NewProc("RegisterHotKey")
	procGetAsyncKeyState    = user32.NewProc("GetAsyncKeyState")
	procTranslateMessage    = user32.NewProc("TranslateMessage")
	procDispatchMessageW    = user32.NewProc("DispatchMessageW")
	procUnhookWindowsHookEx = user32.NewProc("UnhookWindowsHookEx")
	procPostQuitMessage     = user32.NewProc("PostQuitMessage")
)

func SetWindowsHook(wh int, callback HookProc) uintptr {
	hook, _, err := procSetWindowsHookExA.Call(
		uintptr(wh),
		uintptr(syscall.NewCallback(callback)),
		uintptr(0),
		uintptr(0),
	)
	if err != nil && err.Error() != "The operation completed successfully." {
		log.Fatalf("error: %v\n", err)
	}
	return hook
}

func GetMsgProc(nCode int, wParam, lParam uintptr) uintptr {
	return CallNextHookEx(0, nCode, wParam, lParam)
}

func CallNextHookEx(hhk uintptr, nCode int, wParam uintptr, lParam uintptr) uintptr {
	ret, _, _ := procCallNextHookEx.Call(
		uintptr(hhk),
		uintptr(nCode),
		uintptr(wParam),
		uintptr(lParam),
	)
	return ret
}

func GetMessageA(msg *MSG) uintptr {
	ret, _, _ := procGetMessageA.Call(
		uintptr(unsafe.Pointer(&msg)),
		uintptr(0),
		uintptr(0),
		uintptr(0),
	)
	return ret
}

func GetCursorPoint() *Point {
	var pointer *Point
	ret, _, _ := procGetCursorPos.Call(
		uintptr(unsafe.Pointer(&pointer)),
	)
	fmt.Printf("ret: %#v\n", ret)
	fmt.Printf("pointer: %#v\n", pointer)
	return pointer
}

func GetAsyncKeyState(key int) bool {
	state, _, _ := procGetAsyncKeyState.Call(
		uintptr(key),
	)
	return state != 0
}

func TransLateMessage(msg *MSG) int32 {
	ret, _, _ := procTranslateMessage.Call(
		uintptr(unsafe.Pointer(msg)),
	)
	return int32(ret)
}
func DispatchMessage(msg *MSG) int32 {
	ret, _, _ := procDispatchMessageW.Call(uintptr(unsafe.Pointer(msg)))

	return int32(ret)
}

func UnhookWindowsHook(hhk uintptr) {
	procUnhookWindowsHookEx.Call(
		uintptr(hhk),
	)
}
