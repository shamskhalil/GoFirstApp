package main

import (
	"fmt"
	"net/http"
	"time"
)

type response struct {
	code    int
	abbr    string
	err     error
	elapsed string
}

func main() {
	now := time.Now()
	urls := map[string]string{
		"google":    "https://google.com",
		"facebook":  "https://facebook.com",
		"tiktok":    "https://tiktok.com",
		"instagram": "https://instagram.com",
	}
	fmt.Println(len(urls))
	ch := make(chan response)
	for abbr, url := range urls {
		go func(u string, ab string) {
			fetchChan(ch, u, ab)
		}(url, abbr)
	}
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-ch)
	}
	fmt.Printf("Elsapsed: %v\n", time.Since(now))
}

func fetchChan(ch chan response, url string, abbr string) {
	now := time.Now()
	res, err := http.Get(url)
	if err != nil {

		ch <- response{code: 0, abbr: abbr, err: err, elapsed: fmt.Sprintf("%s\n", time.Since(now))}
		fmt.Println(err)
		close(ch)
		return
	}
	ch <- response{code: res.StatusCode, abbr: abbr, err: err, elapsed: fmt.Sprintf("%s\n", time.Since(now))}
}

func fetch(url string) (int, error) {
	fmt.Println("Making resquest ....")
	res, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	code := res.StatusCode
	res.Body.Close()
	return code, nil
}
