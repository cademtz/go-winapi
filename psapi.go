package winapi

import (
	"errors"
	"fmt"
	"syscall"
	"unsafe"
)

const (
	LIST_MODULES_32BIT   uint32 = 0x01
	LIST_MODULES_64BIT   uint32 = 0x02
	LIST_MODULES_ALL     uint32 = 0x03
	LIST_MODULES_DFEAULT uint32 = 0x0
)

var (
	Psapi = syscall.MustLoadDLL("psapi.dll")
	_EnumProcesses = Psapi.MustFindProc("EnumProcesses")
	_EnumProcessModulesEx = Psapi.MustFindProc("EnumProcessModulesEx")
	_GetModuleBaseName = Psapi.MustFindProc("GetModuleBaseNameW")
)

func EnumProcesses() ([]uint32, error) {
	var cbNeeded uint32
	pidList := make([]uint32, 1024)
	
	ok, _, _ := _EnumProcesses.Call(uintptr(unsafe.Pointer(&pidList[0])), uintptr(len(pidList)) * _SIZEOF_PID, uintptr(unsafe.Pointer(&cbNeeded)))
	for ok != 0 && uintptr(cbNeeded) / _SIZEOF_PID >= uintptr(len(pidList)) {
		pidList = make([]uint32, int(float32(len(pidList)) * 1.5)) // Go.... plz...
		ok, _, _ = _EnumProcesses.Call(uintptr(unsafe.Pointer(&pidList[0])), uintptr(len(pidList)) * _SIZEOF_PID, uintptr(unsafe.Pointer(&cbNeeded)))
	}
	if ok == 0 { return nil, errors.New("EnumProcesses failed") }

	return pidList[:uintptr(cbNeeded) / _SIZEOF_PID], nil
}

func EnumProcessModulesEx(hProc syscall.Handle, dwFilterFlag uint32) ([]HMODULE, error) {
	var cbNeeded uint32
	modList := make([]HMODULE, 128)
	var ok uintptr
	var err error
	
	ok, _, err = _EnumProcessModulesEx.Call(
		uintptr(hProc),
		uintptr(unsafe.Pointer(&modList[0])),
		uintptr(len(modList) * int(_SIZEOF_HMODULE)),
		uintptr(unsafe.Pointer(&cbNeeded)),
		uintptr(dwFilterFlag),
	)
	for ok != 0 && uintptr(cbNeeded) / _SIZEOF_PID >= uintptr(len(modList)) {
		modList = make([]HMODULE, int(float32(len(modList)) * 1.5)) // Go.... plz...
		ok, _, err = _EnumProcessModulesEx.Call(
			uintptr(hProc),
			uintptr(unsafe.Pointer(&modList[0])),
			uintptr(len(modList) * int(_SIZEOF_HMODULE)),
			uintptr(unsafe.Pointer(&cbNeeded)),
			uintptr(dwFilterFlag),
		)
	}
	if ok == 0 { return nil, fmt.Errorf("EnumProcessModules failed (%v)", err) }

	return modList[:uintptr(cbNeeded) / _SIZEOF_PID], nil
}

func GetModuleBaseNameW(hProc syscall.Handle, hMod HMODULE) (string, error) {
	lpBaseName := make([]uint16, MAX_PATH)

	len, _, _ := _GetModuleBaseName.Call(uintptr(hProc), uintptr(hMod), uintptr(unsafe.Pointer(&lpBaseName[0])), uintptr(len(lpBaseName)))

	if len == 0 { return "", errors.New("GetModuleBaseNameW failed") }

	return syscall.UTF16ToString(lpBaseName),nil
}
