package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"time"

	"github.com/ifeng0188/go-winocr"
)

func main() {
	winocr.SetOcrDllPath(`D:/Temp/winocr`)

	engine := winocr.NewOcrEngine()
	defer engine.Close()

	engine.EnableModelDelayLoad()

	if len(os.Args) < 2 {
		fmt.Printf("请拖放图片到本程序上\n")
		exit()
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("打开图像文件失败: %v\n", err)
		exit()
	}

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Printf("解码图像失败: %v\n", err)
		exit()
	}

	start := time.Now()
	result, err := engine.Recognize(img, "text")
	if err != nil {
		fmt.Printf("识别失败: %v\n", err)
		exit()
	}
	since := time.Since(start)

	fmt.Printf("识别成功:\n%s\n\n耗时:\n%s\n", result, since)
	exit()
}

func exit() {
	fmt.Printf("\n按回车键退出...")
	fmt.Scanln()
	os.Exit(0)
}
