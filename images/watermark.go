/*
Copyright 2018 Rémi Lavedrine.

Licensed under the Mozilla Public License, version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

https://www.mozilla.org/en-US/MPL/

* The above copyright notice and this permission notice shall be included in all
* copies or substantial portions of the Software.

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package images : Provides manipulation functions on images.
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
