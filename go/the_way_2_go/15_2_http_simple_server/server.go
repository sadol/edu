// very simple HTTP server in golang
package main

import (
    "fmt"
    "log"
    "net"
    "net/http"
    "strings"
)

const (
    HOST = "localhost"
    PORT = "9999"
)

func main() {
    http.HandleFunc("/hello/", HelloServer)         // register handler function
    http.HandleFunc("/shouthello/", ShoutServer)         // register handler function
    err := http.ListenAndServe(net.JoinHostPort(HOST, PORT), nil)
    if err != nil {
        log.Fatal("ListenAndServe", err.Error())
    }
}

// handler function (for `hello' directory in this case)
func HelloServer(w http.ResponseWriter, req *http.Request) {
    fmt.Println("Inside HelloServer handler")
    fmt.Fprint(w, "hello " + req.URL.Path[strings.LastIndex(req.URL.Path, "/") + 1:])
}

// handler function (for `shouthello' directory in this case)
func ShoutServer(w http.ResponseWriter, req *http.Request) {
    fmt.Println("Inside ShoutServer handler")
    fmt.Fprint(w, "hello " + strings.ToUpper(req.URL.Path[strings.LastIndex(req.URL.Path, "/") + 1:]))
}
