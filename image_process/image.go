package image_process

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"strings"
	"sync"
)

// path relative to the current file.
func Test() {
	// img, err := OpenImage("images/elysia.png")
	// if err != nil {
	// 	fmt.Println("Error when opening file:", err)
	// }

	// imgColor := GetImageTensor(img)
	// newImg := ConvertGreyScale(&imgColor)

	// SaveImage(newImg, "images/elysiaGrey.png")

	// img, err = OpenImage("images/griseo.jpg")
	// if err != nil {
	// 	fmt.Println("Error when opening file:", err)
	// }

	// imgColor = GetImageTensor(img)
	// newImg = ConvertGreyScale(&imgColor)

	// SaveImage(newImg, "images/griseoGrey.jpg")

	imgWidth := 28      // integer pixels
	imgHeight := 28 * 5 //integer pixels
	bgColor := image.White
	rgba := image.NewRGBA(image.Rect(0, 0, imgWidth, imgHeight))
	draw.Draw(rgba, rgba.Bounds(), bgColor, image.Pt(28, 28), draw.Src)

	f, err := os.Create("images/idk.png")
	if err != nil {
		fmt.Println("error:", err)
	}
	defer f.Close()

	err = png.Encode(f, rgba)
	if err != nil {
		fmt.Println("error:", err)
	}

	m := image.NewRGBA(image.Rect(0, 0, 280, 280))
	blue := color.RGBA{0, 0, 255, 255}
	draw.Draw(m, m.Bounds(), &image.Uniform{blue}, image.Point{}, draw.Src)

	draw.Draw(m, m.Bounds(), rgba, image.Point{28 * -3, 28 * -2}, draw.Src)

	f, err = os.Create("images/blue.png")
	if err != nil {
		fmt.Println("error:", err)
	}
	defer f.Close()

	err = png.Encode(f, m)
	if err != nil {
		fmt.Println("error:", err)
	}
}

func SaveImage(img image.Image, path string) {
	if strings.Contains(path, ".png") {
		f, err := os.Create(path)
		if err != nil {
			fmt.Println("error:", err)
		}
		defer f.Close()

		err = png.Encode(f, img)
		if err != nil {
			fmt.Println("error:", err)
		}

		return
	}

	if strings.Contains(path, ".jpeg") || strings.Contains(path, ".jpg") {
		f, err := os.Create(path)
		if err != nil {
			fmt.Println("error:", err)
		}
		defer f.Close()

		opt := jpeg.Options{
			Quality: 90,
		}

		err = jpeg.Encode(f, img, &opt)
		if err != nil {
			fmt.Println("error:", err)
		}

		return
	}
}

func GetNewImage(path string) image.Image {
	img, err := OpenImage("elysia.png")
	if err != nil {
		fmt.Println("Error when opening file:", err)
	}

	imgColor := GetImageTensor(img)
	newImg := ConvertGreyScale(&imgColor)

	return newImg
}

func OpenImage(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer f.Close()

	img, format, err := image.Decode(f)
	if err != nil {
		e := fmt.Errorf("error in decoding %w", err)
		return nil, e
	}

	if format != "jpeg" && format != "png" {
		e := fmt.Errorf("error in image format - not jpeg")
		return nil, e
	}

	return img, nil
}

// store all color pixels in a vector
func GetImageTensor(img image.Image) (pixels [][]color.Color) {
	size := img.Bounds().Size()
	for i := 0; i < size.X; i++ {
		var y []color.Color
		for j := 0; j < size.Y; j++ {
			y = append(y, img.At(i, j))
		}
		pixels = append(pixels, y)
	}

	return
}

func ConvertGreyScale(pixels *[][]color.Color) image.Image {
	p := *pixels
	wg := sync.WaitGroup{}
	rect := image.Rect(0, 0, len(p), len(p[0]))
	newImage := image.NewRGBA(rect)
	for x := 0; x < len(p); x++ {
		for y := 0; y < len(p[0]); y++ {
			wg.Add(1)
			go func(x, y int) {
				pix := p[x][y]
				originalColor, ok := color.RGBAModel.Convert(pix).(color.RGBA)
				if !ok {
					log.Fatalf("color.color conversion went wrong")
				}
				grey := uint8(float64(originalColor.R)*0.21 + float64(originalColor.G)*0.72 + float64(originalColor.B)*0.07)
				col := color.RGBA{
					grey,
					grey,
					grey,
					originalColor.A,
				}
				newImage.Set(x, y, col)
				wg.Done()
			}(x, y)
		}
	}
	return newImage
}
