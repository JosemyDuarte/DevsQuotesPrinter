package printer

import (
	"bytes"
	"encoding/json"
	"image/png"
	"log"
	"net/http"
	"strconv"
)

type Request struct {
	BgImgPath string
	FontPath  string
	FontSize  float64
	Text      string
}

type Assets struct {
	BgImgPath string
	FontPath  string
	FontSize  float64
}

//Serve writes on the image referenced on Assets.BgImgPath with the font set on Assets.FontPath and
//the size Assets.FontSize the text present on the http.Request.
func (a *Assets) Serve(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Text string `json:"text"`
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	log.Println("received request[", req, "]")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	img, err := TextOnImg(
		Request{
			BgImgPath: a.BgImgPath,
			FontPath:  a.FontPath,
			FontSize:  a.FontSize,
			Text:      req.Text,
		},
	)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	buffer := new(bytes.Buffer)
	if err := png.Encode(buffer, img); err != nil {
		log.Println("unable to encode image")
	}

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))

	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Println("couldn't write image to response")
	}
}
