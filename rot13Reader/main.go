// Exercise: rot13Reader
// よくあるパターンは、別の io.Reader をラップし、ストリームの内容を何らかの方法で変換するio.Readerです。

// 例えば、 gzip.NewReader は、 io.Reader (gzipされたデータストリーム)を引数で受け取り、 *gzip.Reader を返します。 その *gzip.Reader は、 io.Reader (展開したデータストリーム)を実装しています。

// io.Reader を実装し、 io.Reader でROT13 換字式暗号( substitution cipher )をすべてのアルファベットの文字に適用して読み出すように rot13Reader を実装してみてください。

// rot13Reader 型は提供済みです。 この Read メソッドを実装することで io.Reader インタフェースを満たしてください。
package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

// Read メソットを 13ビット移動
func (r *rot13Reader) Read(p []byte) (n int, e error) {
	// データ量をよみこんで  for で回す
	// n, e = r.r.Read(p)
	for i := range p { // byteはそのままでも for 対応
		p[i] = rot13(p[i])
	}
	return
}

// 13 bitずらす 定石関数
func rot13(n byte) byte {
	switch {
	case ('A' <= n && n <= 'Z'):
		return (n-'A'+13)%26 + 'A'
	case ('a' <= n && n <= 'z'):
		return (n-'a'+13)%26 + 'a'
	default:
		return n
	}
}

func main() {
	//func strings.NewReader(s string) *strings.Readerこれで  *strings.Re	?ader型を .copyの第二引数に当てはめるために  rot13Reader struct で io.Reader から  Read をひろげ
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}

	// func Copy(dst Writer, src Reader) (written int64, err error)
	// Read メソッドを実装しないと うごかない
	io.Copy(os.Stdout, &r)
}
