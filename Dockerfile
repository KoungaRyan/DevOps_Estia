# Étape 1 : Build
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copier le code source
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Compiler le binaire
RUN go build -o app .

# Étape 2 : Image finale
FROM alpine:3.18

WORKDIR /app

# Copier uniquement le binaire depuis le builder
COPY --from=builder /app/app .

# Exposer un port (OPTIONNEL)
EXPOSE 8080

# Lancer le binaire (PAS le fichier main.go)
ENTRYPOINT ["./app"]
