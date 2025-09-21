# ---------- Build Stage ----------
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy go.mod & go.sum lebih dulu agar dependency caching optimal
COPY go.mod go.sum ./
RUN go mod download

# Copy seluruh source code
COPY . .

# Build binary dengan nama app
RUN go build -o main .

# ---------- Runtime Stage ----------
FROM alpine:3.20

WORKDIR /app

# Instal tzdata untuk timezone
RUN apk add --no-cache tzdata \
    && cp /usr/share/zoneinfo/Asia/Jakarta /etc/localtime \
    && echo "Asia/Jakarta" > /etc/timezone \
    && apk del tzdata

# copy binary dari builder
COPY --from=builder /app/main .
COPY .env.prod .env

# expose port aplikasi
EXPOSE 3000

# Jalankan binary
CMD ["./main"]
