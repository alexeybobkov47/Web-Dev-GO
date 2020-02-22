package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	searchReq := "yandex"
	sites := []string{
		"https://yandex.ru",
		"https://golang.org",
		"https://google.com",
		"https://github.com",
		"https://dtf.ru",
		"https://geekbrains.ru",
	}

	var (
		result = make([]string, 0, len(sites))
	)
	for _, site := range sites {
		in := search(searchReq, site)
		result = append(result, in)
	}

	log.Printf("%v", result)

}

func search(searchReq string, site string) string {
	res := getReq(site)
	if strings.Contains(string(res), searchReq) {
		return "\nНа сайте " + site + " есть совпадения по запросу " + searchReq
	}
	return "\nНа сайте " + site + " нет совпадений по запросу " + searchReq
}

func getReq(reqURL string) []byte {
	resp, _ := http.Get(reqURL)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return body
}

// Вывод с консоли
// $ go run searchReq.go
// 2020/02/22 22:05:17 [
// На сайте https://yandex.ru есть совпадения по запросу yandex
// На сайте https://golang.org нет совпадений по запросу yandex
// На сайте https://google.com нет совпадений по запросу yandex
// На сайте https://github.com нет совпадений по запросу yandex
// На сайте https://dtf.ru есть совпадения по запросу yandex
// На сайте https://geekbrains.ru есть совпадения по запросу yandex]
