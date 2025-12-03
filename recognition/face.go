package recognition

import (
	"fmt"
	"log"

	"github.com/Kagami/go-face"
)

const (
	// Directorios de modelos pre-entrenados de dlib
	modelDir = "./models"
	// Umbral de similitud para considerar un rostro como coincidente
	tolerance = 0.4
)

func LoadFromFile(imagePath string) (bool, []face.Face) {

	// Inicializar el reconocedor facial con los modelos de dlib
	rec, err := face.NewRecognizer(modelDir)
	if err != nil {
		log.Fatalf("Error al inicializar el reconocedor: %v", err)
	}
	defer rec.Close()

	// Detectar rostros en la imagen
	faces, err := rec.RecognizeFile(imagePath)
	if err != nil {
		log.Fatalf("Error al procesar la imagen: %v", err)
	}

	// Mostrar resultados
	if len(faces) == 0 {
		fmt.Println("No se detectaron rostros en la imagen.")
		return false, nil
	}

	// Mostrar resultados
	/*for i, f := range faces {
		fmt.Printf("Rostro #%d:\n", i+1)
		fmt.Printf("  Ubicación: X=%d, Y=%d, Ancho=%d, Alto=%d\n",
			f.Rectangle.Min.X, f.Rectangle.Min.Y,
			f.Rectangle.Dx(), f.Rectangle.Dy())
		fmt.Printf("  Descriptor facial (128D): [%.3f, %.3f, %.3f, ...] (%d dimensiones)\n",
			f.Descriptor[0], f.Descriptor[1], f.Descriptor[2], len(f.Descriptor))
		fmt.Println()
	}*/
	// Obtener la cara mas grande
	/*var face face.Face
	max := 0
	for _, f := range faces {
		area := f.Rectangle.Dx() * f.Rectangle.Dy()
		if area > max {
			max = area
			face = f
		}
	}*/
	// Guardar descriptores para comparación futura
	return true, faces
}

// CompareFaces compara dos descriptores faciales y retorna la distancia
func CompareFaces(desc1, desc2 face.Descriptor) float32 {
	var sum float32
	for i := range desc1 {
		diff := desc1[i] - desc2[i]
		sum += diff * diff
	}
	return sum
}

// IsSamePerson determina si dos rostros pertenecen a la misma persona
func IsSamePerson(desc1, desc2 face.Descriptor) bool {
	distance := CompareFaces(desc1, desc2)
	return distance < tolerance
}
