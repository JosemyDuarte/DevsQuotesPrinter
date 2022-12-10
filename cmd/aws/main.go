package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"

	"github.com/josemyduarte/printer/internal/handler"
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

func main() {
	awsHandler := handler.AWS{
		BackgroundImgPath: bgFileName,
		FontPath:          fontFileName,
		FontSize:          60,
	}

	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(awsHandler.Serve)
}
