package main

import "net/http"

type MyHandler struct{}

func (h MyHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	data := []byte("Привет! Сервер находится в разработке")
	res.Write(data)
}

func main() {
	var h MyHandler

	err := http.ListenAndServe(`:8080`, h)
	if err != nil {
		panic(err)
	}
}
