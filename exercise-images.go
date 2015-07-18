package main

//import "golang.org/x/tour/pic"
import (
	"fmt"
	"image/color"
	"image"
)

type Image struct {}

func (this Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (this Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 200, 200)
}

func (this Image) At(x, y int) color.Color {
	v := uint8(x ^ y)
	return color.RGBA{v, v, 255, 255}
}

func (this Image) String() string {
	return fmt.Sprintf("Image %v", this.Bounds())
}

func main() {
	m := Image{}
//	pic.ShowImage(m)
	fmt.Println(m)
}