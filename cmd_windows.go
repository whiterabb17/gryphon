package gryphon

import (
	"os/exec"
	"syscall"
)

func cmdOut(command string) (string, error) {
	cmd := exec.Command("cmd.exe", "/C", command)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	output, err := cmd.CombinedOutput()
	out := string(output)
	return out, err
}

func pwsh(command string) (string, error) {
	cmd := exec.Command("powershell", "-NoLogo", "-Ep", "Bypass", command)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	b, err := cmd.CombinedOutput()
	out := string(b)
	return out, err
}
