package deepfire

import (
	"debug/pe"
	"unsafe"
)

type baseRelocEntry uint16
type IMAGE_REL_BASED uint16
type usp = unsafe.Pointer
type size_t = int
type Row = []byte
type Ptr = unsafe.Pointer
type PFunc = func([]byte, Ptr, *TypeInfo, int) Ptr
type BASE_RELOCATION_ENTRY uint16

type IMAGE_DOS_HEADER struct {
	E_magic    uint16
	E_cblp     uint16
	E_cp       uint16
	E_crlc     uint16
	E_cparhdr  uint16
	E_minalloc uint16
	E_maxalloc uint16
	E_ss       uint16
	E_sp       uint16
	E_csum     uint16
	E_ip       uint16
	E_cs       uint16
	E_lfarlc   uint16
	E_ovno     uint16
	E_res      [4]uint16
	E_oemid    uint16
	E_oeminfo  uint16
	E_res2     [10]uint16
	E_lfanew   int32
}

type IMAGE_NT_HEADERS struct {
	Signature      uint32
	FileHeader     pe.FileHeader
	OptionalHeader pe.OptionalHeader32
}

type IMAGE_NT_HEADERS64 struct {
	Signature      uint32
	FileHeader     pe.FileHeader
	OptionalHeader pe.OptionalHeader64
}

type IMAGE_BASE_RELOCATION struct {
	VirtualAddress uint32
	SizeOfBlock    uint32
}

type Container struct {
	ti   *TypeInfo
	Rows []Row
}

type m128a struct {
	low  uint64
	high int64
}

/*
type winFileTime struct {
	LowDateTime  uint32
	HighDateTime uint32
}

type DATA_BLOB struct {
	cbData uint32
	pbData *byte
}

*/
type WOW64_FLOATING_SAVE_AREA struct {
	ControlWord   uint32
	StatusWord    uint32
	TagWord       uint32
	ErrorOffset   uint32
	ErrorSelector uint32
	DataOffset    uint32
	DataSelector  uint32
	RegisterArea  [80]byte
	Cr0NpxState   uint32
}

type WOW64_CONTEXT struct {
	ContextFlags      uint32
	Dr0               uint32
	Dr1               uint32
	Dr2               uint32
	Dr3               uint32
	Dr6               uint32
	Dr7               uint32
	FloatSave         WOW64_FLOATING_SAVE_AREA
	SegGs             uint32
	SegFs             uint32
	SegEs             uint32
	SegDs             uint32
	Edi               uint32
	Esi               uint32
	Ebx               uint32
	Edx               uint32
	Ecx               uint32
	Eax               uint32
	Ebp               uint32
	Eip               uint32
	SegCs             uint32
	EFlags            uint32
	Esp               uint32
	SegSs             uint32
	ExtendedRegisters [512]byte
}

type CONTEXT struct {
	p1home               uint64
	p2home               uint64
	p3home               uint64
	p4home               uint64
	p5home               uint64
	p6home               uint64
	contextflags         uint32
	mxcsr                uint32
	segcs                uint16
	segds                uint16
	seges                uint16
	segfs                uint16
	seggs                uint16
	segss                uint16
	eflags               uint32
	dr0                  uint64
	dr1                  uint64
	dr2                  uint64
	dr3                  uint64
	dr6                  uint64
	dr7                  uint64
	rax                  uint64
	rcx                  uint64
	rdx                  uint64
	rbx                  uint64
	rsp                  uint64
	rbp                  uint64
	rsi                  uint64
	rdi                  uint64
	r8                   uint64
	r9                   uint64
	r10                  uint64
	r11                  uint64
	r12                  uint64
	r13                  uint64
	r14                  uint64
	r15                  uint64
	rip                  uint64
	anon0                [512]byte
	vectorregister       [26]m128a
	vectorcontrol        uint64
	debugcontrol         uint64
	lastbranchtorip      uint64
	lastbranchfromrip    uint64
	lastexceptiontorip   uint64
	lastexceptionfromrip uint64
}

