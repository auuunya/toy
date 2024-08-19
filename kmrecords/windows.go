package kmrecords

// mouse
var (
	WM_CAPTURECHANGED  UINT_16 = 0x0215
	WM_LBUTTONDBLCLK           = 0x0203
	WM_LBUTTONDOWN     UINT_16 = 0x0201
	WM_LBUTTONUP       UINT_16 = 0x0202
	WM_MBUTTONDBLCLK           = 0x0209
	WM_MBUTTONDOWN             = 0x0207
	WM_MBUTTONUP               = 0x0208
	WM_MOUSEACTIVATE           = 0x0021
	WM_MOUSEHOVER              = 0x02A1
	WM_MOUSEHWHEEL             = 0x020E
	WM_MOUSELEAVE              = 0x02A3
	WM_MOUSEMOVE       UINT_16 = 0x0200
	WM_MOUSEWHEEL      UINT_16 = 0x020A
	WM_NCHITTEST               = 0x0084
	WM_NCLBUTTONDBLCLK         = 0x00A3
	WM_NCLBUTTONDOWN           = 0x00A1
	WM_NCLBUTTONUP             = 0x00A2
	WM_NCMBUTTONDBLCLK         = 0x00A9
	WM_NCMBUTTONDOWN           = 0x00A7
	WM_NCMBUTTONUP             = 0x00A8
	WM_NCMOUSEHOVER            = 0x02A0
	WM_NCMOUSELEAVE            = 0x02A2
	WM_NCMOUSEMOVE             = 0x00A0
	WM_NCRBUTTONDBLCLK         = 0x00A6
	WM_NCRBUTTONDOWN           = 0x00A4
	WM_NCRBUTTONUP             = 0x00A5
	WM_NCXBUTTONDBLCLK         = 0x00AD
	WM_NCXBUTTONDOWN           = 0x00AB
	WM_NCXBUTTONUP             = 0x00AC
	WM_RBUTTONDBLCLK           = 0x0206
	WM_RBUTTONDOWN     UINT_16 = 0x0204
	WM_RBUTTONUP       UINT_16 = 0x0205
	WM_XBUTTONDBLCLK           = 0x020D
	WM_XBUTTONDOWN             = 0x020B
	WM_XBUTTONUP               = 0x020C
)

const (
	// WM_NCXBUTTONDBLCLK
	XBUTTON1 = 0x0001 // 双击第一个 X 按钮
	XBUTTON2 = 0x0002 // 双击第二个 X 按钮
)

var (
	MOUSEEVENTF_MOVE            = 0x0001
	MOUSEEVENTF_LEFTDOWN        = 0x0002
	MOUSEEVENTF_LEFTUP          = 0x0004
	MOUSEEVENTF_RIGHTDOWN       = 0x0008
	MOUSEEVENTF_RIGHTUP         = 0x0010
	MOUSEEVENTF_MIDDLEDOWN      = 0x0020
	MOUSEEVENTF_MIDDLEUP        = 0x0040
	MOUSEEVENTF_XDOWN           = 0x0080
	MOUSEEVENTF_XUP             = 0x0100
	MOUSEEVENTF_WHEEL           = 0x0800
	MOUSEEVENTF_HWHEEL          = 0x1000
	MOUSEEVENTF_MOVE_NOCOALESCE = 0x2000
	MOUSEEVENTF_VIRTUALDESK     = 0x4000
	MOUSEEVENTF_ABSOLUTE        = 0x8000
)

var (
	// WM_LBUTTONDBLCLK
	MK_CONTROL  = 0x0008 // 按下了 CTRL 键
	MK_LBUTTON  = 0x0001 // 按下了鼠标左键
	MK_MBUTTON  = 0x0010 // 按下了鼠标中键
	MK_RBUTTON  = 0x0002 // 按下了鼠标右键
	MK_SHIFT    = 0x0004 // 按下了 SHIFT 键
	MK_XBUTTON1 = 0x0020 // 按下了第一个 X 按钮
	MK_XBUTTON2 = 0x0040 // 按下了第二个 X 按钮
)

const (
	MA_ACTIVATE         int = iota + 1 // 激活窗口，并且不丢弃鼠标消息
	MA_ACTIVATEANDEAT                  // 激活窗口，并丢弃鼠标消息
	MA_NOACTIVATE                      // 不激活窗口，并且不丢弃鼠标消息
	MA_NOACTIVATEANDEAT                // 不激活窗口，但丢弃鼠标消息
)

const (
	INPUT_MOUSE = iota
	INPUT_KEYBOARD
	INPUT_HARDWARE
)

type Point struct {
	X, Y int32
}

type MouseInput struct {
	Dx          int32
	Dy          int32
	MouseData   uintptr
	DwFlags     uintptr
	Time        uintptr
	DwExtraInfo uintptr
}
type TagInput struct {
	Type           int
	DummyUoionName *DummyUoionName
}

