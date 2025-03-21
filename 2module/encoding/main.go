package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"time"
)

/*
Тег omitempty стоит использовать с осторожностью.
Если поле company отсутствует, то при обработке полученных данных
другая программа может возвращать ошибку.
А если значение этого поля равно пустой строке,
то программа будет работать.
*/

type MyType struct {
	User      string    `json:"user,omitempty" example:"Bob"`
	CreatedAt time.Time `json:"created_at"`
}

type Visitor struct {
	ID      int      `json:"id"`
	Name    string   `json:"name,omitempty"`
	Phones  []string `json:"phones,omitempty"`
	Company string   `json:"company,omitempty"`
}

type Data struct {
	ID      int    `json:"-"`
	Name    string `json:"name"`
	Company string `json:"comp,omitempty"`
}

var visitors = map[string]Visitor{
	"1": {
		ID:   1,
		Name: "Guest",
		Phones: []string{
			`789-673-56-90`,
			`612-934-77-23`,
		},
		Company: "KHMS",
	},
	"2": {
		ID:   2,
		Name: "Slava",
		Phones: []string{
			`789-673-56-90`,
			`612-934-77-23`,
		},
		Company: "Ozon",
	},
}

const (
	targetField = "User" // имя поля, о котором нужно получить информацию
	targetTag   = "json" // тег, значение которого нужно получить
)

func main() {
	// validData()
	go func() {
		time.Sleep(time.Second)
		http.Post(`http://localhost:8080`, `application/json`,
			// ключи указаны в разных регистрах, но данные сконвертируются правильно
			bytes.NewBufferString(`{"ID": 10, "NaMe": "Gopher", "company": "Don't Panic"}`))
	}()
	http.ListenAndServe("localhost:8080", http.HandlerFunc(JSONHandler)) // "C:\Program Files\Git\mingw64\bin\curl.exe" --include localhost:8080/?id=1
	// task1()
}

func validData() {

	obj := MyType{}

	// получаем Go-описание типа
	objType := reflect.TypeOf(obj)

	// ищем поле по имени
	field, ok := objType.FieldByName(targetField)
	if !ok {
		panic(fmt.Errorf("field (%s): not found", targetField))
	}

	// ищем тег по имени
	tagValue, ok := field.Tag.Lookup(targetTag)
	if !ok {
		panic(fmt.Errorf("tag (%s) for field (%s): not found", targetTag, targetField))
	}

	fmt.Printf("Значение тега %s поля %s: %s\n", targetTag, targetField, tagValue)
	fmt.Printf("Теги поля %s: %s\n", targetField, field.Tag)
}

func JSONHandler(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	resp, err := json.Marshal(visitors[id])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func JSONPostData(w http.ResponseWriter, req *http.Request) {
	var id string

	if req.Method == http.MethodPost {
		var visitor Visitor
		var buf bytes.Buffer
		// читаем тело запроса
		_, err := buf.ReadFrom(req.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// десериализуем JSON в Visitor
		if err = json.Unmarshal(buf.Bytes(), &visitor); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		id = strconv.Itoa(visitor.ID)
		// добавляем в мапу
		visitors[id] = visitor
	} else {
		id = req.URL.Query().Get("id")
	}
	resp, err := json.Marshal(visitors[id])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func task1() {
	foo := []Data{
		{
			ID:   10,
			Name: "Gopher",
		},
		{
			Name:    "Вася",
			Company: "Яндекс",
		},
	}
	out, err := json.Marshal(foo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))

}
