package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	// myFirstLog()
	// logInFile("info1.log")
	// logInFileNew("server.log")
	// logSetFlags("setFlags.log")
	// logTask1()
}

func myFirstLog() {
	log.Print("Log in Go")
}

func logInFile(nameLogFile string) {
	// создаём файл info.log и обрабатываем ошибку, если что-то пошло не так
	file, err := os.OpenFile(nameLogFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	// откладываем закрытие файла
	defer file.Close()

	// устанавливаем назначение вывода в файл info.log
	log.SetOutput(file)
	log.Print("Logging to a file in Go!")
}

func logInFileNew(nameLogFile string) {
	flog, err := os.OpenFile(nameLogFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer flog.Close()

	myLog := log.New(flog, "serv: ", log.LstdFlags|log.Lshortfile)
	myLog.Print("Start server")
	myLog.Print("Finish server")
}

func logSetFlags(nameLogFile string) {
	// создаём файл info.log и обрабатываем ошибку, если что-то пошло не так
	file, err := os.OpenFile(nameLogFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// откладываем закрытие файла
	defer file.Close()

	// устанавливаем назначение вывода в файл info.log
	log.SetOutput(file)

	// добавляем флаги для вывода даты, времени, имени файла
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Print("Logging to a file in Go!")
}

func logTask1() {
	/*
		Создайте переменную типа *log.Logger,
		которая будет использовать bytes.Buffer для записи данных.
		В результате работы программы в буфере должно быть две строки произвольного содержания
		можете использовать, например, названия любимых групп или песен.
	*/

	var buf bytes.Buffer
	// допишите код
	// 1) создайте переменную типа *log.Logger
	// 2) запишите в неё нужные строки

	logger := log.New(&buf, "mylog: ", 0) // *log.Logger ...
	logger.Print("Hello, world")
	logger.Print("Goodbye")

	fmt.Print(&buf)
	// должна вывести
	// mylog: Hello, world!
	// mylog: Goodbye
}