const (
	EsSystemRequired = 0x00000001
	EsContinuous     = 0x80000000

	STRING = iota
	INT
	CONTEXT_AMD64                                    = 0x100000
	CONTEXT_INTEGER                                  = (CONTEXT_AMD64 | 0x2)
	CREATE_SUSPENDED                                 = 0x00000004
	MEM_RELEASE                                      = 0x8000
	IMAGE_NT_OPTIONAL_HDR32_MAGIC                    = 0x10b
	IMAGE_NT_OPTIONAL_HDR64_MAGIC                    = 0x20b
	IMAGE_DOS_SIGNATURE                              = 0x5A4D
	IMAGE_NT_SIGNATURE                               = 0x00004550
	IMAGE_DIRECTORY_ENTRY_BASERELOC                  = 5
	IMAGE_NUMBEROF_DIRECTORY_ENTRIES                 = 16
	IMAGE_SIZEOF_SECTION_HEADER                      = 40
	IMAGE_REL_BASED_ABSOLUTE         IMAGE_REL_BASED = 0 //The base relocation is skipped. This type can be used to pad a block.
	IMAGE_REL_BASED_HIGH             IMAGE_REL_BASED = 1 //The base relocation adds the high 16 bits of the difference to the 16-bit field at offset. The 16-bit field represents the high value of a 32-bit word.
	IMAGE_REL_BASED_LOW              IMAGE_REL_BASED = 2 //The base relocation adds the low 16 bits of the difference to the 16-bit field at offset. The 16-bit field represents the low half of a 32-bit word.
	IMAGE_REL_BASED_HIGHLOW          IMAGE_REL_BASED = 3 //The base relocation applies all 32 bits of the difference to the 32-bit field at offset.
	IMAGE_REL_BASED_HIGHADJ          IMAGE_REL_BASED = 4 //The base relocation adds the high 16 bits of the difference to the 16-bit field at offset. The 16-bit field represents the high value of a 32-bit word. The low 16 bits of the 32-bit value are stored in the 16-bit word that follows this base relocation. This means that this base relocation occupies two slots.
	IMAGE_REL_BASED_MIPS_JMPADDR     IMAGE_REL_BASED = 5 //The relocation interpretation is dependent on the machine type.When the machine type is MIPS, the base relocation applies to a MIPS jump instruction.
	IMAGE_REL_BASED_ARM_MOV32        IMAGE_REL_BASED = 5 //This relocation is meaningful only when the machine type is ARM or Thumb. The base relocation applies the 32-bit address of a symbol across a consecutive MOVW/MOVT instruction pair.
	IMAGE_REL_BASED_RISCV_HIGH20     IMAGE_REL_BASED = 5 //This relocation is only meaningful when the machine type is RISC-V. The base relocation applies to the high 20 bits of a 32-bit absolute address.
	IMAGE_REL_BASED_THUMB_MOV32      IMAGE_REL_BASED = 7 //This relocation is meaningful only when the machine type is Thumb. The base relocation applies the 32-bit address of a symbol to a consecutive MOVW/MOVT instruction pair.
	IMAGE_REL_BASED_RISCV_LOW12I     IMAGE_REL_BASED = 7 //This relocation is only meaningful when the machine type is RISC-V. The base relocation applies to the low 12 bits of a 32-bit absolute address formed in RISC-V I-type instruction format.
	IMAGE_REL_BASED_RISCV_LOW12S     IMAGE_REL_BASED = 8 //This relocation is only meaningful when the machine type is RISC-V. The base relocation applies to the low 12 bits of a 32-bit absolute address formed in RISC-V S-type instruction format.
	IMAGE_REL_BASED_MIPS_JMPADDR16   IMAGE_REL_BASED = 9 //The relocation is only meaningful when the machine type is MIPS. The base relocation applies to a MIPS16 jump instruction.
	IMAGE_REL_BASED_DIR64            IMAGE_REL_BASED = 10
	PAGE_READWRITE                                   = 0x04
	PAGE_EXECUTE_READ                                = 0x20
	RELOC_32BIT_FIELD                                = 3
	RELOC_64BIT_FIELD                                = 0xA
)

type TypeInfo struct {
	fields int
	Type   []uint
	Pos    []int
	Size   []int
	Offset []uintptr
	Save   []PFunc
	Dump   []PFunc
}

type dataBlob struct {
	cbData uint32
	pbData *byte
}

type IMAGE_OPTIONAL_HEADER64 struct {
	Magic                       uint16
	MajorLinkerVersion          byte
	MinorLinkerVersion          byte
	SizeOfCode                  uint32
	SizeOfInitializedData       uint32
	SizeOfUninitializedData     uint32
	AddressOfEntryPoint         uint32
	BaseOfCode                  uint32
	ImageBase                   uint64
	SectionAlignment            uint32
	FileAlignment               uint32
	MajorOperatingSystemVersion uint16
	MinorOperatingSystemVersion uint16
	MajorImageVersion           uint16
	MinorImageVersion           uint16
	MajorSubsystemVersion       uint16
	MinorSubsystemVersion       uint16
	Win32VersionValue           uint32
	SizeOfImage                 uint32
	SizeOfHeaders               uint32
	CheckSum                    uint32
	Subsystem                   uint16
	DllCharacteristics          uint16
	SizeOfStackReserve          uint64
	SizeOfStackCommit           uint64
	SizeOfHeapReserve           uint64
	SizeOfHeapCommit            uint64
	LoaderFlags                 uint32
	NumberOfRvaAndSizes         uint32
	DataDirectory               uintptr
}

