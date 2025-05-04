package winocr

type BoundingBox struct {
	TopLeft     Point `json:"top_left"`
	TopRight    Point `json:"top_right"`
	BottomRight Point `json:"bottom_right"`
	BottomLeft  Point `json:"bottom_left"`
}

type Point struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

type Image struct {
	Type     int32
	Width    int32
	Height   int32
	Reserved int32
	Step     int64
	DataPtr  *byte
}

type OcrResult struct {
	ImageAngle float32   `json:"image_angle"`
	Lines      []OcrLine `json:"lines"`
}

type OcrLine struct {
	Text         string      `json:"text"`
	BoundingRect BoundingBox `json:"bounding_rect"`
	Words        []OcrWord   `json:"words"`
}

type OcrWord struct {
	Text         string      `json:"text"`
	BoundingRect BoundingBox `json:"bounding_rect"`
	Confidence   float32     `json:"confidence"`
}
