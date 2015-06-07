# Goupil HTTP Client

A simple HTTP Client to manage HTTP requesting.

The client take a base URL completed with a route when requesting and return HTTP status code.

## Usage

```go
import (
    "fmt"
    "github.com/GoGoupil/http"
)

func main() {
    c := http.Client{}
    c.Open("devatoria.info", 80)
    defer c.Close()
    elapsed, code := c.Get("/")
    fmt.Printf("Elapsed: %f\n", elapsed)
    fmt.Printf("HTTP code: %d\n", code)
}
```