type DummyUoionName struct {
	*MouseInput
	*KeyBDInput
	*HardWareInput
}

type KeyBDInput struct {
	Wvk         uint16
	Wscan       uint16
	DwFlags     uintptr
	Time        uintptr
	DwExtraInfo uintptr
}

type HardWareInput struct {
	UMsg    uintptr
	WParamL uint16
	WParamH uint16
}
type MsllHookStruct struct {
	Pt          Point
	MouseData   uintptr
	Flags       uintptr
	Time        uintptr
	DwExtraInfo uintptr
}

// keyboard
var (
	WM_ACTIVATE            = 0x0006
	WM_APPCOMMAND          = 0x0319
	WM_CHAR                = 0x0102
	WM_DEADCHAR            = 0x0103
	WM_HOTKEY              = 0x0312
	WM_KEYDOWN     UINT_16 = 0x0100
	WM_KEYUP       UINT_16 = 0x0101
	WM_KILLFOCUS           = 0x0008
	WM_SETFOCUS            = 0x0007
	WM_SYSDEADCHAR         = 0x0107
	WM_SYSKEYDOWN  UINT_16 = 0x0104
	WM_SYSKEYUP    UINT_16 = 0x0105
	WM_UNICHAR             = 0x0109
)

var (
	IDHOT_SNAPDESKTOP = -2
	IDHOT_SNAPWINDOW  = -1
)

var (
	MOD_ALT     = 0x0001
	MOD_CONTROL = 0x0002
	MOD_SHIFT   = 0x0004
	MOD_WIN     = 0x0008
)

