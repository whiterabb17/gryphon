package shellcode

import (
	"log"
)

func injectIntoProcess(process, args, data string) bool {
	log.Println("Not currently implemented in Linux")
	return false
}

func injectIntoProcessEarlyBird(process, args, data string) bool {
	log.Println("Not currently implemented in Linux")
	return false
}

func syscallInjectShellcode(data string) bool {
	log.Println("Not currently implemented in Linux")
	return false
}

func createThreadInject(data string) bool {
	log.Println("Not currently implemented in Linux")
	return false
}
