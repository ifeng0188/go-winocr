# go-winocr

A Go binding for Windows OCR engine that provides text recognition capabilities.

## Features

- Fast and accurate text recognition
- Support for multiple output formats (text/json)
- Detailed OCR results including:
  - Text content
  - Bounding boxes
  - Word confidence scores
  - Image angle detection

## Installation

```bash
go get github.com/ifeng0188/go-winocr
```

Before using this package, you need to download the Windows OCR engine files and place them in a directory:
- oneocr.dll
- oneocr.onemodel
- onnxruntime.dll

## Quick Start

```go
package main

import (
    "fmt"
    "image"
    "os"
    "github.com/ifeng0188/go-winocr"
)

func main() {
    // Set OCR DLL path
    winocr.SetOcrDllPath("path/to/ocr/dlls")

    // Create OCR engine
    engine := winocr.NewOcrEngine()
    defer engine.Close()

    // Enable model delay load (optional)
    engine.EnableModelDelayLoad()

    // Load and decode image
    file, _ := os.Open("image.png")
    defer file.Close()
    img, _, _ := image.Decode(file)

    // Perform OCR
    result, err := engine.Recognize(img, "text")
    if err != nil {
        fmt.Printf("OCR failed: %v\n", err)
        return
    }

    fmt.Printf("OCR Result:\n%s\n", result)
}
```

## Credits

- [win11-oneocr](https://github.com/b1tg/win11-oneocr) - Inspiration for this project

## License

This project is licensed under the MIT License.