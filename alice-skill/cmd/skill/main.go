package main

import (
	"log"
	"net/http"
	"yandex_practicum/alice-skill/internal/flags"
	"yandex_practicum/alice-skill/internal/logger"

	"go.uber.org/zap"
)

func main() {
	// Обрабатываем аргументы командной строки
	flags.ParseFlags()

	// Запускаем сервер
	if err := run(); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}

// Функция для запуска сервера
func run() error {
	if err := logger.Initialize(flags.FlagLogLevel); err != nil {
		return err
	}
	logger.Log.Info("Running server", zap.String("address", flags.FlagRunAddr))
	// оборачиваем хендлер webhook в middleware с логированием
	return http.ListenAndServe(flags.FlagRunAddr, logger.RequestLogger(webhook))
}

// Обработчик вебхука
func webhook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		logger.Log.Debug("got request with bad method", zap.String("method", r.Method))
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(`
      {
        "response": {
          "text": "Извините, я пока ничего не умею"
        },
        "version": "1.1 with logger"
      }
    `))
	logger.Log.Debug("Sending HTTP 200 response")
}
