FROM golang:latest

# Update the package list and install Solidity
RUN apt-get update && \
    apt-get install -y software-properties-common && \
    add-apt-repository -y ppa:eth/eth && \
    apt-get update && \
    apt-get install -y solc

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Build the Go app
RUN go build -o main ./cmd/main.go

# Expose port 8080
EXPOSE 8080

# Define the command to run the app
CMD ["./main"]
