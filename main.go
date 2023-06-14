package main

import (
	"fmt"
	"log"
	"time"

	"net/http"
)

func main() {

	fmt.Println("go服务启动了，启动时间：", time.Now().Format("2006-1-2 15:04:05"))
	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}
