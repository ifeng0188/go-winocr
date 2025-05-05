package winocr

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/windows"
)

// CreateOcrInitOptions creates and returns a handle to OCR initialization options.
func CreateOcrInitOptions() uintptr {
	var handle uintptr
	proc := oneocr.NewProc("CreateOcrInitOptions")
	proc.Call(
		uintptr(unsafe.Pointer(&handle)),
	)
	return handle
}

// OcrInitOptionsSetUseModelDelayLoad enables or disables lazy loading of the OCR model.
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

// CreateOcrPipeline creates and returns a handle to OCR pipeline.
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

// CreateOcrProcessOptions creates and returns a handle to OCR processing options.
func CreateOcrProcessOptions() uintptr {
	var handle uintptr
	proc := oneocr.NewProc("CreateOcrProcessOptions")
	proc.Call(
		uintptr(unsafe.Pointer(&handle)),
	)
	return handle
}

// OcrProcessOptionsGetMaxRecognitionLineCount returns the maximum number of text lines
func OcrProcessOptionsGetMaxRecognitionLineCount(processOpts uintptr) int {
	var count int
	proc := oneocr.NewProc("OcrProcessOptionsGetMaxRecognitionLineCount")
	proc.Call(
		processOpts,
		uintptr(unsafe.Pointer(&count)),
	)
	return count
}

// OcrProcessOptionsSetMaxRecognitionLineCount sets the maximum number of text lines
func OcrProcessOptionsSetMaxRecognitionLineCount(processOpts uintptr, count int) {
	proc := oneocr.NewProc("OcrProcessOptionsSetMaxRecognitionLineCount")
	proc.Call(
		processOpts,
		uintptr(count),
	)
}

// OcrProcessOptionsGetResizeResolution returns the current width and height settings
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

// OcrProcessOptionsSetResizeResolution sets the resolution that images will be resized to
func OcrProcessOptionsSetResizeResolution(processOpts uintptr, width int, height int) {
	proc := oneocr.NewProc("OcrProcessOptionsSetResizeResolution")
	proc.Call(
		processOpts,
		uintptr(width),
		uintptr(height),
	)
}

// RunOcrPipeline executes the OCR pipeline on the provided image with the specified
// processing options. Returns a handle to the recognition results.
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

// GetImageAngle returns the detected rotation angle of the image in degrees.
func GetImageAngle(result uintptr) float32 {
	var angle float32
	proc := oneocr.NewProc("GetImageAngle")
	proc.Call(
		result,
		uintptr(unsafe.Pointer(&angle)),
	)
	return angle
}

// GetOcrLineCount returns the number of text lines detected in the image.
func GetOcrLineCount(result uintptr) int {
	var count int
	proc := oneocr.NewProc("GetOcrLineCount")
	proc.Call(
		result,
		uintptr(unsafe.Pointer(&count)),
	)
	return count
}

// GetOcrLine returns a handle to the specified text line by index.
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

// GetOcrLineBoundingBox returns the bounding box coordinates for the specified text line.
func GetOcrLineBoundingBox(line uintptr) BoundingBox {
	var box *BoundingBox
	proc := oneocr.NewProc("GetOcrLineBoundingBox")
	proc.Call(
		line,
		uintptr(unsafe.Pointer(&box)),
	)
	return *box
}

// GetOcrLineContent returns the text content of the specified line.
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

// GetOcrLineWordCount returns the number of words in the specified text line.
func GetOcrLineWordCount(line uintptr) int {
	var count int
	proc := oneocr.NewProc("GetOcrLineWordCount")
	proc.Call(
		line,
		uintptr(unsafe.Pointer(&count)),
	)
	return count
}

// GetOcrWord returns a handle to the specified word by index within a text line.
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

// GetOcrWordBoundingBox returns the bounding box coordinates for the specified word.
func GetOcrWordBoundingBox(word uintptr) BoundingBox {
	var box *BoundingBox
	proc := oneocr.NewProc("GetOcrWordBoundingBox")
	proc.Call(
		word,
		uintptr(unsafe.Pointer(&box)),
	)
	return *box
}

// GetOcrWordContent returns the text content of the specified word.
func GetOcrWordContent(word uintptr) string {
	var ptr *byte
	proc := oneocr.NewProc("GetOcrWordContent")
	proc.Call(
		word,
		uintptr(unsafe.Pointer(&ptr)),
	)
	return ansi2String(ptr)
}

// GetOcrWordConfidence returns the recognition confidence score (0-1) for the specified word.
func GetOcrWordConfidence(word uintptr) float32 {
	var confidence float32
	proc := oneocr.NewProc("GetOcrWordConfidence")
	proc.Call(
		word,
		uintptr(unsafe.Pointer(&confidence)),
	)
	return confidence
}

// ReleaseOcrInitOptions releases the resources associated with OCR initialization options.
// This should be called when the options are no longer needed.
func ReleaseOcrInitOptions(initOpts uintptr) {
	proc := oneocr.NewProc("ReleaseOcrInitOptions")
	proc.Call(initOpts)
}

// ReleaseOcrPipeline releases the resources associated with an OCR pipeline.
// This should be called when the pipeline is no longer needed.
func ReleaseOcrPipeline(pipeline uintptr) {
	proc := oneocr.NewProc("ReleaseOcrPipeline")
	proc.Call(pipeline)
}

// ReleaseOcrProcessOptions releases the resources associated with OCR processing options.
// This should be called when the options are no longer needed.
func ReleaseOcrProcessOptions(processOpts uintptr) {
	proc := oneocr.NewProc("ReleaseOcrProcessOptions")
	proc.Call(processOpts)
}

// ReleaseOcrResult releases the resources associated with OCR recognition results.
// This should be called when the results are no longer needed.
func ReleaseOcrResult(result uintptr) {
	proc := oneocr.NewProc("ReleaseOcrResult")
	proc.Call(result)
}
