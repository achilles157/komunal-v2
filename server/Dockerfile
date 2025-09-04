# --- Tahap 1: Build ---
# Menggunakan image Go yang sesuai dengan go.mod Anda
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Menyalin file dependensi
COPY go.mod go.sum ./

# Mengunduh semua dependensi
RUN go mod download

# Menyalin seluruh kode sumber proyek Anda
COPY . .

# Build aplikasi Go Anda. Outputnya adalah file binary bernama 'main'
RUN CGO_ENABLED=0 GOOS=linux go build -o /main ./cmd/api

# --- Tahap 2: Run ---
# Menggunakan image minimalis untuk menjalankan aplikasi
FROM alpine:latest

# Menyalin file binary 'main' yang sudah di-build dari tahap sebelumnya
COPY --from=builder /main /main

# Memberi tahu container port mana yang akan diekspos
EXPOSE 8080

# Perintah untuk menjalankan aplikasi saat container dimulai
CMD ["/main"]
