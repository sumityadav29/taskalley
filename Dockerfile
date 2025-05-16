FROM golang:1.22.4-alpine

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o main ./cmd/server

# Expose port
EXPOSE 8888

# Command to run the executable
CMD ["./main"]
