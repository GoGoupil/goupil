# Goupil HTTP Client

A simple HTTP Client to manage HTTP requesting.

The client take a base URL completed with a route when requesting and return HTTP status code.

## Usage

```go
import "github.com/GoGoupil/http"

func main() {
  c := Client{"http://www.google.com"}
  code := c.Get("/")
}
```
