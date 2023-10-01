package main

import (
	"gotask/web"
	"log"
	"sync"
)

const (
	processes = 5
	subStr    = "Go"
)

var (
	urls = []string{
		"https://golang.org",
		"https://go.dev/",
		"https://ru.wikipedia.org/wiki/Go",
		"http://golang-book.ru/",
		"https://golangify.com/",
		"https://tproger.ru/translations/golang-basics/",
		"https://goforum.info/",
	}
)

func main() {

	var wg sync.WaitGroup
	sem := make(chan struct{}, processes)

	mu := sync.Mutex{}
	res := 0

	for _, url := range urls {
		sem <- struct{}{}
		wg.Add(1)
		go func(s string) {
			defer func() {
				<-sem
				wg.Done()
			}()
			count, err := web.GetCount(s, subStr)
			if err != nil {
				log.Println("Error:", err)
				return
			}
			mu.Lock()
			res += count
			mu.Unlock()
		}(url)
	}

	wg.Wait()
	log.Println("Res:", res)
}
