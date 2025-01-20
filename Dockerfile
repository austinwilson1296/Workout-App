# Build stage
FROM golang:1.23.2-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache gcc musl-dev

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o fitted .

# Final stage
FROM alpine:3.19

WORKDIR /app

# Install necessary runtime dependencies
RUN apk add --no-cache tzdata

# Install goose
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

# Copy binary from builder
COPY --from=builder /app/fitted .

# Copy required project files
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/static ./static
COPY --from=builder /app/sql ./sql
COPY --from=builder /app/migrations ./migrations

# Copy and set up start script
COPY db_migrations.sh .
RUN chmod +x db_migrations.sh

EXPOSE 8080

CMD ["./fitted"]