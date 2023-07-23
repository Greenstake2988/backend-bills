# Establece la imagen de la aplicacion
FROM golang:1.16-alpine

# El directorio de trabajo para la aplicacion
WORKDIR /app

# copiar el archivo al contenedor
COPY main.go .

# creamos la aplicacion
RUN go build -o app

# Exponemos el puerto
EXPOSE 8080

# Establece el punto de ejecucion
CMD ["./app"]