package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

func nilHandler(err error) {
	if err != nil {
		panic(err)
	}
}

func calcDistance(x1, y1, x2, y2 int) float64 {
	dx := float64(x2 - x1)
	dy := float64(y2 - y1)
	return math.Sqrt(math.Pow(dx, 2) + math.Pow(dy, 2))
}

func saveImg(fileName string, imgRGBA *image.RGBA) {
	outputFile, err := os.Create(fileName)
	nilHandler(err)
	defer outputFile.Close()

	err = png.Encode(outputFile, imgRGBA)
	nilHandler(err)
}

func main() {
	println(int(math.Pow(float64(100), 2)))

	FILENAME := "output2.png"
	MN_WIDTH, MX_WIDTH := 0, 204
	MN_HEIGHT, MX_HEIGHT := 0, 249

	TG_RADIUS := 10
	TG_MN_WIDTH, TG_MX_WIDTH := MN_WIDTH+TG_RADIUS, MX_WIDTH-TG_RADIUS
	TG_MN_HEIGHT, TG_MX_HEIGHT := MN_HEIGHT+TG_RADIUS, MX_HEIGHT-TG_RADIUS

	POINTER_X, POINTER_Y := 0, 0

	imgRGBA := image.NewRGBA(image.Rect(0, 0, MX_WIDTH, MX_HEIGHT))
	for x := 0; x < MX_WIDTH; x++ {
		for y := 0; y < MX_HEIGHT; y++ {
			isLeftX := TG_MN_WIDTH >= x
			isRightX := TG_MX_WIDTH <= x

			isTopY := TG_MN_HEIGHT >= y
			isBtmY := TG_MX_HEIGHT <= y

			isTgWidth := isLeftX || isRightX
			isTgHeight := isTopY || isBtmY
			isTg := isTgWidth && isTgHeight
			if !isTg {
				imgRGBA.Set(x, y, color.Black)
				continue
			}

			if isLeftX {
				POINTER_X = TG_MN_WIDTH
			} else {
				POINTER_X = TG_MX_WIDTH
			}

			if isTopY {
				POINTER_Y = TG_MN_HEIGHT
			} else {
				POINTER_Y = TG_MX_HEIGHT
			}

			distance := calcDistance(POINTER_X, POINTER_Y, x, y)
			println(distance, float64(TG_RADIUS), int(distance), TG_RADIUS)
			// tgDistance := float64(TG_RADIUS) - distance
			isInsideRadius := distance < float64(TG_RADIUS)
			// isInsideRadius := distance < float64(TG_RADIUS)
			if isInsideRadius {
				imgRGBA.Set(x, y, color.Black)
				continue
			}

			imgRGBA.Set(x, y, color.White)
		}
	}

	saveImg(FILENAME, imgRGBA)

	println(POINTER_X, POINTER_Y)
	println(TG_MN_WIDTH, TG_MX_WIDTH, TG_MN_HEIGHT, TG_MX_HEIGHT)

	// width := 204
	// height := 249

	// img := image.NewRGBA(image.Rect(0, 0, width, height))

	// for x := 0; x < width; x++ {
	// 	for y := 0; y < height; y++ {

	// 		math.Pow
	// 		distance := math.Sqrt(
	// 			math.Pow(float64(x-10), 2) +
	// 				math.Pow(float64(y-10), 2),
	// 		)

	// 		println(x, y, int(distance))

	// 		if distance <= 10 {
	// 			img.Set(x, y, color.Black)
	// 			continue
	// 		}

	// 		img.Set(x, y, color.Transparent)
	// 		continue
	// 	}
	// }

	// // 이미지 파일로 저장
	// outputFile, err := os.Create("output_black.png")
	// if err != nil {
	// 	panic(err)
	// }
	// defer outputFile.Close()

	// // PNG 형식으로 이미지 저장
	// err = png.Encode(outputFile, img)
	// if err != nil {
	// 	panic(err)
	// }

	// println("검은색 이미지 생성이 완료되었습니다.")
}
