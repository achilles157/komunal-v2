# --- Tahap 1: Build ---
FROM golang:1.21-alpine AS builder
WORKDIR /app

# Salin dan unduh dependensi terlebih dahulu
COPY go.mod go.sum ./
RUN go mod download

# Salin sisa kode aplikasi
COPY . .

# Build aplikasi Go Anda menjadi satu file binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /main ./cmd/api

# --- Tahap 2: Run ---
FROM alpine:latest

# Salin file binary yang sudah di-build dari tahap sebelumnya
COPY --from=builder /main /main

# Expose port yang digunakan oleh aplikasi Go Anda
EXPOSE 8080

# Perintah untuk menjalankan aplikasi
CMD ["/main"]