var (
	// https://learn.microsoft.com/zh-cn/windows/win32/inputdev/virtual-key-codes
	VK_LBUTTON  = 0x01
	VK_RBUTTON  = 0x02
	VK_CANCEL   = 0x03
	VK_MBUTTON  = 0x04
	VK_XBUTTON1 = 0x05
	VK_XBUTTON2 = 0x06
	_           = 0x07
	VK_BACK     = 0x08
	VK_TAB      = 0x09
	// _=0x0A-0B
	VK_CLEAR  = 0x0C
	VK_RETURN = 0x0D
	// _=	0x0E-0F
	VK_SHIFT      = 0x10
	VK_CONTROL    = 0x11
	VK_MENU       = 0x12
	VK_PAUSE      = 0x13
	VK_CAPITAL    = 0x14
	VK_KANA       = 0x15
	VK_HANGUL     = 0x15
	VK_IME_ON     = 0x16
	VK_JUNJA      = 0x17
	VK_FINAL      = 0x18
	VK_HANJA      = 0x19
	VK_KANJI      = 0x19
	VK_IME_OFF    = 0x1A
	VK_ESCAPE     = 0x1B
	VK_CONVERT    = 0x1C
	VK_NONCONVERT = 0x1D
	VK_ACCEPT     = 0x1E
	VK_MODECHANGE = 0x1F
	VK_SPACE      = 0x20
	VK_PRIOR      = 0x21
	VK_NEXT       = 0x22
	VK_END        = 0x23
	VK_HOME       = 0x24
	VK_LEFT       = 0x25
	VK_UP         = 0x26
	VK_RIGHT      = 0x27
	VK_DOWN       = 0x28
	VK_SELECT     = 0x29
	VK_PRINT      = 0x2A
	VK_EXECUTE    = 0x2B
	VK_SNAPSHOT   = 0x2C
	VK_INSERT     = 0x2D
	VK_DELETE     = 0x2E
	VK_HELP       = 0x2F
	VK_0          = 0x30
	VK_1          = 0x31
	VK_2          = 0x32
	VK_3          = 0x33
	VK_4          = 0x34
	VK_5          = 0x35
	VK_6          = 0x36
	VK_7          = 0x37
	VK_8          = 0x38
	VK_9          = 0x39
	// _=	0x3A-40
	VK_A    = 0x41
	VK_B    = 0x42
	VK_C    = 0x43
	VK_D    = 0x44
	VK_E    = 0x45
	VK_F    = 0x46
	VK_G    = 0x47
	VK_H    = 0x48
	VK_I    = 0x49
	VK_J    = 0x4A
	VK_K    = 0x4B
	VK_L    = 0x4C
	VK_M    = 0x4D
	VK_N    = 0x4E
	VK_O    = 0x4F
	VK_P    = 0x50
	VK_Q    = 0x51
	VK_R    = 0x52
	VK_S    = 0x53
	VK_T    = 0x54
	VK_U    = 0x55
	VK_V    = 0x56
	VK_W    = 0x57
	VK_X    = 0x58
	VK_Y    = 0x59
	VK_Z    = 0x5A
	VK_LWIN = 0x5B
	VK_RWIN = 0x5C
	VK_APPS = 0x5D
	// _=	0x5E
	VK_SLEEP     = 0x5F
	VK_NUMPAD0   = 0x60
	VK_NUMPAD1   = 0x61
	VK_NUMPAD2   = 0x62
	VK_NUMPAD3   = 0x63
	VK_NUMPAD4   = 0x64
	VK_NUMPAD5   = 0x65
	VK_NUMPAD6   = 0x66
	VK_NUMPAD7   = 0x67
	VK_NUMPAD8   = 0x68
	VK_NUMPAD9   = 0x69
	VK_MULTIPLY  = 0x6A
	VK_ADD       = 0x6B
	VK_SEPARATOR = 0x6C
	VK_SUBTRACT  = 0x6D
	VK_DECIMAL   = 0x6E
	VK_DIVIDE    = 0x6F
	VK_F1        = 0x70
	VK_F2        = 0x71
	VK_F3        = 0x72
	VK_F4        = 0x73
	VK_F5        = 0x74
	VK_F6        = 0x75
	VK_F7        = 0x76
	VK_F8        = 0x77
	VK_F9        = 0x78
	VK_F10       = 0x79
	VK_F11       = 0x7A
	VK_F12       = 0x7B
	VK_F13       = 0x7C
	VK_F14       = 0x7D
	VK_F15       = 0x7E
	VK_F16       = 0x7F
	VK_F17       = 0x80
	VK_F18       = 0x81
	VK_F19       = 0x82
	VK_F20       = 0x83
	VK_F21       = 0x84
	VK_F22       = 0x85
	VK_F23       = 0x86
	VK_F24       = 0x87
	// _=	0x88-8F
	VK_NUMLOCK = 0x90
	VK_SCROLL  = 0x91
	// -	0x92-96	OEM 特有
	// -	0x97-9F	未分配
	VK_LSHIFT              = 0xA0
	VK_RSHIFT              = 0xA1
	VK_LCONTROL            = 0xA2
	VK_RCONTROL            = 0xA3
	VK_LMENU               = 0xA4
	VK_RMENU               = 0xA5
	VK_BROWSER_BACK        = 0xA6
	VK_BROWSER_FORWARD     = 0xA7
	VK_BROWSER_REFRESH     = 0xA8
	VK_BROWSER_STOP        = 0xA9
	VK_BROWSER_SEARCH      = 0xAA
	VK_BROWSER_FAVORITES   = 0xAB
	VK_BROWSER_HOME        = 0xAC
	VK_VOLUME_MUTE         = 0xAD
	VK_VOLUME_DOWN         = 0xAE
	VK_VOLUME_UP           = 0xAF
	VK_MEDIA_NEXT_TRACK    = 0xB0
	VK_MEDIA_PREV_TRACK    = 0xB1
	VK_MEDIA_STOP          = 0xB2
	VK_MEDIA_PLAY_PAUSE    = 0xB3
	VK_LAUNCH_MAIL         = 0xB4
	VK_LAUNCH_MEDIA_SELECT = 0xB5
	VK_LAUNCH_APP1         = 0xB6
	VK_LAUNCH_APP2         = 0xB7
	// -	0xB8-B9	预留
	VK_OEM_1      = 0xBA
	VK_OEM_PLUS   = 0xBB
	VK_OEM_COMMA  = 0xBC
	VK_OEM_MINUS  = 0xBD
	VK_OEM_PERIOD = 0xBE
	VK_OEM_2      = 0xBF
	VK_OEM_3      = 0xC0
	// -	0xC1-DA	保留
	VK_OEM_4 = 0xDB
	VK_OEM_5 = 0xDC
	VK_OEM_6 = 0xDD
	VK_OEM_7 = 0xDE
	VK_OEM_8 = 0xDF
	// -	0xE0	预留
	// -	0xE1	OEM 特有
	VK_OEM_102 = 0xE2
	// -	0xE3-E4	OEM 特有
	VK_PROCESSKEY = 0xE5
	// -	0xE6	OEM 特有
	VK_PACKET = 0xE7
	// -	0xE8	未分配
	// -	0xE9-F5	OEM 特有
	VK_ATTN      = 0xF6
	VK_CRSEL     = 0xF7
	VK_EXSEL     = 0xF8
	VK_EREOF     = 0xF9
	VK_PLAY      = 0xFA
	VK_ZOOM      = 0xFB
	VK_NONAME    = 0xFC
	VK_PA1       = 0xFD
	VK_OEM_CLEAR = 0xFE
)

type KbDllHookStruct struct {
	VkCode      uintptr
	ScanCode    uintptr
	Flags       uintptr
	Time        uintptr
	DwExtraInfo uintptr
}
