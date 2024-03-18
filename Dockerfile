# Constuir imagen
FROM golang:1.22.0-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main cmd/main.go

# Ejecutar imagen
FROM alpine
WORKDIR /app

# Crear archivo .env
RUN touch ./.env
# Copiar archivo .env y cambiar permisos
RUN chmod +r .env && cat .env

# Copiar el binario y definir el comando de ejecución
COPY --from=builder /app/main .
EXPOSE 3700
CMD ["sh", "-c", "source .env && ./main"]