type PEB struct {
	InheritedAddressSpace    byte    // BYTE	0
	ReadImageFileExecOptions byte    // BYTE	1
	BeingDebugged            byte    // BYTE	2
	reserved2                [1]byte // BYTE 3

	Mutant                 uintptr     // BYTE 4
	ImageBaseAddress       uintptr     // BYTE 8
	Ldr                    uintptr     // PPEB_LDR_DATA
	ProcessParameters      uintptr     // PRTL_USER_PROCESS_PARAMETERS
	reserved4              [3]uintptr  // PVOID
	AtlThunkSListPtr       uintptr     // PVOID
	reserved5              uintptr     // PVOID
	reserved6              uint32      // ULONG
	reserved7              uintptr     // PVOID
	reserved8              uint32      // ULONG
	AtlThunkSListPtr32     uint32      // ULONG
	reserved9              [45]uintptr // PVOID
	reserved10             [96]byte    // BYTE
	PostProcessInitRoutine uintptr     // PPS_POST_PROCESS_INIT_ROUTINE
	reserved11             [128]byte   // BYTE
	reserved12             [1]uintptr  // PVOID
	SessionId              uint32      // ULONG
}

// https://github.com/elastic/go-windows/blob/master/ntdll.go#L77
type PROCESS_BASIC_INFORMATION struct {
	reserved1                    uintptr    // PVOID
	PebBaseAddress               uintptr    // PPEB
	reserved2                    [2]uintptr // PVOID
	UniqueProcessId              uintptr    // ULONG_PTR
	InheritedFromUniqueProcessID uintptr    // PVOID
}

type IMAGE_FILE_HEADER struct {
	Machine              uint16
	NumberOfSections     uint16
	TimeDateStamp        uint32
	PointerToSymbolTable uint32
	NumberOfSymbols      uint32
	SizeOfOptionalHeader uint16
	Characteristics      uint16
}

type IMAGE_OPTIONAL_HEADER32 struct {
	Magic                       uint16
	MajorLinkerVersion          byte
	MinorLinkerVersion          byte
	SizeOfCode                  uint32
	SizeOfInitializedData       uint32
	SizeOfUninitializedData     uint32
	AddressOfEntryPoint         uint32
	BaseOfCode                  uint32
	BaseOfData                  uint32 // Different from 64 bit header
	ImageBase                   uint64
	SectionAlignment            uint32
	FileAlignment               uint32
	MajorOperatingSystemVersion uint16
	MinorOperatingSystemVersion uint16
	MajorImageVersion           uint16
	MinorImageVersion           uint16
	MajorSubsystemVersion       uint16
	MinorSubsystemVersion       uint16
	Win32VersionValue           uint32
	SizeOfImage                 uint32
	SizeOfHeaders               uint32
	CheckSum                    uint32
	Subsystem                   uint16
	DllCharacteristics          uint16
	SizeOfStackReserve          uint64
	SizeOfStackCommit           uint64
	SizeOfHeapReserve           uint64
	SizeOfHeapCommit            uint64
	LoaderFlags                 uint32
	NumberOfRvaAndSizes         uint32
	DataDirectory               uintptr
}

type Type int

func (r *BASE_RELOCATION_ENTRY) GetOffset() (_offset uint16) {
	_offset = uint16(*r) & 0x0fff
	return
}

func (r *BASE_RELOCATION_ENTRY) SetOffset(_offset uint16) {
	*r = *r | BASE_RELOCATION_ENTRY(_offset&0x0fff)
}

func (r *BASE_RELOCATION_ENTRY) SetType(_type uint16) {
	*r = *r | BASE_RELOCATION_ENTRY(_type&0xf000)
}

func (r *BASE_RELOCATION_ENTRY) GetType() (_type uint16) {
	_type = (uint16(*r) & 0xf000) >> 12
	return
}

func (b baseRelocEntry) Type() IMAGE_REL_BASED {
	return IMAGE_REL_BASED(uint16(b) >> 12)
}

func (b baseRelocEntry) Offset() uint32 {
	return uint32(uint16(b) & 0x0FFF)
}
