package main

import (
	"log"
	"os"
	"recognition/recognition"
)

func main() {
	response, faces := recognition.LoadFromFile(os.Args[1])
	if !response {
		log.Fatalf("No se encontro ningun rostro en: %s", os.Args[1])
		return
	}

	// Obtener la cara mas grande
	var face recognition.Face
	max := 0
	for _, f := range faces {
		area := f.Rectangle.Dx() * f.Rectangle.Dy()
		if area > max {
			max = area
			face = recognition.Face(f)
		}
	}

}
