# Start from golang base image
FROM golang:alpine

# Set the current working directory inside the container 
WORKDIR /app

# Copy the source from the current directory to the working Directory inside the container 
COPY . .

# Build the Go app
CMD ["go","run","main.go"]
