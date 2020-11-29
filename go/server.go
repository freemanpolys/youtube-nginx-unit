package main
import (
    "fmt"
    "log"
	"net/http"
	"unit.nginx.org/go"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Hello!")
    })


    fmt.Printf("Starting server at port 8080\n")
    if err := unit.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}