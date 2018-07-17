package main

import (
	"fmt"
	graphics "github.com/disintegration/imaging"
	"image"
	"image/png"
	"log"
	"os"
)

func main() {
	// src, err := LoadImage("lena.jpg")
	// src, err := LoadImage("wm0_90.jpeg")
	src, err := LoadImage("1.png")
	if err != nil {
		log.Fatal("1,", err)
	}

	// dst := image.NewRGBA(image.Rect(0, 0, 350, 400))

	// err = graphics.Rotate(dst, src, &graphics.RotateOptions{3.5})
	dst := graphics.Rotate90(src)
	if err != nil {
		log.Fatal("2", err)
	}

	// 需要保存的文件
	imgcounter := 123
	saveImage(fmt.Sprintf("%03d.png", imgcounter), dst)
}

// LoadImage decodes an image from a file.
func LoadImage(path string) (img image.Image, err error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("l1:", err)
		return
	}
	defer file.Close()
	img, _, err = image.Decode(file)
	fmt.Println("l2:", err)
	return
}

// 保存Png图片
func saveImage(path string, img image.Image) (err error) {
	// 需要保存的文件
	imgfile, err := os.Create(path)
	defer imgfile.Close()

	// 以PNG格式保存文件
	err = png.Encode(imgfile, img)
	if err != nil {
		log.Fatal("3", err)
	}
	return
}
