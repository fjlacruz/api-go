FROM golang:1.24-alpine

WORKDIR /app

# Copiar archivos de configuración
COPY go.mod go.sum ./

# Instalar dependencias
RUN go mod download

# Copiar el resto del código
COPY . .

# Compilar la aplicación
RUN go build -o main .

# Exponer el puerto
EXPOSE 8088

# Ejecutar la aplicación
#CMD ["./main"]
CMD ["sh", "-c", "sleep 5 && ./main"]