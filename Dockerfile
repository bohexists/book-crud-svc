# Use the official Go image as the base image
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /app


RUN ping -c 4 google.com


# Copy the go.mod and go.sum files to install dependencies
COPY go.mod go.sum ./

# Download and cache the dependencies
RUN go mod download

# Copy the remaining application files to the working directory
COPY . .

# Build the Go application
RUN go build -o main .

# Set environment variables
ENV PORT=8080

# Command to run the application
CMD ["./main"]