# Use an official Golang runtime as a parent image
FROM golang:latest

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Build the Go app
RUN go build -o main

# Expose port 8000 to the outside world
EXPOSE 8004

# Command to run the executable
CMD ["./main"]