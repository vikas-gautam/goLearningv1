package main

import (
    "fmt"
    "net/http"
    "log"
    "html"
)

func helloWorld(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Hello World")
}

func vikas(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "vikas is here")
    fmt.Println("vikas is serving here")
    log.Println("initalising go server")

}
//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request))


func main() {
    http.HandleFunc("/", helloWorld)
    http.HandleFunc("/vikas", vikas)
    http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    })
    log.Fatal(http.ListenAndServe(":9090", nil))
}
//https://betakuang.medium.com/setup-go-development-environment-with-vs-code-and-wsl-on-windows-62bd4625c6a7

// smart scale

