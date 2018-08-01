package images

import (
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log"
	"os"

	"github.com/shosta/androSecTest/logging"

	extension "github.com/shosta/androSecTest/file"
)

func decodeSrcImage(imgSrcPath string) image.Image {
	imgSrc, err := os.Open(imgSrcPath)
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}
	defer imgSrc.Close()

	var srcImg image.Image
	if extension.IsPNG(imgSrcPath) {
		src, err := png.Decode(imgSrc)
		if err != nil {
			log.Fatalf("failed to decode: %s", err)
		}
		srcImg = src
	} else if extension.IsJPG(imgSrcPath) {
		src, err := jpeg.Decode(imgSrc)
		if err != nil {
			log.Fatalf("failed to decode: %s", err)
		}
		srcImg = src
	}

	return srcImg
}

// Watermark :
func Watermark(watermarkPath string, imgSrcPath string) {
	// Open and create the Source image
	logging.PrintlnDebug("Decode image : " + imgSrcPath)
	src := decodeSrcImage(imgSrcPath)

	// Open and create the watermark image
	watermarkImg, err := os.Open(watermarkPath)
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}

	watermark, err := png.Decode(watermarkImg)
	if err != nil {
		log.Fatalf("failed to decode: %s", err)
	}
	defer watermarkImg.Close()

	// Add watermak on Source image
	offset := image.Pt(2, 2)
	b := src.Bounds()
	image3 := image.NewRGBA(b)
	draw.Draw(image3, b, src, image.ZP, draw.Src)
	draw.Draw(image3, watermark.Bounds().Add(offset), watermark, image.ZP, draw.Over)

	os.RemoveAll(imgSrcPath)
	third, err := os.Create(imgSrcPath)
	if err != nil {
		log.Fatalf("failed to create: %s", err)
	}
	// jpeg.Encode(third, image3, &jpeg.Options{jpeg.DefaultQuality})
	png.Encode(third, image3)
	defer third.Close()
}
