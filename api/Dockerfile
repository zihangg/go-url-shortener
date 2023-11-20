# Start from the official golang base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the local code to the container's working directory
COPY . .

# Download dependencies
RUN go mod download

# Build the Go application
RUN go build -o app

# Command to run the application
CMD ["./app"]
