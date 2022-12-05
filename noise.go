package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"net/http"
)

// generate noise png image
func main() {
	// create new image
	img := image.NewRGBA(image.Rect(0, 0, 1024, 1024))
	// fill with noise
	for x := 0; x < 1024; x++ {
		for y := 0; y < 1024; y++ {
			img.Set(x, y, color.RGBA{
				R: uint8(rand.Intn(255)),
				G: uint8(rand.Intn(255)),
				B: uint8(rand.Intn(255)),
				A: 255,
			})
		}
	}
	// encode to buffer
	var buf bytes.Buffer
	err := png.Encode(&buf, img)
	if err != nil {
		fmt.Println(err)
		return
	}
	// write to client
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/png")
		w.Write(buf.Bytes())
	})
	http.ListenAndServe(":8080", nil)
}
