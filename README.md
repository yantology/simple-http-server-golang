# Simple HTTP Server in Golang

This project is a simple HTTP server written in Go. It serves static files and provides basic form handling and API endpoints.

## File Structure

- simple-http-server-golang/
  - static/
    - form.html
  - main.go
  - main_test.go
  - go.mod
  - go.sum
  - README.md

- `static/`: Directory containing static files to be served.
- `main.go`: Main application file containing the server code.
- `main_test.go`: Test file for the application.
- `go.mod`: Go module file.
- `go.sum`: Go dependencies checksum file.
- `README.md`: Project documentation.

## Libraries/Packages Used

- `net/http`: Standard library for HTTP client and server implementations.
- `github.com/stretchr/testify`: Library for writing unit tests.

## How to Use the Project

### Development

1. **Clone the repository:**

    ```sh
    git clone https://github.com/yantology/simple-http-server-golang.git
    cd simple-http-server-golang
    ```

2. **Install dependencies:**

    ```sh
    go mod tidy
    ```

3. **Run the server:**

    ```sh
    go run main.go
    ```

4. **Run tests:**

    ```sh
    go test ./...
    ```

### Production

1. **Build the application:**

    ```sh
    go build -o simple-http-server.exe
    ```

2. **Run the built application:**

    ```sh
    ./simple-http-server
    ```

The server will start on port 8080. You can access it at `http://localhost:8080`.

- To access the form, navigate to `http://localhost:8080/form`.
- To access the hello endpoint, navigate to `http://localhost:8080/api/v1/hello`.
- To submit the form, use the endpoint `http://localhost:8080/api/v1/form`.

### Docker

You can also run the application using Docker(look dockerfile for the setup).

1. **Build the Docker image:**

    ```sh
    docker build -t simple-http-server .
    ```

2. **Run the Docker container:**

    ```sh
    docker run -p 8080:8080 simple-http-server
    ```

The server will start on port 8080. You can access it at `http://localhost:8080`.

- To access the form, navigate to `http://localhost:8080/form`.
- To access the hello endpoint, navigate to `http://localhost:8080/api/v1/hello`.
- To submit the form, use the endpoint `http://localhost:8080/api/v1/form`.

## License

This project is licensed under the MIT License.
