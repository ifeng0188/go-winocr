package winocr

import (
	"encoding/json"
	"image"
	"image/draw"
	"path/filepath"
	"strings"
)

// OcrEngine represents the OCR engine instance that performs text recognition.
type OcrEngine struct {
	initOpts    uintptr
	processOpts uintptr
}

// NewOcrEngine creates a new instance of the OCR engine with default settings.
func NewOcrEngine() *OcrEngine {
	e := &OcrEngine{
		initOpts:    CreateOcrInitOptions(),
		processOpts: CreateOcrProcessOptions(),
	}
	return e
}

// EnableModelDelayLoad enables lazy loading of the OCR model.
func (e *OcrEngine) EnableModelDelayLoad() error {
	OcrInitOptionsSetUseModelDelayLoad(e.initOpts, true)
	return nil
}

// GetMaxRecognitionLineCount returns the maximum number of text lines
func (e *OcrEngine) GetMaxRecognitionLineCount() int {
	return OcrProcessOptionsGetMaxRecognitionLineCount(e.processOpts)
}

// SetMaxRecognitionLineCount sets the maximum number of text lines
func (e *OcrEngine) SetMaxRecognitionLineCount(count int) error {
	OcrProcessOptionsSetMaxRecognitionLineCount(e.processOpts, count)
	return nil
}

// GetResizeResolution returns the current resolution settings used for image resizing
func (e *OcrEngine) GetResizeResolution() (int, int) {
	return OcrProcessOptionsGetResizeResolution(e.processOpts)
}

// SetResizeResolution sets the resolution that images will be resized to
func (e *OcrEngine) SetResizeResolution(width, height int) error {
	OcrProcessOptionsSetResizeResolution(e.processOpts, width, height)
	return nil
}

// Recognize performs OCR on the provided image and returns the results.
// The format parameter can be either "text" for plain text output or "json" for
// detailed recognition results including bounding boxes and confidence scores.
func (e *OcrEngine) Recognize(img image.Image, format string) (string, error) {
	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()
	rgba := image.NewRGBA(bounds)
	draw.Draw(rgba, bounds, img, bounds.Min, draw.Src)

	processed := &Image{
		Type:     3,
		Width:    int32(width),
		Height:   int32(height),
		Reserved: 0,
		Step:     int64(width * 4),
		DataPtr:  &rgba.Pix[0],
	}

	pipeline := CreateOcrPipeline(e.initOpts, filepath.Join(ocrPath, "oneocr.onemodel"), modelKey)
	defer ReleaseOcrPipeline(pipeline)

	result, err := RunOcrPipeline(pipeline, e.processOpts, processed)
	if err != nil {
		return "", err
	}
	defer ReleaseOcrResult(result)

	var resultFormat string
	if format == "json" {
		resultFormat = e.getJson(result)
	} else {
		resultFormat = e.getText(result)
	}
	return resultFormat, nil
}

// Close releases all resources associated with the OCR engine.
// This should be called when the engine is no longer needed.
func (e *OcrEngine) Close() {
	ReleaseOcrProcessOptions(e.processOpts)
	ReleaseOcrInitOptions(e.initOpts)
}

func (e *OcrEngine) getText(result uintptr) string {
	var textBuilder strings.Builder
	lineCount := GetOcrLineCount(result)
	for i := range lineCount {
		if i > 0 {
			textBuilder.WriteString("\n")
		}
		line := GetOcrLine(result, i)
		lineContent := GetOcrLineContent(line)
		textBuilder.WriteString(lineContent)
	}
	return textBuilder.String()
}

func (e *OcrEngine) getJson(result uintptr) string {
	lineCount := GetOcrLineCount(result)
	data := &OcrResult{
		ImageAngle: GetImageAngle(result),
		Lines:      make([]OcrLine, lineCount),
	}
	for i := range lineCount {
		line := GetOcrLine(result, i)
		lineWordCount := GetOcrLineWordCount(line)
		data.Lines[i] = OcrLine{
			Text:         GetOcrLineContent(line),
			BoundingRect: GetOcrLineBoundingBox(line),
			Words:        make([]OcrWord, lineWordCount),
		}
		for j := range lineWordCount {
			word := GetOcrWord(line, j)
			data.Lines[i].Words[j] = OcrWord{
				Text:         GetOcrWordContent(word),
				BoundingRect: GetOcrWordBoundingBox(word),
				Confidence:   GetOcrWordConfidence(word),
			}
		}
	}
	json, _ := json.Marshal(data)
	return string(json)
}
