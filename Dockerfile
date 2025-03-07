# Build stage
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /gatewayx ./cmd/gatewayx

# Runtime stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /gatewayx .
EXPOSE 8080
CMD ["./gatewayx"]