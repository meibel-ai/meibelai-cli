# Build stage
FROM golang:1.21-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git ca-certificates

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o meibel .

# Final stage
FROM alpine:latest

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates

# Create non-root user
RUN addgroup -g 1000 -S meibel && \
    adduser -u 1000 -S meibel -G meibel

# Copy binary from builder
COPY --from=builder /app/meibel /usr/local/bin/meibel

# Set ownership
RUN chown -R meibel:meibel /usr/local/bin/meibel && \
    chmod +x /usr/local/bin/meibel

# Create config directory
RUN mkdir -p /home/meibel/.config/meibel && \
    chown -R meibel:meibel /home/meibel

# Switch to non-root user
USER meibel
WORKDIR /home/meibel

# Set default command
ENTRYPOINT ["meibel"]
CMD ["--help"]

# Labels
LABEL org.opencontainers.image.title="Meibel CLI"
LABEL org.opencontainers.image.description="CLI for interacting with Meibel AI API"
LABEL org.opencontainers.image.vendor="Meibel AI"
LABEL org.opencontainers.image.licenses="MIT"
LABEL org.opencontainers.image.source="https://github.com/meibel-ai/meibel-cli"