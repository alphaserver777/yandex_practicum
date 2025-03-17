// go run 1module\flags\1\main.go --help
// go run 1module\flags\1\main.go -file=<FILENAME>

package main

import (
	"flag"
	"fmt"
)

func main() {
	// указываем имя флага, значение по умолчанию и описание
	imgFile := flag.String("file", "", "input image file")
	// делаем разбор командной строки
	flag.Parse()
	fmt.Println("Image file:", *imgFile)
}
