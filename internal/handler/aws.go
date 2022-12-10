package handler

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"image/png"
	"log"
	"net/http"
	"strconv"

	"github.com/aws/aws-lambda-go/events"

	"github.com/josemyduarte/printer/internal/printer"
)

type AWS struct {
	BackgroundImgPath string
	FontPath          string
	FontSize          float64
}

func (a *AWS) Serve(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("request received [%+v]", request)

	var ParsedRequest struct {
		Text string `json:"text"`
	}

	b := []byte(request.Body)
	err := json.Unmarshal(b, &ParsedRequest)
	if err != nil {
		log.Printf("failed to parse request body: %w", err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       err.Error(),
		}, err
	}

	log.Printf("text = [%+v]", ParsedRequest.Text)

	img, err := printer.TextOnImg(printer.Request{
		BgImgPath: a.BackgroundImgPath,
		FontPath:  a.FontPath,
		FontSize:  a.FontSize,
		Text:      ParsedRequest.Text,
	})
	if err != nil {
		log.Printf("failed to print text on image: %w", err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, err
	}

	log.Printf("printed %s on image", ParsedRequest.Text)

	imgBuf := new(bytes.Buffer)
	if png.Encode(imgBuf, img) != nil {
		log.Printf("failed to encode image bytes: %w", err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, err
	}

	log.Printf("image encoded")

	return events.APIGatewayProxyResponse{
		StatusCode:      http.StatusOK,
		Body:            base64.StdEncoding.EncodeToString(imgBuf.Bytes()),
		IsBase64Encoded: true,
		Headers: map[string]string{
			"Content-Type":   "image/png",
			"Content-Length": strconv.Itoa(len(imgBuf.Bytes())),
		},
	}, nil
}
