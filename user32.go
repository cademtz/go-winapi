package winapi

import (
	"syscall"
	"unsafe"
)

const (
	MB_ABORTRETRYIGNORE  uint32 = 0x00000002
	MB_CANCELTRYCONTINUE uint32 = 0x00000006
	MB_HELP              uint32 = 0x00004000
	MB_OK                uint32 = 0x00000000
	MB_OKCANCEL          uint32 = 0x00000001
	MB_RETRYCANCEL       uint32 = 0x00000005
	MB_YESNO             uint32 = 0x00000004
	MB_YESNOCANCEL       uint32 = 0x00000003

	MB_ICONEXCLAMATION uint32 = 0x00000030
	MB_ICONWARNING     uint32 = 0x00000030
	MB_ICONINFORMATION uint32 = 0x00000040
	MB_ICONASTERISK    uint32 = 0x00000040
	MB_ICONQUESTION    uint32 = 0x00000020
	MB_ICONSTOP        uint32 = 0x00000010
	MB_ICONERROR       uint32 = 0x00000010
	MB_ICONHAND        uint32 = 0x00000010

	MB_DEFBUTTON1 uint32 = 0x00000000
	MB_DEFBUTTON2 uint32 = 0x00000100
	MB_DEFBUTTON3 uint32 = 0x00000200
	MB_DEFBUTTON4 uint32 = 0x00000300

	MB_APPLMODAL   uint32 = 0x00000000
	MB_SYSTEMMODAL uint32 = 0x00001000
	MB_TASKMODAL   uint32 = 0x00002000

	IDOK       = 1
	IDCANCEL   = 2
	IDABORT    = 3
	IDRETRY    = 4
	IDIGNORE   = 5
	IDYES      = 6
	IDNO       = 7
	IDTRYAGAIN = 10
	IDCONTINUE = 11
)

var (
	User32       = syscall.MustLoadDLL("user32.dll")
	_MessageBoxW = User32.MustFindProc("MessageBoxW")
)

func MessageBox(Msg, Title string, Buttons uint32) int {
	szTitle, err := syscall.UTF16PtrFromString(Title)
	if err != nil { panic(err) }
	
	szMsg, err := syscall.UTF16PtrFromString(Msg)
	if err != nil { panic(err) }

	result, _, _ := _MessageBoxW.Call(
		0,
		uintptr(unsafe.Pointer(szMsg)),
		uintptr(unsafe.Pointer(szTitle)),
		uintptr(Buttons),
	)
	return int(result)
}
