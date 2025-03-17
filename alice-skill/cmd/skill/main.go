package main

import (
	"log"
	"net/http"
)

func main() {
	// Обрабатываем аргументы командной строки
	parseFlags()

	// Запускаем сервер
	if err := run(); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}

// Функция для запуска сервера
func run() error {
	log.Println("Сервер запущен на:", flagRunAddr)
	return http.ListenAndServe(flagRunAddr, http.HandlerFunc(webhook))
}

// Обработчик вебхука
func webhook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(`{"error": "Разрешены только POST-запросы"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write([]byte(`
      {
        "response": {
          "text": "Извините, я пока ничего не умею"
        },
        "version": "1.0"
      }
    `))
	if err != nil {
		log.Printf("Ошибка при отправке ответа: %v", err)
	}
}
