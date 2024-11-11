# Etapa de build
FROM golang:1.19 AS builder
WORKDIR /app

# Copiar arquivos do projeto
COPY . .

# Instalar dependências e buildar o binário
RUN go mod download
RUN go build -o main .

# Etapa final
FROM alpine:3.16
WORKDIR /app

# Copiar o binário do build
COPY --from=builder /app/main .

# Expor a porta (ajuste para a porta que o seu serviço usa)
EXPOSE 8080

# Executar o binário
CMD ["./main"]
