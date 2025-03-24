package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"yandex_practicum/alice-skill/internal/flags"
	"yandex_practicum/alice-skill/internal/logger"
	"yandex_practicum/alice-skill/internal/models"

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
	http.Handle("/ping", pingHandler())
	// десериализуем запрос в структуру модели
	logger.Log.Debug("decoding request")
	var req models.Request
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&req); err != nil {
		logger.Log.Debug("cannot decode request JSON body", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// проверяем, что пришёл запрос понятного типа
	if req.Request.Type != models.TypeSimpleUtterance {
		logger.Log.Debug("unsupported request type", zap.String("type", req.Request.Type))
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	// заполняем модель ответа
	resp := models.Response{
		Response: models.ResponsePayload{
			Text: "Извините, я пока ничего не умею",
		},
		Version: "1.0",
	}

	w.Header().Set("Content-Type", "application/json")

	// сериализуем ответ сервера
	enc := json.NewEncoder(w)
	if err := enc.Encode(resp); err != nil {
		logger.Log.Debug("error encoding response", zap.Error(err))
		return
	}
	logger.Log.Debug("sending HTTP 200 response")

}

func pingHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprintf(w, "pong\n")
	}
	return http.HandlerFunc(fn)
}
