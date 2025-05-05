package winocr

// BoundingBox represents a rectangular region in the image defined by four corner points.
type BoundingBox struct {
	TopLeft     Point `json:"top_left"`
	TopRight    Point `json:"top_right"`
	BottomRight Point `json:"bottom_right"`
	BottomLeft  Point `json:"bottom_left"`
}

// Point represents a 2D coordinate point in the image.
type Point struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

// Image represents an input image for OCR processing.
type Image struct {
	Type     int32
	Width    int32
	Height   int32
	Reserved int32
	Step     int64
	DataPtr  *byte
}

// OcrResult represents the complete OCR recognition result for an image.
type OcrResult struct {
	ImageAngle float32   `json:"image_angle"`
	Lines      []OcrLine `json:"lines"`
}

// OcrLine represents a single line of recognized text in the image.
type OcrLine struct {
	Text         string      `json:"text"`
	BoundingRect BoundingBox `json:"bounding_rect"`
	Words        []OcrWord   `json:"words"`
}

// OcrWord represents a single word within a line of recognized text.
type OcrWord struct {
	Text         string      `json:"text"`
	BoundingRect BoundingBox `json:"bounding_rect"`
	Confidence   float32     `json:"confidence"`
}
