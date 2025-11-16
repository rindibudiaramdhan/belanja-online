# ---------- BUILDER ----------
FROM golang:1.25.3-alpine AS builder
WORKDIR /app

# Install git (required for go mod download)
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build binary
RUN go build -o server ./main.go

# ---------- FINAL IMAGE ----------
FROM alpine:3.19
WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 8080
CMD ["./server"]
