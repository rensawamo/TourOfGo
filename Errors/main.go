// Exercise: Errors
// Sqrt 関数を 以前の演習 からコピーし、 error の値を返すように修正してみてください。

// Sqrt は、複素数をサポートしていないので、負の値が与えられたとき、nil以外のエラー値を返す必要があります。

// 新しい型:

// type ErrNegativeSqrt float64
// を作成してください。

// そして、 ErrNegativeSqrt(-2).Error() で、 "cannot Sqrt negative number: -2" を返すような:

// func (e ErrNegativeSqrt) Error() string
// メソッドを実装し、 error インタフェースを満たすようにします。

// 注意: Error メソッドの中で、 fmt.Sprint(e) を呼び出すことは、無限ループのプログラムになることでしょう。 最初に fmt.Sprint(float64(e)) として e を変換しておくことで、これを避けることができます。 なぜでしょうか？

// 負の値が与えられたとき、 ErrNegativeSqrt の値を返すように Sqrt 関数を修正してみてください。

package main

import (
	"fmt"
	"math"
)

const cd = 0.00001

// error 型は fmt.Stringer に組み込まれているいる
// type error interface {
//     Error() string
// }

type ErrNegativeSqrt float64

// error inter face の戻り値も    Springf でカスタムできる
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

func Sqrt(x float64) (float64, error) {
	// ネガティブ平方根は 実在しない
	if x < 0 {
		// しかし fmt.Errorf がくみこまれてるから errorインターフェイスをわざわざいじらなくていい
		return 0, ErrNegativeSqrt(x)
	}

	z, d := x/2, 0.

	for i := 1; i <= 10; i++ {
		z -= (z*z - x) / (2 * z)
		if math.Abs(z-d) < cd {
			break
		}
		d = z
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
