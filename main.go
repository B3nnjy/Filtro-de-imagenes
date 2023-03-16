package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	imagenes := Archivo()

	for i := 0; i < len(imagenes); i++ {
		wg.Add(1)
		go CambiarImagen(imagenes[i], &wg)
	}

	wg.Wait()
	fmt.Println("Fin!!")
}

func Archivo() []string {
	var nombres []string
	files, err := ioutil.ReadDir("./imgs")

	if err != nil {
		log.Fatal("Error al abrir el directorio!! ", err)
	}

	nombres = make([]string, len(files))
	var x int = 0

	for _, arch := range files {
		if !arch.IsDir() {

			nombres[x] = arch.Name()
			x++
		}
	}

	return nombres
}

func CambiarImagen(dir string, wg *sync.WaitGroup) {
	defer wg.Done()
	// Abre la imagen
	file, err := os.Open("./imgs/" + dir)
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

	CambiarARojo(dir, width, height, rgb)

}

func CambiarARojo(dir string, width, height int, rgb [][]color.RGBA) {
	imagen := image.Rect(0, 0, width, height)
	newImg := image.NewRGBA(imagen)

	for i := imagen.Min.X; i < imagen.Max.X; i++ {
		for j := imagen.Min.Y; j < imagen.Max.Y; j++ {
			rgb[j][i].R = 255
			newImg.Set(i, j, rgb[j][i])
		}
	}

	re := regexp.MustCompile(`\.`)
	subCad := re.Split(dir, -1)

	newFile, err := os.Create("./imgs/rojo/" + subCad[0] + "_Roja" + ".jpg")
	if err != nil {
		panic(err)
	}
	defer newFile.Close()

	err = jpeg.Encode(newFile, newImg, nil)
	if err != nil {
		panic(err)
	}
}
