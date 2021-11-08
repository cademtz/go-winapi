package winapi

import (
	"syscall"
	"unsafe"
)

const (
	PAGE_NOACCESS                   = 0x01
	PAGE_READONLY                   = 0x02
	PAGE_READWRITE                  = 0x04
	PAGE_WRITECOPY                  = 0x08
	PAGE_EXECUTE                    = 0x10
	PAGE_EXECUTE_READ               = 0x20
	PAGE_EXECUTE_READWRITE          = 0x40
	PAGE_EXECUTE_WRITECOPY          = 0x80
	PAGE_GUARD                      = 0x100
	PAGE_NOCACHE                    = 0x200
	PAGE_WRITECOMBINE               = 0x400
	PAGE_GRAPHICS_NOACCESS          = 0x0800
	PAGE_GRAPHICS_READONLY          = 0x1000
	PAGE_GRAPHICS_READWRITE         = 0x2000
	PAGE_GRAPHICS_EXECUTE           = 0x4000
	PAGE_GRAPHICS_EXECUTE_READ      = 0x8000
	PAGE_GRAPHICS_EXECUTE_READWRITE = 0x10000
	PAGE_GRAPHICS_COHERENT          = 0x20000
	PAGE_GRAPHICS_NOCACHE           = 0x40000
	PAGE_ENCLAVE_THREAD_CONTROL     = 0x80000000
	PAGE_REVERT_TO_FILE_MAP         = 0x80000000
	PAGE_TARGETS_NO_UPDATE          = 0x40000000
	PAGE_TARGETS_INVALID            = 0x40000000
	PAGE_ENCLAVE_UNVALIDATED        = 0x20000000
	PAGE_ENCLAVE_MASK               = 0x10000000
	PAGE_ENCLAVE_DECOMMIT           = (PAGE_ENCLAVE_MASK | 0)
	PAGE_ENCLAVE_SS_FIRST           = (PAGE_ENCLAVE_MASK | 1)
	PAGE_ENCLAVE_SS_REST            = (PAGE_ENCLAVE_MASK | 2)
	MEM_COMMIT                      = 0x00001000
	MEM_RESERVE                     = 0x00002000
	MEM_REPLACE_PLACEHOLDER         = 0x00004000
	MEM_RESERVE_PLACEHOLDER         = 0x00040000
	MEM_RESET                       = 0x00080000
	MEM_TOP_DOWN                    = 0x00100000
	MEM_WRITE_WATCH                 = 0x00200000
	MEM_PHYSICAL                    = 0x00400000
	MEM_ROTATE                      = 0x00800000
	MEM_DIFFERENT_IMAGE_BASE_OK     = 0x00800000
	MEM_RESET_UNDO                  = 0x01000000
	MEM_LARGE_PAGES                 = 0x20000000
	MEM_4MB_PAGES                   = 0x80000000
	MEM_64K_PAGES                   = (MEM_LARGE_PAGES | MEM_PHYSICAL)
	MEM_UNMAP_WITH_TRANSIENT_BOOST  = 0x00000001
	MEM_COALESCE_PLACEHOLDERS       = 0x00000001
	MEM_PRESERVE_PLACEHOLDER        = 0x00000002
	MEM_DECOMMIT                    = 0x00004000
	MEM_RELEASE                     = 0x00008000
	MEM_FREE                        = 0x00010000

	STANDARD_RIGHTS_REQUIRED = uint32(0x000F0000)
	SYNCHRONIZE              = uint32(0x00100000)

	PROCESS_TERMINATE                 = uint32(0x0001)
	PROCESS_CREATE_THREAD             = uint32(0x0002)
	PROCESS_SET_SESSIONID             = uint32(0x0004)
	PROCESS_VM_OPERATION              = uint32(0x0008)
	PROCESS_VM_READ                   = uint32(0x0010)
	PROCESS_VM_WRITE                  = uint32(0x0020)
	PROCESS_DUP_HANDLE                = uint32(0x0040)
	PROCESS_CREATE_PROCESS            = uint32(0x0080)
	PROCESS_SET_QUOTA                 = uint32(0x0100)
	PROCESS_SET_INFORMATION           = uint32(0x0200)
	PROCESS_QUERY_INFORMATION         = uint32(0x0400)
	PROCESS_SUSPEND_RESUME            = uint32(0x0800)
	PROCESS_QUERY_LIMITED_INFORMATION = uint32(0x1000)
	PROCESS_SET_LIMITED_INFORMATION   = uint32(0x2000)
	PROCESS_ALL_ACCESS                = uint32(STANDARD_RIGHTS_REQUIRED | SYNCHRONIZE | 0xFFFF)

	HEAP_NO_SERIALIZE          = uint32(0x00000001)
	HEAP_GENERATE_EXCEPTIONS   = uint32(0x00000004)
	HEAP_ZERO_MEMORY           = uint32(0x00000008)
	HEAP_CREATE_ENABLE_EXECUTE = uint32(0x00040000)
)

