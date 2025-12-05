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
	tolerance = 0.3
)

type Face face.Face

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
	return true, faces
}

// CompareFaces compara dos descriptores faciales y retorna la distancia
func CompareFaces(desc1, desc2 Face) float32 {
	var sum float32
	for i := range desc1.Descriptor {
		diff := desc1.Descriptor[i] - desc2.Descriptor[i]
		sum += diff * diff
	}
	return sum
}

// IsSamePerson determina si dos rostros pertenecen a la misma persona
func IsSamePerson(desc1, desc2 Face) bool {
	distance := CompareFaces(desc1, desc2)
	return distance < tolerance
}
