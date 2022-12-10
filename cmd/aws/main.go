package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image/jpeg"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/josemyduarte/printer/internal/printer"
)

var (
	fontURL      = "https://github.com/DevsQuotes/DevsQuotesPrinter/raw/master/assets/FiraSans-Light.ttf"
	fontFileName = "/tmp/FiraSans-Light.ttf"

	bgURL      = "https://github.com/DevsQuotes/DevsQuotesPrinter/raw/master/assets/00-instagram-background.png"
	bgFileName = "/tmp/00-instagram-background.png"
)

// It seems that in AWS lambdas you don't have access to the asset folder.
// Downloading the assets on init to make sure we have it available on runtime.
func init() {
	if err := downloadFile(fontURL, fontFileName); err != nil {
		panic(fmt.Errorf("couldn't download font: %w", err))
	}

	if err := downloadFile(bgURL, bgFileName); err != nil {
		panic(fmt.Errorf("couldn't download background image: %w", err))
	}
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var req struct {
		Text string `json:"text"`
	}

	b := []byte(request.Body)
	err := json.Unmarshal(b, &req)
	if err != nil {
		log.Println(err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       err.Error(),
		}, err
	}

	img, err := printer.TextOnImg(printer.Request{
		BgImgPath: bgFileName,
		FontPath:  fontFileName,
		FontSize:  60,
		Text:      req.Text,
	})
	if err != nil {
		log.Println(err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, err
	}

	imgBuf := new(bytes.Buffer)
	if jpeg.Encode(imgBuf, img, nil) != nil {
		log.Println(err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode:      http.StatusOK,
		Body:            base64.StdEncoding.EncodeToString(imgBuf.Bytes()),
		IsBase64Encoded: true,
		Headers: map[string]string{
			"Content-Type": "image/png",
		},
	}, nil

}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
