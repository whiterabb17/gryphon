package deepfire

import (
	goLift "github.com/whiterabb17/goLift"
)

func escalate(path string) string {
	err := goLift.WEscalate(path)
	return err.Error()
}
