# Use the official Go image as the base image
FROM golang:1.22-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to install dependencies
COPY go.mod go.sum ./

# Download and cache the dependencies
RUN go mod download

# Copy the remaining application files to the working directory
COPY . .

# Copy the .env file to the working directory
COPY .env .env

# Install Dockerize
RUN apk add --no-cache wget \
    && wget https://github.com/jwilder/dockerize/releases/download/v0.8.0/dockerize-linux-arm64-v0.8.0.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-linux-arm64-v0.8.0.tar.gz \
    && rm dockerize-linux-arm64-v0.8.0.tar.gz


# Wait for the database to be ready and run tests
CMD dockerize -wait tcp://db:5432 -timeout 30s \
    && go test -v ./...

# Build the Go application
RUN go build -o main .

# Set environment variables
ENV PORT=8080

# Command to run the application
#CMD ["sh", "-c", "dockerize -wait tcp://db:5432 -timeout 60s && go test -v ./... && ./main"]
CMD ["sh", "-c", "dockerize -wait tcp://db:5432 -timeout 60s && ./main"]