var (
	Kernel32 = syscall.MustLoadDLL("kernel32.dll")
	_GetProcAddress = Kernel32.MustFindProc("GetProcAddress")
	_VirtualAllocEx = Kernel32.MustFindProc("VirtualAllocEx")
	_VirtualAlloc = Kernel32.MustFindProc("VirtualAlloc")
	_ReadProcessMemory = Kernel32.MustFindProc("ReadProcessMemory")
	_WriteProcessMemory = Kernel32.MustFindProc("WriteProcessMemory")
	_CreateRemoteThread = Kernel32.MustFindProc("CreateRemoteThread")
	_ExitThread = Kernel32.MustFindProc("ExitThread")
	_GetModuleHandleW = Kernel32.MustFindProc("GetModuleHandleW")
	_AllocConsole = Kernel32.MustFindProc("AllocConsole")
	_GetStdHandle = Kernel32.MustFindProc("GetStdHandle")
	_VirtualProtect = Kernel32.MustFindProc("VirtualProtect")
	_HeapCreate = Kernel32.MustFindProc("HeapCreate")
	_HeapDestroy = Kernel32.MustFindProc("HeapDestroy")
	_HeapAlloc = Kernel32.MustFindProc("HeapAlloc")
	_HeapReAlloc = Kernel32.MustFindProc("HeapReAlloc")
	_HeapFree = Kernel32.MustFindProc("HeapFree")
)

func GetProcAddress(hModule syscall.Handle, Ordinal uint16) uintptr {
	addr, _, _ := _GetProcAddress.Call(uintptr(hModule), uintptr(Ordinal))
	return addr
}

func VirtualAllocEx(hProcess syscall.Handle, lpAddress uintptr, dwSize uintptr, flAllocationType uint32, flProtect uint32) uintptr {
	addr, _, _ := _VirtualAllocEx.Call(
		uintptr(hProcess),
		lpAddress,
		dwSize,
		uintptr(flAllocationType),
		uintptr(flProtect),
	)
	return addr
}

func VirtualAlloc(lpAddress uintptr, dwSize uintptr, flAllocationType uint32, flProtect uint32) uintptr {
	addr, _, _ := _VirtualAlloc.Call(
		lpAddress,
		dwSize,
		uintptr(flAllocationType),
		uintptr(flProtect),
	)
	return addr
}

func ReadProcessMemory(hProcess syscall.Handle, lpBaseAddress uintptr, Buffer []byte) bool {
	ok, _, _ := _ReadProcessMemory.Call(
		uintptr(hProcess),
		lpBaseAddress,
		uintptr(unsafe.Pointer(&Buffer[0])),
		uintptr(len(Buffer)),
		0,
	)
	return ok != 0
}

func WriteProcessMemory(hProcess syscall.Handle, lpBaseAddress uintptr, Buffer []byte) bool {
	ok, _, _ := _WriteProcessMemory.Call(
		uintptr(hProcess),
		lpBaseAddress,
		uintptr(unsafe.Pointer(&Buffer[0])),
		uintptr(len(Buffer)),
		0,
	)
	return ok != 0
}

func CreateRemoteThread(hProcess syscall.Handle, dwStackSize uintptr, lpStartAddress uintptr, lpParameter uintptr, dwCreationFlags uint32) syscall.Handle {
	handle, _, _ := _CreateRemoteThread.Call(
		uintptr(hProcess),
		0,
		dwStackSize,
		lpStartAddress,
		lpParameter,
		uintptr(dwCreationFlags),
		0,
	)
	return syscall.Handle(handle)
}

func ExitThread(dwExitCode uint32) {
	_ExitThread.Call(uintptr(dwExitCode))
}

func GetModuleHandle(Name string) syscall.Handle {
	namePtr, err := syscall.UTF16PtrFromString(Name)
	if err != nil { panic(err) }

	hMod, _, _ := _GetModuleHandleW.Call(uintptr(unsafe.Pointer(namePtr)))
	return syscall.Handle(hMod)
}

func AllocConsole() bool {
	ok, _, _ := _AllocConsole.Call()
	return ok != 0
}

func GetStdHandle(nStdHandle int) syscall.Handle {
	handle, _, _ := _GetStdHandle.Call(uintptr(nStdHandle))
	return syscall.Handle(handle)
}

func VirtualProtect(lpAddress uintptr, dwSize uint32, flNewProtect uint32, lpflOldProtect *uint32) bool {
	ok, _, _ := _VirtualProtect.Call(
		lpAddress,
		uintptr(dwSize),
		uintptr(flNewProtect),
		uintptr(unsafe.Pointer(lpflOldProtect)),
	)

	return ok != 0
}

func HeapCreate(flOptions uint32, dwInitialSize uintptr, dwMaximumSize uintptr) syscall.Handle {
	handle, _, _ := _HeapCreate.Call(uintptr(flOptions), dwInitialSize, dwMaximumSize)
	return syscall.Handle(handle)
}

func HeapAlloc(hHeap syscall.Handle, dwFlags uint32, dwBytes uintptr) uintptr {
	newmem, _, _ := _HeapAlloc.Call(uintptr(hHeap), uintptr(dwFlags), dwBytes)
	return newmem
}

func HeapReAlloc(hHeap syscall.Handle, dwFlags uint32, lpMem uintptr, dwBytes uintptr) uintptr {
	newmem, _, _ := _HeapReAlloc.Call(uintptr(hHeap), uintptr(dwFlags), lpMem, dwBytes)
	return newmem
}

func HeapDestroy(hHeap syscall.Handle) bool {
	ok, _, _ := _HeapDestroy.Call(uintptr(hHeap))
	return ok != 0
}

func HeapFree(hHeap syscall.Handle, dwFlags uint32, lpMem uintptr) bool {
	ok, _, _ := _HeapFree.Call(uintptr(hHeap), uintptr(dwFlags), lpMem)
	return ok != 0
}
