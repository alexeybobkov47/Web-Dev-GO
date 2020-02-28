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
		"345345345https://yandex.ru",
		"https://golang.org",
		"https://google.com",
		"https://github.com",
		"https://dtf.ru",
		"https://geekbrains.ru",
	}

	var result = make([]string, 0, len(sites))
	result = search(searchReq, sites)
	log.Printf("Найдены совпадения: %q", result)

}

func search(searchReq string, sites []string) []string {
	out := make([]string, 0, 1)
	for _, site := range sites {
		res := getReq(site)
		if strings.Contains(string(res), searchReq) {
			out = append(out, site)

		}

	}
	return out
}

func getReq(reqURL string) []byte {
	resp, err := http.Get(reqURL)
	if err != nil {
		log.Println("Ошибка в getReq http.Get " + reqURL)
		return nil

	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Ошибка в getReq ioutil.ReadAll")
		return nil
	}
	return body
}
