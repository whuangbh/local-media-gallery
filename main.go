package main

import (
	"bufio"
	"fmt"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"local-media-gallery/config"
	"local-media-gallery/database"
	"local-media-gallery/scripts"
	"local-media-gallery/server"
	"log"
	"os"
	"strings"

	_ "golang.org/x/image/webp"
)

var MEDIA_HOST_PATH string

func main() {
	config.LoadConfig()

	err := database.InIt()
	if err != nil {
		log.Fatalf("error trying initialize database: \n %v \n", err)
	}

	defer database.CloseConnection()
	log.Println("Database initialized successfully.")

	input, err := handleInput()
	if err != nil {
		log.Fatalf("error getting user input: %v", err)
	}

	if input != "1" && input != "2" {
		fmt.Println("Please enter 1 or 2!")
		return
	}

	if input == "2" {
		log.Println("Start preprocessing media folders...")
		scripts.PreprocessMediaFolder()
	}

	log.Println("Start hosting web server...")
	if err := server.Init(); err != nil {
		log.Fatalf("error trying to start web server: \n %v \n", err)
	}
}

func handleInput() (string, error) {
	fmt.Println("What do you want to do?")
	fmt.Println("1: Start web server")
	fmt.Println("2: Preprocess media folders, and start web server")
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("error reading input: %w", err)
	}

	input = strings.TrimSpace(input)
	return input, nil
}
