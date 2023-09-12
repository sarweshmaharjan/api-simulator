# Use an official Go runtime as a parent image
FROM golang:1.20

# Set the working directory inside the container
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .

# Copy the .env file from the ./cmd/ directory to /app/
COPY ./cmd/.env ./.env

# Build the Go application
RUN go mod download
RUN go build -o api_response_simulation ./cmd/main.go

# Expose the port your application will run on
EXPOSE 9999

# Run your application when the container starts, listening on 0.0.0.0
CMD ["./api_response_simulation"]
