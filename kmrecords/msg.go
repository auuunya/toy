package kmrecords

type UINT_16 = uint16
type MSG struct {
	Hwnd     uintptr
	Message  uint32
	WParam   uintptr
	LParam   uintptr
	Time     uint32
	Pt       *Point
	LPrivate uint32
}

var (
	QS_KEY            UINT_16 = 0x0001
	QS_MOUSEMOVE      UINT_16 = 0x0002
	QS_MOUSEBUTTON    UINT_16 = 0x0004
	QS_POSTMESSAGE    UINT_16 = 0x0008
	QS_TIMER          UINT_16 = 0x0010
	QS_PAINT          UINT_16 = 0x0020
	QS_SENDMESSAGE    UINT_16 = 0x0040
	QS_HOTKEY         UINT_16 = 0x0080
	QS_ALLPOSTMESSAGE UINT_16 = 0x0100
	QS_RAWINPUT       UINT_16 = 0x0400
	QS_TOUCH          UINT_16 = 0x0800
	QS_POINTER        UINT_16 = 0x1000
	QS_MOUSE                  = QS_MOUSEMOVE | QS_MOUSEBUTTON
	QS_INPUT                  = QS_MOUSE | QS_KEY | QS_RAWINPUT | QS_TOUCH | QS_POINTER
	QS_ALLEVENTS              = QS_INPUT | QS_POSTMESSAGE | QS_TIMER | QS_PAINT | QS_HOTKEY
	QS_ALLINPUT               = QS_INPUT | QS_POSTMESSAGE | QS_TIMER | QS_PAINT | QS_HOTKEY | QS_SENDMESSAGE
)

var (
	PM_NOREMOVE       UINT_16 = 0x0000
	PM_REMOVE         UINT_16 = 0x0001
	PM_NOYIELD        UINT_16 = 0x0002
	PM_QS_INPUT               = QS_INPUT << 16
	PM_QS_PAINT               = QS_PAINT << 16
	PM_QS_POSTMESSAGE         = (QS_POSTMESSAGE | QS_HOTKEY | QS_TIMER) << 16
	PM_QS_SENDMESSAGE         = QS_SENDMESSAGE << 16
)
