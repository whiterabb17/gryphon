package deepfire

import (
	goLift "github.com/whiterabb17/goLift"
)

func escalate(path string) string {
	err := goLift.NEscalate(path)
	return err.Error()
}
