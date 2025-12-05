package main

import (
	"encoding/json"
	"log"
	"os"
	"recognition/recognition"
)

func save(face recognition.Face, filename string) error {
	log.Println("Guardando...")
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(face); err != nil {
		return err
	}
	return nil
}

const fileName = "decriptor.json"

func main() {
	filename := os.Args[1]
	response, faces := recognition.LoadFromFile(filename)
	if !response {
		log.Fatalf("No se encontro ningun rostro en: %s", filename)
		return
	}
	// Obtener la cara mas grande
	face := recognition.GetBiggerFace(faces)

	// Leer descriptor guardado en descriptores.txt
	file, err := os.Open(fileName)
	if err != nil {
		log.Println("Error al abrir", fileName, ":", err)
		save(face, fileName)
		return
	}
	defer file.Close()

	// Leer descriptor del archivo
	var descriptor recognition.Face

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&descriptor); err != nil {
		log.Println("Error al leer", fileName, ":", err)
		save(face, fileName)
		return
	}
	// Comparar descriptores
	if recognition.IsSamePerson(face, descriptor) {
		log.Println("Los rostros pertenecen a la misma persona.")
	} else {
		log.Println("Los rostros NO pertenecen a la misma persona.")
	}
	save(face, fileName)
}
