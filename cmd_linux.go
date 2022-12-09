package gryphon

import (
	"errors"
	"os/exec"
)

func cmdOut(command string) (string, error) {
	cmd := exec.Command("/bin/bash", "-c", command)
	output, err := cmd.CombinedOutput()
	out := string(output)
	return out, err
}

func pwsh(command string) (string, error) {
	return "", errors.New("unimplemented on this system")
}
