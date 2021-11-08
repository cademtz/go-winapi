package winapi

import "unsafe"

type HMODULE uintptr

const (
	IMAGE_ORDINAL_FLAG64 = 0x8000000000000000
	IMAGE_ORDINAL_FLAG32 = 0x80000000

	MAX_PATH = uint32(261)

	_SIZEOF_PID     = unsafe.Sizeof(uint32(0))
	_SIZEOF_HMODULE = unsafe.Sizeof(HMODULE(0))
)
