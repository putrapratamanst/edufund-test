FROM golang:latest

RUN mkdir -p /app

# Set the current working directory inside the container 
WORKDIR /app

# Copy go mod and sum files 
COPY go.mod go.sum ./
COPY . .

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 
# RUN go mod download
# Generate swagger from GOPATH.
RUN go mod vendor

RUN go build -o ./app ./cmd/main.go

EXPOSE 80

ENTRYPOINT ["./app"]