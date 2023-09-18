# Stage 1: Build aplikasi Golang
FROM golang:1.20-alpine3.17 AS builder

RUN apk add --update --no-cache ca-certificates git

WORKDIR /app

# Salin file aplikasi Golang dan file lainnya yang diperlukan
COPY . .

# Instal dependensi yang dibutuhkan
RUN go mod download

# Build aplikasi Golang
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# Stage 2: Copy binary yang sudah dibuat ke image yang baru
FROM bash:5

RUN apk update && apk add \
  grep \
  sed \
  coreutils \
  docker \
  wget


WORKDIR /app

# Salin binary yang sudah dibuat dari stage 1
COPY --from=builder /app/app .

# Jalankan aplikasi saat container berjalan
ENTRYPOINT ["./app"]
