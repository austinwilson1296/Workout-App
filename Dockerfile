# Build stage
FROM golang:1.23.2-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache gcc musl-dev

# Install goose
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

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



# Copy binary and project files from builder
COPY --from=builder /app/fitted .
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/static ./static
COPY --from=builder /app/sql ./sql
COPY --from=builder /app/scripts ./scripts

# Ensure db_migrations script is executable
RUN chmod +x ./scripts/db_migration.sh

EXPOSE 8080

# Run migrations and start the application
CMD ["sh", "-c", "./scripts/db_migration && ./fitted"]