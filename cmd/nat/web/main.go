package main

import (
	"encoding/json"
	"log"
	"net/http"
)

const (
	LocalServerAddr = ":7700"
)

func main() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		q := request.URL.Query()
		b, err := json.Marshal(q)
		if err != nil {
			log.Printf("Marshal Error：%v", err)
		}
		writer.Write(b)
	})
	log.Println("本地服务启动", LocalServerAddr)
	http.ListenAndServe(LocalServerAddr, nil)
}
