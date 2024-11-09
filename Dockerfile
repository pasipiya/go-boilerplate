# Use the official Go image
FROM golang:1.20-alpine

# Set up work directory
WORKDIR /app

# Copy and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the app
RUN go build -o main ./cmd/app

# Expose port and run
EXPOSE 8080 6060
CMD ["./main"]