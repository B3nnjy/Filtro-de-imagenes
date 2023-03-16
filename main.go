package main

import (
	"image/color"
	"image/jpeg"
	"log"
	"os"
)

func main() {
	// Abre la imagen
	file, err := os.Open("./imgs/Explorer.jpeg")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Decodifica la imagen
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	// Crea una matriz para almacenar los valores RGB
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	rgb := make([][]color.RGBA, height)
	for y := 0; y < height; y++ {
		rgb[y] = make([]color.RGBA, width)
		for x := 0; x < width; x++ {
			// Obtiene el valor de color RGB en cada punto de la imagen
			rgba := color.RGBAModel.Convert(img.At(x, y)).(color.RGBA)
			rgb[y][x] = rgba
		}
	}

	// Usa la matriz RGB para realizar operaciones adicionales, como
	// análisis de imagen o procesamiento de datos

	// ...
}
