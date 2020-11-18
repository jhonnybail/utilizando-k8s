package main

import "net/http"
import "fmt"

func greetings(value string) string {
    return "<b>" + value + "</b>"
}

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, greetings("Code.education Rocks!"))
    })

    http.ListenAndServe(":8000", nil)
}
