package main

import (
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"net/http"
)

func main() {
	http.HandleFunc("/", httpHandler)
	http.ListenAndServe(":8080", nil)
}

// Gaussian noise with variance

func createGaussianNoiseImage(width, height, variance int) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, color.RGBA{
				R: uint8(rand.NormFloat64()*float64(variance) + 128),
				G: uint8(rand.NormFloat64()*float64(variance) + 128),
				B: uint8(rand.NormFloat64()*float64(variance) + 128),
				A: 255,
			})
		}
	}
	return img
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	img := createGaussianNoiseImage(512, 512, 64)
	w.Header().Set("Content-Type", "image/png")
	png.Encode(w, img)
}
