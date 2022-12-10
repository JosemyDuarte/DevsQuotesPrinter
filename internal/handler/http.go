package handler

import (
	"bytes"
	"encoding/json"
	"image/png"
	"log"
	"net/http"
	"strconv"

	"github.com/josemyduarte/printer/internal/printer"
)

var (
	BgImgPath         = "assets/00-instagram-background.png"
	FontPath          = "assets/FiraSans-Light.ttf"
	FontSize  float64 = 60
)

// Serve writes on the image referenced on Assets.BgImgPath with the font set on Assets.FontPath and
// the size Assets.FontSize the text present on the http.Request.
func Serve(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Text string `json:"text"`
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	log.Println("received request[", req, "]")
	if err != nil {
		log.Printf("failed parsing request: %w", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	img, err := printer.TextOnImg(
		printer.Request{
			BgImgPath: BgImgPath,
			FontPath:  FontPath,
			FontSize:  FontSize,
			Text:      req.Text,
		},
	)
	if err != nil {
		log.Printf("couldn't print text on image: %w", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	buffer := new(bytes.Buffer)
	if err := png.Encode(buffer, img); err != nil {
		log.Printf("unable to encode image: %w", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))

	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Printf("couldn't write image to response: %w", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
