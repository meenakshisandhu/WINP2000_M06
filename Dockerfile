# Step 1: Use a Go base image
FROM golang:1.20-alpine

# Step 2: Set the working directory inside the container
WORKDIR /app

# Step 3: Copy the Go module files
COPY go.mod go.sum ./

# Step 4: Download the Go modules
RUN go mod tidy

# Step 5: Copy the application source code
COPY . .

# Step 6: Build the Go application
RUN go build -o app .

# Step 7: Set the entrypoint to run the Go application
CMD ["./app"]
