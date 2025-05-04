package winocr

import (
	"os"
	"path/filepath"
	"sync"
	"unsafe"

	"golang.org/x/sys/windows"
)

const (
	modelKey = `kj)TGtrK>f]b[Piow.gU+nC@s""""""4`
)

var (
	ocrPath    = "."
	oneocr     *windows.LazyDLL
	oneocrOnce sync.Once
)

func SetOcrDllPath(path string) {
	validateOcrDllPath(path)
	oneocrOnce.Do(func() {
		windows.SetDllDirectory(path)
		oneocr = windows.NewLazyDLL(filepath.Join(path, "oneocr.dll"))
		if oneocr.Handle() == 0 {
			panic("failed to load oneocr.dll")
		}
		ocrPath = path
	})
}

func GetOcrDllPath() string {
	return ocrPath
}

func validateOcrDllPath(path string) {
	for _, file := range []string{"oneocr.dll", "oneocr.onemodel", "onnxruntime.dll"} {
		if _, err := os.Stat(filepath.Join(path, file)); os.IsNotExist(err) {
			panic(file + " not found")
		}
	}
}

func ansi2String(ptr *byte) string {
	if ptr == nil {
		return ""
	}

	length := 0

	for {
		chunk := *(*uint64)(unsafe.Add(unsafe.Pointer(ptr), length))
		if (chunk-0x0101010101010101) & ^chunk & 0x8080808080808080 != 0 {
			break
		}
		length += 8
	}

	for *(*byte)(unsafe.Add(unsafe.Pointer(ptr), length)) != 0 {
		length++
	}

	return unsafe.String(ptr, length)
}
