# Use official Golang base image with version 1.24.2
FROM golang:1.24.2

# Set working directory inside the container
WORKDIR /app

# Copy Go modules files first (enables Docker layer caching)
COPY go.mod ./
COPY go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire application source code
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside
EXPOSE 8080

# Run the application
CMD ["./main"]
