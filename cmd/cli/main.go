package main

import (
	"flag"
	"log"

	"github.com/josemyduarte/printer/internal/printer"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	var (
		fontSize          = flag.Float64("fontSize", 60, "font fontSize in points")
		fontPath          = flag.String("fontPath", "assets/FiraSans-Light.ttf", "filename of the ttf font")
		backgroundImgPath = flag.String("bgImg", "assets/00-instagram-background.png", "image to use as background")
		text              = flag.String("text", "NOTHING", "text to print on the image")
		outputPath        = flag.String("output", "cool_img.png", "output path for the resulting image")
	)
	flag.Parse()
	img, err := printer.TextOnImg(
		printer.Request{
			BgImgPath: *backgroundImgPath,
			FontPath:  *fontPath,
			FontSize:  *fontSize,
			Text:      *text,
		},
	)
	if err != nil {
		return err
	}

	if err := printer.Save(img, *outputPath); err != nil {
		return err
	}

	log.Println("image saved on [", *outputPath, "]")
	return nil
}
