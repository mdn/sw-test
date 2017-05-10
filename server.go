package main
// Simple file server
// Install Google Go 1.x and then run `go run server.go` in this directory.
import (
        "log"
        "fmt"
        "net/http"
)
const PORT = ":8080"
func main() {
        fmt.Println("Visit http://localhost" + PORT )
        log.Fatal(http.ListenAndServe(PORT, http.FileServer(http.Dir("./"))))
}
