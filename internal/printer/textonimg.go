package printer

import (
	"image"
	"image/color"

	"github.com/fogleman/gg"
)

type Request struct {
	BgImgPath string
	FontPath  string
	FontSize  float64
	Text      string
}

// TextOnImg given a path to an image, a font & size and a text it will return an Image with the
// given text printed in the ~middle of the image
func TextOnImg(request Request) (image.Image, error) {
	bgImage, err := gg.LoadImage(request.BgImgPath)
	if err != nil {
		return nil, err
	}
	imgWidth := bgImage.Bounds().Dx()
	imgHeight := bgImage.Bounds().Dy()

	dc := gg.NewContext(imgWidth, imgHeight)
	dc.DrawImage(bgImage, 0, 0)

	if err := dc.LoadFontFace(request.FontPath, request.FontSize); err != nil {
		return nil, err
	}

	x := float64(imgWidth / 2)
	signatureOff := 80
	y := float64((imgHeight / 2) - signatureOff)
	textMargin := 60.0
	maxWidth := float64(imgWidth) - textMargin
	dc.SetColor(color.Black)
	dc.DrawStringWrapped(request.Text, x+3, y+3, 0.5, 0.5, maxWidth, 1.5, gg.AlignCenter)
	dc.SetColor(color.White)
	dc.DrawStringWrapped(request.Text, x, y, 0.5, 0.5, maxWidth, 1.5, gg.AlignCenter)

	return dc.Image(), nil
}

func Save(img image.Image, path string) error {
	if err := gg.SavePNG(path, img); err != nil {
		return err
	}
	return nil
}
