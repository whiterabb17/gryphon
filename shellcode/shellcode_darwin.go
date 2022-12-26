package shellcode

import (
	"log"
)

func injectIntoProcess(process, args, data string) bool {
	log.Println("Not currently implemented in Darwin")
	return false
}

func injectIntoProcessEarlyBird(process, args, data string) bool {
	log.Println("Not currently implemented in Darwin")
	return false
}

func syscallInjectShellcode(data string) bool {
	log.Println("Not currently implemented in Darwin")
	return false
}

func createThreadInject(data string) bool {
	log.Println("Not currently implemented in Darwin")
	return false
}
