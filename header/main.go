package main

import (
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	client := &http.Client{}
	request, err := http.NewRequest("GET", "http://localhost:18888", nil)
	request.Header.Add("Content-Type", "image/jpeg")
	// BASIC認証
	request.SetBasicAuth("ユーザ", "パスワード")
	// クッキー
	request.AddCookie(&http.Cookie{Name: "test", Value: "value"})
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	log.Println(string(dump))
}
