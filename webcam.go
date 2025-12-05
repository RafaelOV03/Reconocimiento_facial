package main

import (
	"log"
	"os"
	"path/filepath"
	"recognition/recognition"

	"gocv.io/x/gocv"
)

type Image struct {
	Path string
	Face recognition.Face
}

func main() {
	ruta := os.Args[1]
	// Cargar todas las imagenes jpg de la ruta
	files, err := os.ReadDir(ruta)
	if err != nil {
		log.Fatal(err)
	}

	var images []Image
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".jpg" {
			path := filepath.Join(ruta, file.Name())
			success, faces := recognition.LoadFromFile(path)
			if success {
				face := recognition.GetBiggerFace(faces)
				images = append(images, Image{
					Path: path,
					Face: face,
				})
			} else {
				log.Printf("No se encontro ningun rostro en: %s", path)
			}
		}
	}
	log.Println("Iniciando webcam...")
	// Iniciar webcam
	webcam, err := gocv.OpenVideoCapture(0)
	if err != nil {
		log.Fatal(err)
	}
	defer webcam.Close()

	window := gocv.NewWindow("Reconocimiento Facial")
	defer window.Close()

	img := gocv.NewMat()
	defer img.Close()

	frameCount := 0
	captureInterval := 30 // Capturar cada 30 frames (~1 segundo a 30fps)

	log.Println("Webcam activa. Presiona ESC para salir")

	for {
		if ok := webcam.Read(&img); !ok {
			log.Println("No se pudo leer de la webcam")
			break
		}
		if img.Empty() {
			continue
		}

		frameCount++

		// Capturar y procesar cada N frames
		if frameCount%captureInterval == 0 {
			// Guardar frame temporal
			tempFile := "temp_webcam.jpg"
			gocv.IMWrite(tempFile, img)

			// Detectar rostro en la captura
			success, faces := recognition.LoadFromFile(tempFile)
			if success && len(faces) > 0 {
				capturedFace := recognition.GetBiggerFace(faces)

				// Comparar con todas las imágenes cargadas usando goroutines
				go compareWithAll(capturedFace, images)
			} else {
				log.Println("No se detectó ningún rostro en la captura")
			}
		}

		// Mostrar video en ventana
		window.IMShow(img)

		// Salir con ESC
		if window.WaitKey(1) == 27 {
			break
		}
	}
}

// Función que compara un rostro con todas las imágenes de forma concurrente
func compareWithAll(capturedFace recognition.Face, images []Image) {
	type Result struct {
		Path    string
		IsMatch bool
	}

	results := make(chan Result, len(images))

	// Lanzar una goroutine por cada comparación
	for _, img := range images {
		go func(img Image) {
			isMatch := recognition.IsSamePerson(capturedFace, img.Face)
			results <- Result{
				Path:    img.Path,
				IsMatch: isMatch,
			}
		}(img)
	}

	// Recolectar resultados
	matches := []string{}
	for i := 0; i < len(images); i++ {
		result := <-results
		if result.IsMatch {
			matches = append(matches, result.Path)
		}
	}

	// Mostrar resultados
	if len(matches) > 0 {
		log.Printf("✓ Coincidencia encontrada con %d imagen(es):", len(matches))
		for _, match := range matches {
			log.Printf("  - %s", match)
		}
	} else {
		log.Println("✗ No se encontraron coincidencias")
	}
}
