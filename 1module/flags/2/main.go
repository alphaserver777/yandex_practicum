// go run 1module\flags\1\main.go --help
// go run 1module\flags\1\main.go -file=<FILENAME>
// go run 1module\flags\2\main.go filter -gray true -sepia false
// go run 1module\flags\2\main.go cnv -thumb -w 2048  -dest "/home/user/images"
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	twoSubcommandFlags()
	//myFlags()
}

func twoSubcommandFlags() {
	// декларируем наборы флагов для подкоманд
	cnvFlags := flag.NewFlagSet("cnv", flag.ExitOnError)
	filterFlags := flag.NewFlagSet("filter", flag.ExitOnError)
	// декларируем флаги набора cnvFlags
	destDir := cnvFlags.String("dest", "./output", "destination folder")
	width := cnvFlags.Int("w", 1024, "width of the image")
	isThumb := cnvFlags.Bool("thumb", false, "create thumb")

	// флаги набора filterFlags
	isGray := filterFlags.Bool("gray", false, "convert to grayscale")
	isSepia := filterFlags.Bool("sepia", false, "convert to sepia")
	// проверяем, задана ли подкоманда
	// os.Arg[0] имя команды
	// os.Arg[1] имя подкоманды
	if len(os.Args) < 2 {
		fmt.Println("set or get subcommand required")
		os.Exit(1)
	}
	// в зависимости от переданной подкоманды
	// делаем парсинг флагов соответствующего набора,
	// передаём функции FlagSet.Parse() аргументы командной строки
	// os.Args[2:] содержит все аргументы,
	// следующие за os.Args[1], за именем подкоманды
	switch os.Args[1] {
	case "cnv":
		cnvFlags.Parse(os.Args[2:])
	case "filter":
		filterFlags.Parse(os.Args[2:])
	default:
		// PrintDefaults выводит параметры командной строки
		flag.PrintDefaults()
		os.Exit(1)
	}
	// проверяем, какой набор флагов использовался,
	// то есть какая подкоманда была передана,
	// функция FlagSet.Parsed() возвращает false, если
	// парсинг флагов набора не проводился
	if cnvFlags.Parsed() {
		fmt.Println("Destination folder:", *destDir)
		fmt.Println("Width:", *width)
		fmt.Println("Thumbs:", *isThumb)
	}
	if filterFlags.Parsed() {
		fmt.Println("isGrai:", *isGray)
		fmt.Println("isSepia:", *isSepia)
	}
}

func myFlags() {
	imgFile := flag.String("file", "", "input image file")
	destDir := flag.String("dest", "./output", "destination folder")
	width := flag.Int("w", 1024, "width of the image")
	isThumb := flag.Bool("thumb", false, "create thumb")

	// разбор командной строки
	flag.Parse()
	fmt.Println("Image file:", *imgFile)
	fmt.Println("Destination folder:", *destDir)
	fmt.Println("Width:", *width)
	fmt.Println("Thumbs:", *isThumb)
}
