package network

import (
	"fmt"
	"io"
	"net/http"
)

func CreateHttpServer() {
	http.HandleFunc("/go", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(request.RemoteAddr, "connection is ok")
		fmt.Println("method:", request.Method)
		fmt.Println("url:", request.URL.Path)
		fmt.Println("header:", request.Header)
		fmt.Println("body:", request.Body)

		writer.Write([]byte("i got it"))
	})
	err := http.ListenAndServe("127.0.0.1:7890", nil)
	if err != nil {
		fmt.Println("http server start error:", err)
		return
	}
}

func CreateHttpClient() {
	get, err := http.Get("http://127.0.0.1:7890/go")
	if err != nil {
		fmt.Println("http get is failed,err:", err)
		return
	}
	defer get.Body.Close()

	fmt.Println("http status", get.Status)
	fmt.Println("http response header", get.Header)

	buf := make([]byte, 1024)

	for {
		n, err := get.Body.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("read http Get body error:", err)
			return
		} else {
			fmt.Println("response data:", string(buf[:n]))
			break
		}
	}
}
