package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	// 이미지의 크기 정의
	width := 204
	height := 249

	// 새 이미지 생성
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// 모든 픽셀을 검은색으로 설정
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, color.Black)
		}
	}

	// 이미지 파일로 저장
	outputFile, err := os.Create("output_black.png")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	// PNG 형식으로 이미지 저장
	err = png.Encode(outputFile, img)
	if err != nil {
		panic(err)
	}

	println("검은색 이미지 생성이 완료되었습니다.")
}
