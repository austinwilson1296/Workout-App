# Build stage
FROM golang:1.23.2-alpine AS builder

WORKDIR /app

# Install goose
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

# Copy the entire project into the builder
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o fitted .

# Final stage
FROM alpine:3.19

WORKDIR /app

# Install necessary runtime dependencies
RUN apk add --no-cache tzdata bash

# Copy the Go binary, goose, and other necessary files from the builder stage
COPY --from=builder /go/bin/goose /usr/local/bin/goose
COPY --from=builder /app/fitted /app/fitted
COPY --from=builder /app/scripts /app/scripts
COPY --from=builder /app/templates /app/templates
COPY --from=builder /app/static /app/static
COPY --from=builder /app/sql /app/sql

# Ensure the db_migration.sh script is executable
RUN chmod +x ./scripts/db_migration.sh

# List all files in /app (for debugging purposes)
RUN echo "Listing all files in /app:" && ls -R /app > /app/files_list.txt

# Run the db_migration.sh script at container startup and then run the app
CMD ["sh", "-c", "./scripts/db_migration.sh && ./fitted"]
