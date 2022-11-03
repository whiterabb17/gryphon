package gryphon

import (
	"log"
)

func bypassAV() {
	log.Println("Not implemented on Darwin yet")
}

func detectDebugProc() bool { return false }

func checkForProc(variable string) bool { return false }
