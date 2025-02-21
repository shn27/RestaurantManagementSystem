FROM golang:alpine3.20

WORKDIR /app

# Copy dependencies and cache Go module downloads
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o app .

# Set the entry point for the container
CMD ["./app"]
