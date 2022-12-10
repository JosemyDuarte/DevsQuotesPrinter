package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/josemyduarte/printer/internal/handler"
)

func determineListenAddress() string {
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("$PORT not set, using :80 as default")
		return ":80"
	}
	return ":" + port
}

func main() {
	httpHandler := handler.HTTP{
		BackgroundImgPath: "assets/00-instagram-background.png",
		FontPath:          "assets/FiraSans-Light.ttf",
		FontSize:          60,
	}

	http.HandleFunc("/", httpHandler.Handle)

	addr := determineListenAddress()
	log.Printf("Listening on %s...\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}
