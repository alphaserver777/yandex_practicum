package main

import (
	"fmt"
	"log"
	"os"

	"github.com/caarlos0/env/v6"
)

func main() {
	// getUser()
	// getAllEnv()
	printEnvUser()
}

func getUser() {
	u := os.Getenv("USERNAME")
	log.Print(u)
}

func getAllEnv() {
	envList := os.Environ()
	for i := 0; i < len(envList); i++ {
		fmt.Println(envList[i])
	}
}

func printEnvUser() {

	type Config struct {
		User string `env:"USERNAME"`
	}

	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Current user is %s\n", cfg.User)
}
