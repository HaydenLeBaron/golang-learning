/*
A demo for creating an HTTP server in Go.

To start the HTTP server run:

```sh
go run http-server.go
```

And then point your web browser URL to 
- localhost:8080
to see the first handlers page--and then to
- localhost:8080/earth
to see handler2's page
*/
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/earth", handler2)
    http.ListenAndServe(":8080", nil) // Start server on port 8080
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World\n")
}

func handler2(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello Earth\n")
}
