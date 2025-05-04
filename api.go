package winocr

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/windows"
)

func CreateOcrInitOptions() uintptr {
	var handle uintptr
	proc := oneocr.NewProc("CreateOcrInitOptions")
	proc.Call(
		uintptr(unsafe.Pointer(&handle)),
	)
	return handle
}

func OcrInitOptionsSetUseModelDelayLoad(initOpts uintptr, enable bool) {
	enabled := 0
	if enable {
		enabled = 1
	}
	proc := oneocr.NewProc("OcrInitOptionsSetUseModelDelayLoad")
	proc.Call(
		initOpts,
		uintptr(enabled),
	)
}

func CreateOcrPipeline(initOpts uintptr, modelPath, modelKey string) uintptr {
	var handle uintptr
	modelPathPtr, _ := windows.BytePtrFromString(modelPath)
	modelKeyPtr, _ := windows.BytePtrFromString(modelKey)
	proc := oneocr.NewProc("CreateOcrPipeline")
	proc.Call(
		uintptr(unsafe.Pointer(modelPathPtr)),
		uintptr(unsafe.Pointer(modelKeyPtr)),
		initOpts,
		uintptr(unsafe.Pointer(&handle)),
	)
	return handle
}

func CreateOcrProcessOptions() uintptr {
	var handle uintptr
	proc := oneocr.NewProc("CreateOcrProcessOptions")
	proc.Call(
		uintptr(unsafe.Pointer(&handle)),
	)
	return handle
}

func OcrProcessOptionsGetMaxRecognitionLineCount(processOpts uintptr) int {
	var count int
	proc := oneocr.NewProc("OcrProcessOptionsGetMaxRecognitionLineCount")
	proc.Call(
		processOpts,
		uintptr(unsafe.Pointer(&count)),
	)
	return count
}

func OcrProcessOptionsSetMaxRecognitionLineCount(processOpts uintptr, count int) {
	proc := oneocr.NewProc("OcrProcessOptionsSetMaxRecognitionLineCount")
	proc.Call(
		processOpts,
		uintptr(count),
	)
}

func OcrProcessOptionsGetResizeResolution(processOpts uintptr) (int, int) {
	var width, height int
	proc := oneocr.NewProc("OcrProcessOptionsGetResizeResolution")
	proc.Call(
		processOpts,
		uintptr(unsafe.Pointer(&width)),
		uintptr(unsafe.Pointer(&height)),
	)
	return width, height
}

func OcrProcessOptionsSetResizeResolution(processOpts uintptr, width int, height int) {
	proc := oneocr.NewProc("OcrProcessOptionsSetResizeResolution")
	proc.Call(
		processOpts,
		uintptr(width),
		uintptr(height),
	)
}

func RunOcrPipeline(pipeline, processOpts uintptr, img *Image) (uintptr, error) {
	var (
		result uintptr
		err    error
	)
	proc := oneocr.NewProc("RunOcrPipeline")
	ret, _, _ := proc.Call(
		pipeline,
		uintptr(unsafe.Pointer(img)),
		processOpts,
		uintptr(unsafe.Pointer(&result)),
	)
	if ret != 0 {
		err = fmt.Errorf("error code %d", ret)
	}
	return result, err
}

func GetImageAngle(result uintptr) float32 {
	var angle float32
	proc := oneocr.NewProc("GetImageAngle")
	proc.Call(
		result,
		uintptr(unsafe.Pointer(&angle)),
	)
	return angle
}

func GetOcrLineCount(result uintptr) int {
	var count int
	proc := oneocr.NewProc("GetOcrLineCount")
	proc.Call(
		result,
		uintptr(unsafe.Pointer(&count)),
	)
	return count
}

func GetOcrLine(result uintptr, index int) uintptr {
	var handle uintptr
	proc := oneocr.NewProc("GetOcrLine")
	proc.Call(
		result,
		uintptr(index),
		uintptr(unsafe.Pointer(&handle)),
	)
	return handle
}

func GetOcrLineBoundingBox(line uintptr) BoundingBox {
	var box *BoundingBox
	proc := oneocr.NewProc("GetOcrLineBoundingBox")
	proc.Call(
		line,
		uintptr(unsafe.Pointer(&box)),
	)
	return *box
}

func GetOcrLineContent(line uintptr) string {
	var ptr *byte
	proc := oneocr.NewProc("GetOcrLineContent")
	proc.Call(
		line,
		uintptr(unsafe.Pointer(&ptr)),
	)
	return ansi2String(ptr)
}

// TODO: GetOcrLineStyle

func GetOcrLineWordCount(line uintptr) int {
	var count int
	proc := oneocr.NewProc("GetOcrLineWordCount")
	proc.Call(
		line,
		uintptr(unsafe.Pointer(&count)),
	)
	return count
}

func GetOcrWord(line uintptr, index int) uintptr {
	var handle uintptr
	proc := oneocr.NewProc("GetOcrWord")
	proc.Call(
		line,
		uintptr(index),
		uintptr(unsafe.Pointer(&handle)),
	)
	return handle
}

func GetOcrWordBoundingBox(word uintptr) BoundingBox {
	var box *BoundingBox
	proc := oneocr.NewProc("GetOcrWordBoundingBox")
	proc.Call(
		word,
		uintptr(unsafe.Pointer(&box)),
	)
	return *box
}

func GetOcrWordContent(word uintptr) string {
	var ptr *byte
	proc := oneocr.NewProc("GetOcrWordContent")
	proc.Call(
		word,
		uintptr(unsafe.Pointer(&ptr)),
	)
	return ansi2String(ptr)
}

func GetOcrWordConfidence(word uintptr) float32 {
	var confidence float32
	proc := oneocr.NewProc("GetOcrWordConfidence")
	proc.Call(
		word,
		uintptr(unsafe.Pointer(&confidence)),
	)
	return confidence
}

func ReleaseOcrInitOptions(initOpts uintptr) {
	proc := oneocr.NewProc("ReleaseOcrInitOptions")
	proc.Call(initOpts)
}

func ReleaseOcrPipeline(pipeline uintptr) {
	proc := oneocr.NewProc("ReleaseOcrPipeline")
	proc.Call(pipeline)
}

func ReleaseOcrProcessOptions(processOpts uintptr) {
	proc := oneocr.NewProc("ReleaseOcrProcessOptions")
	proc.Call(processOpts)
}

func ReleaseOcrResult(result uintptr) {
	proc := oneocr.NewProc("ReleaseOcrResult")
	proc.Call(result)
}
