FROM golang:1.24-bullseye

# Carpeta de trabajo igual a la que monta VS Code
WORKDIR /workspaces/golang

# Instalar Air para hot reload
RUN go install github.com/air-verse/air@latest

# Copiar go.mod y go.sum si existen para aprovechar cache
#COPY go.mod go.sum* ./

# Copiar el resto del código
COPY . .

RUN [ ! -f go.mod ] && go mod init example.com/go-api || true
RUN go mod tidy



# Puerto para API (si aplicara)
EXPOSE 8080

# Ejecutar Air por defecto
CMD ["air", "-c", ".air.toml"]

