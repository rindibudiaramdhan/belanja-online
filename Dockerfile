# =========================
# 1. Builder Container
# =========================
FROM golang:1.23-alpine AS builder

ENV GOTOOLCHAIN=auto

WORKDIR /app

# Install dependencies needed for Go build
RUN apk add --no-cache git

# Copy go modules first (better use build cache)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server main.go


# =========================
# 2. Runtime Container
# =========================
FROM alpine:3.19

WORKDIR /app

# Install certificates (for HTTPS calls)
RUN apk add --no-cache ca-certificates

# Copy binary from builder stage
COPY --from=builder /app/server .

# Expose API port
EXPOSE 8080

CMD ["./server"]
