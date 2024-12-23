# Gomatcher

Gomatcher is a Go-based project designed for processing data streams, identifying patterns, and serving as a modular endpoint manager. It includes functionalities for managing and composing data, running servers, and handling HTTP-based endpoints.

## Features

- **Data Management**: Efficiently manage and compose data across multiple endpoints.
- **Pattern Matching**: Detect and validate patterns in data using regular expressions.
- **HTTP Endpoint Support**: Easily set up and run HTTP endpoints to process incoming requests.
- **Concurrency**: Leverages Go's concurrency model for scalable and efficient operations.

## Installation

### Prerequisites

- [Go](https://golang.org/dl/) 1.19 or later

### Steps

1. Clone the repository:

   ```bash
   git clone https://github.com/amfonelic/gomatcher.git
   cd gomatcher
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Build the project:

   ```bash
   go build ./...
   ```

## Usage

### Environment Variables

Set the environment variables in .env.example before running the application


### Running the Application

Run the application with:

```bash
go run ./cmd
```

### Running Tests

Run all tests to verify the implementation:

```bash
go test ./...
```

