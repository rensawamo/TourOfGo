// Exercise: Readers
// ASCII文字 'A' の無限ストリームを出力する Reader 型を実装してください。

package main

import (
	"fmt"

	"golang.org/x/tour/reader"
)

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader) Read(b []byte) (int, error) {
	if b == nil {
		return 0, fmt.Errorf("nil error")
	}
	for i := range b { // byteをAの バイトにしてしまう
		b[i] = 'A'
	}
	fmt.Print("実行") // 無限ストリームの証明
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}

//テスト構文
//  Validateは   1mbにforか lenが達したときにとじられて okが出力
// func Validate(r io.Reader)
// 	b := make([]byte, 1024, 2048)
// 	i, o := 0, 0
// 	for ; i < 1<<20 && o < 1<<20; i++ { // test 1mb
// 		n, err := r.Read(b) // 読み込む byteを最大 1024にしている
// 		for i, v := range b[:n] {  // ひと文字づつ読み出して  文字がA であることをチェックする
// 			if v != 'A' {
// 				fmt.Fprintf(os.Stderr, "got byte %x at offset %v, want 'A'\n", v, o+i)
// 				return
// 			}
// 		}
// 		o += n
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "read error: %v\n", err)
// 			return
// 		}
// 	}
// 	if o == 0 {
// 		fmt.Fprintf(os.Stderr, "read zero bytes after %d Read calls\n", i)
// 		return
// 	}
// 	fmt.Println("OK!")
// }
