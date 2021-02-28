package printer

import (
	"image"
	"image/color"

	"github.com/fogleman/gg"
)

//TextOnImg 
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
	textShadowColor := color.Black
	textColor := color.White

	textMargin := 60.0
	x := float64(imgWidth / 2)
	y := float64((imgHeight / 2) - 80)
	maxWidth := float64(dc.Width()) - textMargin
	dc.SetColor(textShadowColor)
	dc.DrawStringWrapped(request.Text, x+3, y+3, 0.5, 0.5, maxWidth, 1.5, gg.AlignCenter)
	dc.SetColor(textColor)
	dc.DrawStringWrapped(request.Text, x, y, 0.5, 0.5, maxWidth, 1.5, gg.AlignCenter)

	return dc.Image(), nil
}

func Save(img image.Image, path string) error {
	if err := gg.SavePNG(path, img); err != nil {
		return err
	}
	return nil
}
