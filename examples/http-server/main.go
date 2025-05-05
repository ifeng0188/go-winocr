package main

import (
	"encoding/json"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"net/http"

	"github.com/ifeng0188/go-winocr"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Result  any    `json:"result,omitempty"`
}

var engine *winocr.OcrEngine

func init() {
	winocr.SetOcrDllPath(`D:/Temp/winocr`)

	engine = winocr.NewOcrEngine()

	engine.EnableModelDelayLoad()
}

func handleOcr(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "只支持 POST 请求", http.StatusMethodNotAllowed)
		return
	}

	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("image")
	if err != nil {
		json.NewEncoder(w).Encode(Response{
			Code:    -1,
			Message: "请上传图片文件",
		})
		return
	}
	defer file.Close()

	contentType := handler.Header.Get("Content-Type")
	if contentType != "image/jpeg" && contentType != "image/png" {
		json.NewEncoder(w).Encode(Response{
			Code:    -1,
			Message: "只支持 JPEG 和 PNG 格式的图片",
		})
		return
	}

	img, _, err := image.Decode(file)
	if err != nil {
		json.NewEncoder(w).Encode(Response{
			Code:    -1,
			Message: fmt.Sprintf("图片解码失败: %v", err),
		})
		return
	}

	format := r.FormValue("format")
	if format == "" {
		format = "text"
	}
	if format != "text" && format != "json" {
		json.NewEncoder(w).Encode(Response{
			Code:    -1,
			Message: "不支持的输出格式，只支持 text 或 json",
		})
		return
	}

	result, err := engine.Recognize(img, format)
	if err != nil {
		json.NewEncoder(w).Encode(Response{
			Code:    -1,
			Message: fmt.Sprintf("OCR 识别失败: %v", err),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if format == "json" {
		var jsonResult any
		if err := json.Unmarshal([]byte(result), &jsonResult); err != nil {
			json.NewEncoder(w).Encode(Response{
				Code:    -1,
				Message: "JSON 解析失败",
			})
			return
		}
		json.NewEncoder(w).Encode(Response{
			Code:    0,
			Message: "success",
			Result:  jsonResult,
		})
	} else {
		json.NewEncoder(w).Encode(Response{
			Code:    0,
			Message: "success",
			Result:  result,
		})
	}
}

func main() {
	http.HandleFunc("/ocr", handleOcr)

	port := ":8080"
	log.Printf("服务器启动在 http://localhost%s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
