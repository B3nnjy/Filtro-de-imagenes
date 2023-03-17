package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
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
		go CambiarImagen(imagenes[i], &wg, i)
	}

	wg.Wait()
	fmt.Println("Fin!!")
}

func Archivo() []string {
	var x int = 0
	var nombres []string
	files, err := os.ReadDir("./imgs")

	if err != nil {
		log.Fatal("Error al abrir el directorio!! ", err)
	}

	for _, arch := range files {
		if !arch.IsDir() {
			x++
		}
	}

	nombres = make([]string, x)
	x = 0

	for _, arch := range files {
		if !arch.IsDir() {
			nombres[x] = arch.Name()
			x++
		}
	}
	return nombres
}

func CambiarImagen(dir string, wg *sync.WaitGroup, id int) {
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

	CambiarA(dir, "rojo", width, height, rgb)
	CambiarA(dir, "azul", width, height, rgb)
	CambiarA(dir, "verde", width, height, rgb)

	fmt.Println("Hilo ", id, " finalizado")
}

func CambiarA(dir, color string, width, height int, rgb [][]color.RGBA) {
	imagen := image.Rect(0, 0, width, height)
	newImg := image.NewRGBA(imagen)

	var R, G, B uint8

	for i := imagen.Min.X; i < imagen.Max.X; i++ {
		for j := imagen.Min.Y; j < imagen.Max.Y; j++ {
			R = rgb[j][i].R
			G = rgb[j][i].G
			B = rgb[j][i].B

			switch color {
			case "rojo":
				rgb[j][i].R = 255
			case "azul":
				rgb[j][i].B = 255
			case "verde":
				rgb[j][i].G = 255
			}

			newImg.Set(i, j, rgb[j][i])

			rgb[j][i].R = R
			rgb[j][i].B = B
			rgb[j][i].G = G
		}
	}

	re := regexp.MustCompile(`\.`)
	subCad := re.Split(dir, -1)

	newFile := guardarImg(subCad[0], color)
	defer newFile.Close()

	err := jpeg.Encode(newFile, newImg, nil)
	if err != nil {
		panic(err)
	}
}

func guardarImg(img, color string) *os.File {
	_, err := os.Stat("./imgs/" + color + "/")

	if os.IsNotExist(err) {
		err := os.Mkdir("./imgs/"+color, 0755)

		if err != nil {
			fmt.Println("Error al crear el directorio ./imgs/" + color)
		} else {
			fmt.Println("Se ha creado el directorio ./imgs/" + color)
		}
	}

	newFile, err := os.Create("./imgs/" + color + "/" + img + "_" + color + ".jpg")
	if err != nil {
		panic(err)
	}

	return newFile
}
