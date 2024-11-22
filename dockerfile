# Etapa de build
FROM golang:1.19 AS builder
WORKDIR /app

# Copiar o código para o contêiner
COPY . .

# Baixar dependências
RUN go mod download

# Compilar o binário
RUN go build -o main .

# Etapa final
FROM alpine:3.16

# Copiar o binário compilado da etapa de build
WORKDIR /app
COPY --from=builder /app/main .

# Expor a porta 3000
EXPOSE 3000

# Iniciar o binário
CMD ["./main"]
