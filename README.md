# Reconocimiento Facial con Go y Dlib

Sistema simple de reconocimiento facial en Go usando dlib a través de la biblioteca `go-face`.

## Características

- Detección de rostros en imágenes
- Extracción de descriptores faciales de 128 dimensiones
- Comparación de rostros para identificación
- Guardado de descriptores para análisis posterior

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

#### macOS
```bash
brew install dlib
```

### 2. Descargar modelos pre-entrenados

Necesitas descargar los modelos de dlib y colocarlos en el directorio `models/`:

```bash
mkdir -p models
cd models

# Modelo de detección de rostros (68 puntos)
wget https://github.com/davisking/dlib-models/raw/master/shape_predictor_5_face_landmarks.dat.bz2
bunzip2 shape_predictor_5_face_landmarks.dat.bz2

# Modelo de reconocimiento facial
wget https://github.com/davisking/dlib-models/raw/master/dlib_face_recognition_resnet_model_v1.dat.bz2
bunzip2 dlib_face_recognition_resnet_model_v1.dat.bz2

# Modelo alternativo con 68 puntos (opcional)
wget https://github.com/davisking/dlib-models/raw/master/shape_predictor_68_face_landmarks.dat.bz2
bunzip2 shape_predictor_68_face_landmarks.dat.bz2

cd ..
```

**Nota:** El reconocedor espera encontrar estos archivos en el directorio `models/`:
- `shape_predictor_5_face_landmarks.dat`
- `dlib_face_recognition_resnet_model_v1.dat`

### 3. Instalar dependencias de Go

```bash
go mod download
```

## Uso

### Detectar rostros en una imagen

```bash
go run main.go imagen.jpg
```

### Ejemplo de salida

```
Analizando imagen: foto.jpg

✓ Se detectaron 2 rostro(s):

Rostro #1:
  Ubicación: X=150, Y=100, Ancho=200, Alto=250
  Descriptor facial (128D): [0.123, -0.456, 0.789, ...] (128 dimensiones)

Rostro #2:
  Ubicación: X=450, Y=120, Ancho=180, Alto=220
  Descriptor facial (128D): [-0.234, 0.567, -0.123, ...] (128 dimensiones)

✓ Descriptores guardados en descriptores.txt
```

## Estructura del código

- **`main()`**: Procesa la imagen y detecta rostros
- **`saveDescriptors()`**: Guarda los descriptores faciales en un archivo
- **`CompareFaces()`**: Calcula la distancia euclidiana entre dos descriptores
- **`IsSamePerson()`**: Determina si dos rostros pertenecen a la misma persona

## Funcionamiento

1. **Detección**: El algoritmo detecta rostros en la imagen usando HOG (Histogram of Oriented Gradients)
2. **Extracción**: Para cada rostro detectado, extrae un descriptor de 128 dimensiones usando una red neuronal ResNet
3. **Comparación**: Los descriptores pueden compararse usando distancia euclidiana (menor distancia = mayor similitud)
4. **Umbral**: Se considera la misma persona si la distancia es menor a 0.4

## Ejemplo de comparación entre imágenes

Para comparar rostros entre dos imágenes, puedes modificar el código o crear un script adicional que:

1. Cargue los descriptores de la primera imagen
2. Cargue los descriptores de la segunda imagen
3. Use la función `IsSamePerson()` para compararlos

## Notas

- El programa guarda automáticamente los descriptores en `descriptores.txt`
- Funciona con formatos: JPG, PNG, BMP
- La precisión mejora con imágenes de buena calidad y rostros bien iluminados
- El umbral de tolerancia (0.4) puede ajustarse según necesidades

## Troubleshooting

Si encuentras errores de compilación:

1. Verifica que dlib esté instalado: `pkg-config --libs dlib`
2. Asegúrate de que los modelos estén descargados en `models/`
3. Verifica la versión de Go: `go version` (debe ser >= 1.21)

## Licencia

Este proyecto usa go-face que envuelve dlib. Dlib está bajo licencia Boost Software License.
