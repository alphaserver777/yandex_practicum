package flags

import (
	"flag"
	"os"
)

// неэкспортированная переменная FlagRunAddr содержит адрес и порт для запуска сервера
var (
	FlagRunAddr  string
	FlagLogLevel string
)

// parseFlags обрабатывает аргументы командной строки
// и сохраняет их значения в соответствующих переменных
func ParseFlags() {
	flag.StringVar(&FlagRunAddr, "a", ":8080", "addres and port to run server")
	flag.StringVar(&FlagLogLevel, "l", "info", "log level")
	flag.Parse()

	if envRunAddr := os.Getenv("RUN_ADDR"); envRunAddr != "" {
		FlagRunAddr = envRunAddr
	}
	if envLogLevel := os.Getenv("LOG_LEVEL"); envLogLevel != "" {
		FlagLogLevel = envLogLevel
	}
}
