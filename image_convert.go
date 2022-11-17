package main

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"image/png"
	"net/http"
)

func main() {
	http.HandleFunc("/", serveConvertForm)
	http.ListenAndServe(":8080", nil)
}

func serveConvertForm(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`
			<form method="POST" enctype="multipart/form-data">
				<p>Image: <input type="file" name="image"></p>
				<input type="submit">
			</form>
		`))
		return
	}

	convertJpgToPng(w, r)
}

func convertJpgToPng(w http.ResponseWriter, r *http.Request) {
	// open "test.jpg"
	file, _, err := r.FormFile("image")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	// encode to buffer
	var buf bytes.Buffer
	err = png.Encode(&buf, img)
	if err != nil {
		fmt.Println(err)
		return
	}

	// write to client
	w.Header().Set("Content-Type", "image/png")
	w.Write(buf.Bytes())
}
