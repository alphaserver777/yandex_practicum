package main

import (
	"fmt"
	"net/http"
)

func mainPage(res http.ResponseWriter, req *http.Request) {
	body := fmt.Sprintf("Method: %s\r\n", req.Method)
	/*
		Создает строку body, которая будет содержать информацию о запросе.
		req.Method возвращает метод HTTP-запроса (например, GET, POST, PUT).
		fmt.Sprintf форматирует строку, подставляя значение req.Method.
	*/
	body += "Header ===============\r\n"
	for k, v := range req.Header {
		body += fmt.Sprintf("%s: %v\r\n", k, v)
	}
	/*
		Добавляет в body заголовки запроса.
		req.Header — это карта (map), где ключи — это названия заголовков, а значения — списки строк.
		Цикл for k, v := range req.Header проходит по всем заголовкам и добавляет их в body
	*/

	body += "Query parameters ===============\r\n"
	for k, v := range req.URL.Query() {
		body += fmt.Sprintf("%s: %v\r\n", k, v)
	}
	res.Write([]byte(body))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc(`/`, mainPage)

	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}
