# --- Tahap 1: Build ---
# Menggunakan image Go resmi sebagai basis untuk build
FROM golang:1.21-alpine AS builder

# Menetapkan direktori kerja di dalam container
WORKDIR /app

# Menyalin file dependensi
COPY go.mod ./
COPY go.sum ./

# Mengunduh semua dependensi
RUN go mod download

# Menyalin seluruh kode sumber proyek Anda
COPY . .

# Build aplikasi Go Anda. Outputnya adalah file binary bernama 'main'
# Perhatikan path ke main.go Anda
RUN CGO_ENABLED=0 GOOS=linux go build -o /main ./cmd/api

# --- Tahap 2: Run ---
# Menggunakan image minimalis untuk menjalankan aplikasi (lebih aman dan kecil)
FROM alpine:latest

# Menyalin file binary 'main' yang sudah di-build dari tahap sebelumnya
COPY --from=builder /main /main

# Memberi tahu container port mana yang akan diekspos
EXPOSE 8080

# Perintah untuk menjalankan aplikasi saat container dimulai
CMD ["/main"]