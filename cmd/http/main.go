package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/josemyduarte/printer"
)

var assets printer.Assets

func init() {
	assets = printer.Assets{
		BgImgPath: "assets/00-instagram-background.png",
		FontPath:  "assets/FiraSans-Light.ttf",
		FontSize:  60,
	}
}

func main() {
	http.HandleFunc("/", assets.Serve)

	addr := determineListenAddress()
	log.Printf("Listening on %s...\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}

func determineListenAddress() string {
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("$PORT not set, using :80 as default")
		return ":80"
	}
	return ":" + port
}
