// Exercise: Stringers
// IPAddr 型を実装してみましょう IPアドレスをドットで4つに区切った( dotted quad )表現で出力するため、 fmt.Stringer インタフェースを実装してください。

// 例えば、 IPAddr{1, 2, 3, 4} は、 "1.2.3.4" として出力するようにします。

package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

// defaltの IPAdder 構造体の String関数をいじる
func (i IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", i[0], i[1], i[2], i[3])
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
