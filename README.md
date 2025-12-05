# Reconocimiento Facial con Go y Dlib

Sistema simple de reconocimiento facial en Go usando dlib a través de la biblioteca `go-face`.

## Requisitos

### 1. Instalar dependencias del sistema

#### Ubuntu/Debian
```bash
sudo apt-get update
sudo apt-get install -y libdlib-dev libblas-dev libatlas-base-dev liblapack-dev libjpeg-turbo8-devsudo libopencv-dev
```

#### Arch Linux
```bash
sudo pacman -S dlib blas lapack libjpeg-turbo opencv vtk
```
### 2. Dependencias go
#### Go-face
```bash
go get github.com/Kagami/go-face
```
#### GoCV (Para webcam)
```bash
go get gocv.io/x/gocv
```

## Funcionamiento
### 1. main.go
```bash
go run main.go <Nombre de la imagen jpg>
```
El programa comparara la cara con la ultima cargada, te dira si son de la misma persona, y la nueva cara se guardara para la proxima comparacion
### 1. webcam.go
```bash
go run main.go <Nombre de la ruta>
```
El programa cargara todas las imagenes jpg de una ruta, y verificara si la persona en pantalla coincide
con alguna de esas fotos