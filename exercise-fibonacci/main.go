// Exercise: Fibonacci closure
// 関数を用いた面白い例を見てみましょう。

// fibonacci (フィボナッチ)関数を実装しましょう。この関数は、連続するフィボナッチ数(0, 1, 1, 2, 3, 5, ...)を返す関数(クロージャ)を返します。

package main

import "fmt"

func fibonacci() func() int {
	// クロージャは 関数内変数を更新する
	f2, f1 := 0, 1
	return func() int {
		f := f2
		f2, f1 = f1, f+f1
		return f
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
