# Chapter2: Request / Response Cycle

```go
r, err := http.Get("http://www.example.com/")
 ```

```go
// http.Response
type Response struct {
    Status string
    StatusCode int
    Proto string
    ProtoMajor int
    ProtoMinor int
    Header Header
    Body io.ReadCloser
    ContentLength int64
    TransferEncoding []string
    close bool
    Uncompressed bool
    Trailer Header
    Request *Request
    TLS *tls.ConnectionsState
}
```

```go
// example
package main
import (
    "log"
    "net/http"
    "os"
)
func main() {
    r, err := http.Get("http://www.google.com/")
    if err != nil {
        panic(err)
    }

    code := r.Statuscode
    switch {
    case code == 200:
        bodyLength := 1270
        webPageContent := make([]byte, bodyLength)
        r.Body.Read(webPageContent) // Read the data from the server
        out, err := os.OpenFile("index.html", os.O-CREATE|os.O_WRONLY, 0664) // Open a writable file on your computer. Create if it does not exist
        if err != nil {
            panic(err)
        }

        out.Write(webPageContent)
        out.Close()
    default:
        log.Fatal("Failed to retrieve the webpage. Received status code", r.Status)
    }
}
```