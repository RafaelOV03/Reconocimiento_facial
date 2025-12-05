# Reconocimiento Facial con Go y Dlib

Sistema simple de reconocimiento facial en Go usando dlib a través de la biblioteca `go-face`.

## Requisitos

### 1. Instalar dependencias del sistema

#### Ubuntu/Debian
```bash
sudo apt-get update
sudo apt-get install -y libdlib-dev libblas-dev libatlas-base-dev liblapack-dev libjpeg-turbo8-dev
```

#### Arch Linux
```bash
sudo pacman -S dlib blas lapack libjpeg-turbo
```

## Funcionamiento
```bash
main.go < Nombre de la imagen jpg >
```
El programa comparara la cara con la ultima cargada, te dira si son de la misma persona, y la nueva cara se guardara para la proxima comparacion