// Exercise: Images
// 前に解いた、 画像ジェネレーター　を覚えていますか？ 今回は、データのスライスの代わりに image.Image インタフェースの実装を返すようにしてみましょう。

// 自分の Image 型を定義し、 インタフェースを満たすのに必要なメソッド を実装し、 pic.ShowImage を呼び出してみてください。

// Bounds は、 image.Rect(0, 0, w, h) のようにして image.Rectangle を返すようにします。

// ColorModel は、 color.RGBAModel を返すようにします。

// At は、ひとつの色を返します。 生成する画像の色の値 v を color.RGBA{v, v, 255, 255} を利用して返すようにします。

package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct{}

// 言語サーバが必要な関数を提示してくれう
// At implements image.Image.
// func (*Image) At(x int, y int) color.Color {
// 	panic("unimplemented")
// }

// // Bounds implements image.Image.
// func (*Image) Bounds() image.Rectangle {
// 	panic("unimplemented")
// }

// // ColorModel implements image.Image.
// func (*Image) ColorModel() color.Model {
// 	panic("unimplemented")
// }

func (m *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (m *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 128, 128)
}

func (m *Image) At(x, y int) color.Color {
	v := uint8(x ^ y)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(&m)
}
