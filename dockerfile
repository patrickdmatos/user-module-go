# Usando uma imagem Go para compilar o código
FROM golang:1.20 AS builder
WORKDIR /app

# Copiar arquivos do projeto
COPY . .

# Baixar dependências e build
RUN go mod download
RUN go build -o main .

# Preparando a imagem final
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .

# Porta em que o serviço será exposto
EXPOSE 3030

# Executa o binário gerado
CMD ["./main"]
