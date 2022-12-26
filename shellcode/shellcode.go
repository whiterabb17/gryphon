package shellcode

func InjectIntoProcess(process, args, data string) bool {
	return injectIntoProcess(process, args, data)
}

func InjectIntoProcessEarlyBird(process, args, data string) bool {
	return injectIntoProcessEarlyBird(process, args, data)
}

func SyscallInjectShellcode(data string) bool {
	return syscallInjectShellcode(data)
}

func CreateThreadInject(data string) bool {
	return createThreadInject(data)
}
