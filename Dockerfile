FROM golang:1.22-alpine as builder

# Install git and python3 and java 21
RUN apk update && apk add --no-cache git python3 openjdk21-jdk

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Run setup and build the app
RUN go run setup.go

# Build the Go app
RUN go build -o app cmd/battery-historian/battery-historian.go

# Compile the thing together
RUN mkdir output && \
    cp -r third_party output/ && \
    cp -r static output/ && \
    cp -r compiled output/ && \
    cp -r templates output/ && \
    cp app output/ && \
    cp -r output/* .

# Start a new stage from scratch
FROM alpine:latest

WORKDIR /app

# Copy the Pre-built binary file from the previous stage, as well as any of the required files to run the app
COPY --from=builder /app/output/ /app

# Command to run the executable
CMD ["./app"]
