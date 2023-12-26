// Exercise: Maps
// WordCount 関数を実装してみましょう。string s で渡される文章の、各単語の出現回数のmapを返す必要があります。 wc.Test 関数は、引数に渡した関数に対しテストスイートを実行し、成功か失敗かを結果に表示します。

// strings.Fields で、何かヒントを得ることができるはずです。

// Note: このテストスイートで何を入力とし、何を期待しているかについては、golang.org/x/tour/wcを見てみてください。

package main

import (
	"fmt"
	"strings"
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	//  帰り値となる変数を定義する  mapはmakeでつくる
	m := make(map[string]int)
	sl := strings.Fields(s)
	fmt.Print(sl)

	for _, w := range sl {
		if v, ok := m[w]; !ok {
			m[w] = 0
		} else {
			m[w] = m[w] + 1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
