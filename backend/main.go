package main

import (
	"fmt"
	_ "golang.org/x/image/webp"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"local-media-gallery/database"
	"local-media-gallery/scripts"
	"local-media-gallery/server"
	"log"
	"os"
)

func main() {
	err := database.InIt()
	if err != nil {
		log.Fatalf("error trying initialize database: \n %v \n", err)
	}

	defer database.CloseConnection()
	log.Println("Database initialized successfully.")

	appMode := os.Getenv("APP_MODE")

	switch appMode {
	case "web_server":
		log.Println("Start hosting web server...")
		if err := server.Init(); err != nil {
			log.Fatalf("error trying to start web server: \n %v \n", err)
		}
	case "preprocess":
		log.Println("Start preprocessing media folders...")
		scripts.PreprocessMediaFolder()
	default:
		fmt.Printf("Unknown appMode: %s! \n", appMode)
		return
	}
}
