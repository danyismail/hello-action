# Use the official Go 1.18.1 image as the base image
FROM golang:1.18.1-alpine

# Set the working directory in the container
WORKDIR /app

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go mod init
RUN go mod tidy
RUN go build -o myapp .

# Expose a port if your application listens on it
# EXPOSE 8080

# Command to run the executable
CMD ["./myapp"]