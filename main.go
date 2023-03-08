package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"time"

	"github.com/go-vgo/robotgo"
)

const (
	red   = "ffff0000"
	white = "ffffffffffff"
)

func main() {
	for {
		checkColor()
	}
}

func checkColor() {
	x1, y1 := 1013, 1190 // 1440p
	w, h := 530, 1

	img := robotgo.CaptureImg(x1, y1, w, h)

	color2hex := ""

	for x := 0; x <= w-1; x++ {
		color1 := img.At(x, 0)

		color1hex := colortoHex(color1)

		if (color1hex == red && color2hex == white) || (color1hex == white && color2hex == red) {
			robotgo.Click()
			return
		}
		color2hex = color1hex
	}
}

func colortoHex(c color.Color) string {
	r, g, b, _ := c.RGBA()
	return fmt.Sprintf("%02x%02x%02x", r, g, b)
}

// for debugging
func saveImageToFile(img image.Image, name string) {
	f, _ := os.Create(fmt.Sprintf("%d_%s.png", time.Now().UnixMilli(), name))
	png.Encode(f, img)
}
