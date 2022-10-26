package deepfire

import "os/exec"

func cmdOut(command string) (string, error) {
	cmd := exec.Command("/bin/zsh", "-c", command)
	output, err := cmd.CombinedOutput()
	out := string(output)
	return out, err
}
