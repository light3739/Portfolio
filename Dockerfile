# Step 1: Start with a base image for building the Go application.
FROM golang:1.21-alpine as builder

# Set the working directory inside the container.
WORKDIR /app

# Copy go.mod and go.sum files to the workspace.
COPY go.mod .
COPY go.sum .

# Download all dependencies.
# They will be cached if the go.mod and go.sum files are not changed.
RUN go mod download

# Install additional packages required for building with CGO.
RUN apk --no-cache add gcc libc-dev

# Copy the source code from the current directory to the working directory inside the container.
COPY . .

# Build the Go app for a smaller and secured application.
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage from scratch for a smaller image size.
FROM alpine:3.14

# Set the working directory.
WORKDIR /app

# Copy the binary file and necessary files from the builder stage.
COPY --from=builder /app/main .
COPY --from=builder /app/templates /app/templates
COPY --from=builder /app/static /app/static

# Include the .env file in the Docker image
COPY .env .

# Expose port 8080 to the outside world.
EXPOSE 8080

# Command to run the executable.
CMD ["./main"]
