package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	// logLogrus("logrus.log")
	// DifferentLevels()
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
	log.SetLevel(log.WarnLevel)

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
