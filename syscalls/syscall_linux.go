package syscalls

func SyscallExecute(Shellcode []byte) bool {

	// Addr, _, _ := variables.VirtualAlloc.Call(0, uintptr(len(Shellcode)), variables.MEM_RESERVE|variables.MEM_COMMIT, variables.PAGE_EXECUTE_READWRITE)

	// AddrPtr := (*[990000]byte)(unsafe.Pointer(Addr))

	// for i := 0; i < len(Shellcode); i++ {
	// 	AddrPtr[i] = Shellcode[i]
	// }

	// go syscall.Syscall(Addr, 0, 0, 0, 0)
	return false
}
