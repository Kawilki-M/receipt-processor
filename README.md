# Go Web Server

This is a web server written in Go that implements an API for the Fetch [Receipt Processor Challenge](https://github.com/fetch-rewards/receipt-processor-challenge/blob/main/README.md). The API specification is detailed in the ['api.yaml'](https://github.com/Kawilki-M/receipt-processor/blob/main/api/api.yml) file.

## Requirements

- Go 1.20+ 

## Installation

Clone the repository and navigate to the project directory:

```sh
git clone https://github.com/Kawilki-M/receipt-processor.git
cd receipt-processor
```

Ensure all dependencies are installed:
```sh
go mod tidy
```

## Running the Server
To start the server, run:

```sh
go run cmd/api/main.go
```

## Configuration
By default, the server runs on **localhost:8000**. If you need to change the port, you can do so by setting the `PORT` environment variable before running the server:

```sh
export PORT=8080  # Change 8000 to 8080 or another port
go run cmd/api/main.go
```

## API Documentation
For detailed API specifications, refer to api.yaml.
