package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

func main2() {
	// 이미지 크기 설정
	width := 21 // 10 * 2 + 1
	height := 21

	// 새 이미지 생성
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// 각 픽셀을 순회하며 원 안이면 흰색으로 채우기, 원 밖이면 투명으로 채우기
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			distance := math.Sqrt(math.Pow(float64(x-10), 2) + math.Pow(float64(y-10), 2)) // 중심으로부터의 거리 계산

			// 반지름 10인 원 내부인지 확인
			if distance <= 10 {
				img.Set(x, y, color.White) // 흰색
			} else {
				img.Set(x, y, color.Transparent) // 투명
			}
		}
	}

	// 파일 생성
	file, err := os.Create("output.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// PNG 형식으로 이미지 파일 작성
	err = png.Encode(file, img)
	if err != nil {
		panic(err)
	}
}
