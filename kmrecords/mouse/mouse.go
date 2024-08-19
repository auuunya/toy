package mouse

import (
	"fmt"
	"kmrecords"
	"sync"
	"time"
	"unsafe"
)

type MouseInfo struct {
	X         int32
	Y         int32
	Ctrl      bool
	Alt       bool
	Shift     bool
	TimeStamp int64
}

var hook struct {
	sync.Mutex
	pointer uintptr
}

type MouseEvent struct {
	Event uint16
	MouseInfo
}

func defaultHook(c chan<- MouseEvent) kmrecords.HookProc {
	return func(nCode int, wParam, lParam uintptr) uintptr {
		if lParam != 0 {
			point := (*kmrecords.MsllHookStruct)(unsafe.Pointer(lParam)).Pt
			mouseinfo := &MouseInfo{
				X:         point.X,
				Y:         point.Y,
				Ctrl:      kmrecords.GetAsyncKeyState(kmrecords.VK_CONTROL),
				Shift:     kmrecords.GetAsyncKeyState(kmrecords.VK_SHIFT),
				Alt:       kmrecords.GetAsyncKeyState(kmrecords.VK_MENU),
				TimeStamp: time.Now().Unix(),
			}
			c <- MouseEvent{
				Event:     uint16(wParam),
				MouseInfo: *mouseinfo,
			}
		}
		return kmrecords.CallNextHookEx(0, nCode, wParam, lParam)
	}
}

func use(fn kmrecords.HookProc, c chan<- MouseEvent) error {
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
		hhk := kmrecords.SetWindowsHook(kmrecords.WH_MOUSE_LL, fn)

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

func Use(c chan<- MouseEvent) {
	use(nil, c)
}

func Stop() {
	stop()
}
