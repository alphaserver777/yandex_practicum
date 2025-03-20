package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

func main() {
	// logLogrus("logrus.log")
	// DifferentLevels(log.InfoLevel)
	zapLoger()
}

func logLogrus(nameLogFile string) {
	// создаём файл info.log и обрабатываем ошибку
	file, err := os.OpenFile(nameLogFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// откладываем закрытие файла
	defer file.Close()

	// устанавливаем вывод логов в файл
	log.SetOutput(file)
	// устанавливаем вывод логов в формате JSON
	log.SetFormatter(&log.JSONFormatter{})
	// устанавливаем уровень предупреждений
	log.SetLevel(log.WarnLevel) // chose different LEVEL -> log.InfoLevel log.WarnLevel etc

	// определяем стандартные поля JSON
	log.WithFields(log.Fields{
		"genre": "metal",
		"name":  "Rammstein",
	}).Info("Немецкая метал-группа, образованная в январе 1994 года в Берлине.")

	log.WithFields(log.Fields{
		"omg":  true,
		"name": "Garbage",
	}).Warn("В 2021 году вышел новый альбом No Gods No Masters.")

	log.WithFields(log.Fields{
		"omg":  true,
		"name": "Linkin Park",
	}).Fatal("Группа Linkin Park взяла паузу после смерти вокалиста Честера Беннингтона 20 июля 2017 года.")
}

// chose different LEVEL -> log.InfoLevel log.WarnLevel etc
func DifferentLevels(level log.Level) {
	log.SetOutput(os.Stdout)
	log.SetLevel(level)

	log.WithFields(log.Fields{
		"genre": "metal",
		"name":  "Rammstein",
	}).Info("Ich Will")

	log.WithFields(log.Fields{
		"genre": "post-grunge",
		"name":  "Garbage",
	}).Warn("I Think I’m Paranoid")

	contextLogger := log.WithFields(log.Fields{
		"common": "Any music is awesome",
		"other":  "I also should be logged always",
	})

	contextLogger.Warn("I will be logged with common and other fields")
	contextLogger.Error("Me too, maybe")

	log.WithFields(log.Fields{
		"genre": "rock",
		"name":  "The Rasmus",
	}).Fatal("Livin' in a World Without You")
}

func zapLoger() {
	// добавляем предустановленный логер NewDevelopment
	logger, err := zap.NewDevelopment()
	if err != nil {
		// вызываем панику, если ошибка
		panic("cannot initialize zap")
	}
	// это нужно добавить, если логер буферизован
	// в данном случае не буферизован, но привычка хорошая
	defer logger.Sync()

	// для примера берём простой URL
	const url = "http://example.com"

	// делаем логер SugaredLogger
	sugar := logger.Sugar()

	// выводим сообщение уровня Info с парой "url": url в виде JSON, это SugaredLogger
	sugar.Infow(
		"Failed to fetch URL",
		"url", url,
	)

	// выводим сообщение уровня Info, но со строкой URL, это тоже SugaredLogger
	sugar.Infof("Failed to fetch URL: %s", url)
	// выводим сообщение уровня Error со строкой URL, и это SugaredLogger
	sugar.Errorf("Failed to fetch URL: %s", url)

	// переводим в обычный Logger
	plain := sugar.Desugar()

	// выводим сообщение уровня Info обычного регистратора (не SugaredLogger)
	plain.Info("Hello, Go!")
	// также уровня Warn (не SugaredLogger)
	plain.Warn("Simple warning")
	// и уровня Error, но добавляем строго типизированное поле "url" (не SugaredLogger)
	plain.Error("Failed to fetch URL", zap.String("url", url))
}
