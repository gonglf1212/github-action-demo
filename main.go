package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"net/http"
)

func main() {

	fmt.Println("go服务启动了，启动时间：", time.Now().Format("2006-1-2 15:04:05"))
	HttpGet()
}

func HttpGet() {
	resp, err := http.Get("http://baidu.com/")
	if err != nil {
		// 处理错误
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// 处理错误
		panic(err)
	}

	fmt.Println(string(body))
}

// Http接口估计外面访问不了，死心了
func RunHttp() {
	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}
