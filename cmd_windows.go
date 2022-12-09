package gryphon

import (
	"os/exec"
	"syscall"
)

func cmdOut(command string) (string, error) {
	cmd := exec.Command("cmd", command)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	output, err := cmd.CombinedOutput()
	out := string(output)
	return out, err
}

/*
func pwsh(command string) (string, error) {
	binary, _ := exec.LookPath("powershell")
	_cmd := exec.Command(binary, fmt.Sprintf(`PowerShell -WindowStyle hidden -encodedCommand "%s"`, command))
	output, err := _cmd.CombinedOutput()
	return string(output), err
}
*/
