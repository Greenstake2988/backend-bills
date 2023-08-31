# Establece la imagen de la aplicacion
FROM golang:1.20-alpine

# El directorio de trabajo para la aplicacion
WORKDIR /app

#RUN apk add --no-cache gcc
COPY go.mod go.sum config ./
# Instala el paquete de SQLite3
# Actualiza el sistema de paquetes
#RUN apk update
#RUN apk add --no-cache sqlite sqlite-dev
#RUN apk add --no-cache build-base

RUN go mod download
COPY . .
# copiar el archivo al contenedor
RUN go build -o ./out/dist .

ENV PORT=8080

# Expose the same port that the application listens on
EXPOSE $PORT

# Establece el punto de ejecucion
CMD ./out/dist