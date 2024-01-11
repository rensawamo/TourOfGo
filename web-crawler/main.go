// Exercise: Web Crawler
// この演習では、ウェブクローラ( web crawler )を並列化するため、Goの並行性の特徴を使います。

// 同じURLを2度取ってくることなく並列してURLを取ってくるように、 Crawl 関数を修正してみてください(注1)。

// 補足: 工夫すれば Crawl 関数のみの修正で実装できますが、無理に Crawl 関数内部に収める必要はありません。

// ひとこと: mapにフェッチしたURLのキャッシュを保持できますが、mapだけでは並行実行時の安全性はありません!

package main

import (
	"fmt"
	"sync"
)

func Crawl(url string, depth int, fetcher Fetcher) {
	var (
		cache = &sync.Map{}
		wg    = &sync.WaitGroup{}
		crawl func(url string, depth int)
	)

	// 再起関数の定義
	crawl = func(url string, depth int) {
		if depth <= 0 {
			return
		}

		// sync.Mapで 参照もともメモリのキャッシュですでによみこんだurl かどうかみることができる
		if _, ok := cache.Load(url); ok {
			return
		}
		// まだの場合cacheメモリについかする
		cache.Store(url, struct{}{})

		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("found: %s %q\n", url, body)
		wg.Add(len(urls)) // 同時に実行する goroutinの数を指定  しっかり数をあわさないと 残りあるっていわれる
		// bodyのitereater
		for _, u := range urls {
			go func(u string) {
				crawl(u, depth-1) // urls を読み込むごとに goroutinのカウントダウンを減らしていく
				wg.Done()
			}(u) // そして次のurlsを読み込んでいく再帰

		}
	}

	crawl(url, depth)
	wg.Wait() // カウントダウンがおわるまでまたせている
	return
}

// depthはkekyの数
func main() {
	Crawl("https://golang.org/", 4, fetcher)
}

type fakeFetcher map[string]*fakeResult

// fetcherに独自のmethodを付与
type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

type fakeResult struct {
	body string
	urls []string
}

// fakeFetcherのダミーに対して （f fakeFetcher） 動く関数
func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok { // urlが存在すれば ボディとかを返す
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
