# Especifica la imagen base
FROM golang:1.22.0

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /go/src

# Copia el código fuente de tu aplicación al directorio de trabajo en el contenedor
COPY . .

# Compila la aplicación dentro del contenedor
RUN go mod tidy
RUN go build -v ./...

# Expone el puerto en el que la aplicación escucha
EXPOSE 8080

# Comando para ejecutar la aplicación cuando el contenedor se inicie
CMD ["go run src/main.go"]
