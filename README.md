# gin-error

A Gin middleware sending error message to response

## Installation

Download Go module:

```sh
go get -u github.com/r04922101/gin-error
```

## How to Use

### Import it in your code

```go
import ginerror "github.com/r04922101/gin-error"
```

### Usage

```go
r := gin.Default()
r.Use(ginerror.RespondError)
```

### Example

Check the example code in [example/main.go](./example/main.go)

Run `example/main.go`

```sh
go run main.go
```

Send an HTTP POST request to `localhost:8080/error`

```sh
curl -X POST 'localhost:8080/error'
```

### Result

Check your console which sent the POST request, and check the response.

```text
Error #01: error occurs
```
