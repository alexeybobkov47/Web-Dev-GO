package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/", firstHandle)
	router.HandleFunc("/setcookie", setCookieHandle)
	router.HandleFunc("/getcookie", getCookieHandle)
	router.HandleFunc("/search", searchHandle)
	port := "8090"
	log.Printf("start listen on port %v", port)
	log.Fatal(http.ListenAndServe(":"+port, router))

}

func firstHandle(wr http.ResponseWriter, _ *http.Request) {
	_, _ = wr.Write([]byte("Hello World!"))
}

func setCookieHandle(wr http.ResponseWriter, req *http.Request) {

	http.SetCookie(wr, &http.Cookie{
		Name:  "AlexeyBobkov",
		Value: "323",
	})

}
func getCookieHandle(wr http.ResponseWriter, req *http.Request) {

	_, _ = fmt.Fprintln(wr, req.Header.Get("Cookie"))
	cookie, _ := req.Cookie("AlexeyBobkov")
	_, _ = fmt.Fprintln(wr, cookie.Name)
	_, _ = fmt.Fprintln(wr, cookie.Value)

}

func searchHandle(wr http.ResponseWriter, req *http.Request) {

	var sReq = []byte(`{
		"searchReq": "yandex",
		"sites": [
			"345345345https://yandex.ru",
			"https://golang.org",
			"https://google.com",
			"https://github.com",
			"https://dtf.ru",
			"https://geekbrains.ru"
		]
	}`)

	type searchRequest struct {
		SearchReq string   `json:"searchReq"`
		Sites     []string `json:"sites"`
	}

	sR := searchRequest{}
	if err := json.Unmarshal(sReq, &sR); err != nil {
		log.Println(err)
	}

	result := searchRequest{
		SearchReq: sR.SearchReq,
		Sites:     search(sR.SearchReq, sR.Sites),
	}
	resultJSON, err := json.Marshal(result)
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("%s", resultJSON)

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
