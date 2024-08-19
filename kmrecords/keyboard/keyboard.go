package keyboard

import (
	"fmt"
	"kmrecords"
	"sync"
	"time"
	"unsafe"
)

type KeyBoradInfo struct {
	VK_Code   uintptr
	VK_Type   uint16
	Ctrl      bool
	Alt       bool
	Shift     bool
	Win       bool
	CapsLock  bool
	TimeStamp int64
}

var hook struct {
	sync.Mutex
	pointer uintptr
}

type KeyBoardEvent struct {
	Event uint16
	KeyBoradInfo
}

func defaultHook(c chan<- KeyBoardEvent) kmrecords.HookProc {
	return func(nCode int, wParam, lParam uintptr) uintptr {
		if lParam != 0 {
			point := (*kmrecords.KbDllHookStruct)(unsafe.Pointer(lParam))
			info := &KeyBoradInfo{
				VK_Code:   point.VkCode & 0xff,
				Ctrl:      kmrecords.GetAsyncKeyState(kmrecords.VK_CONTROL),
				Shift:     kmrecords.GetAsyncKeyState(kmrecords.VK_SHIFT),
				Alt:       kmrecords.GetAsyncKeyState(kmrecords.VK_MENU),
				Win:       kmrecords.GetAsyncKeyState(kmrecords.VK_LWIN | kmrecords.VK_RWIN),
				CapsLock:  kmrecords.GetAsyncKeyState(kmrecords.VK_CAPITAL),
				TimeStamp: time.Now().Unix(),
			}
			c <- KeyBoardEvent{
				Event:        uint16(wParam),
				KeyBoradInfo: *info,
			}
		}
		return kmrecords.CallNextHookEx(0, nCode, wParam, lParam)
	}
}

func use(fn kmrecords.HookProc, c chan<- KeyBoardEvent) error {
	hook.Lock()
	defer hook.Unlock()
	if hook.pointer != 0 {
		return fmt.Errorf("mouse: hook function is already install")
	}
	if c == nil {
		return fmt.Errorf("mouse: chan must not be nil")
	}
	if fn == nil {
		fn = defaultHook(c)
	}
	go func() {
		hhk := kmrecords.SetWindowsHook(kmrecords.WH_KEYBOARD_LL, fn)

		if hhk == 0 {
			panic("mouse: failed to install hook function")
		}

		hook.pointer = hhk

		var msg *kmrecords.MSG

		for {
			if hook.pointer == 0 {
				break
			}
			if result := kmrecords.GetMessageA(msg); result != 0 {
				fmt.Printf("msg: %v\n", msg)
				if result < 0 {
					// We don't care what's went wrong, ignore the result value.
					continue
				} else {
					kmrecords.TransLateMessage(msg)
					kmrecords.DispatchMessage(msg)
				}
			}
		}
	}()

	return nil
}
func stop() {
	hook.Lock()
	defer hook.Unlock()
	kmrecords.UnhookWindowsHook(hook.pointer)
}

func Use(c chan<- KeyBoardEvent) {
	use(nil, c)
}

func Stop() {
	stop()
}
