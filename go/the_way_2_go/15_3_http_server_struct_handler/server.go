// very simple HTTP server in golang
package main

import (
    "fmt"
    "log"
    "net"
    "net/http"
    "strings"
)

// mock struct for storing handler fields (css,html maybe)
type Hello struct {
}

// used only for this handler, signature of such methods must be congruent with
// `http.HandleFunc' method's secong argument.
func (h *Hello)Handler(w http.ResponseWriter, req *http.Request) {
    fmt.Println("Inside HelloServer handler")
    fmt.Fprint(w, "hello " + req.URL.Path[strings.LastIndex(req.URL.Path, "/") + 1:])
}

const (
    HOST = "localhost"
    PORT = "9999"
)

func main() {
    var hello *Hello = new(Hello)
    http.HandleFunc("/hello/", hello.Handler)         // register handler function
    if err := http.ListenAndServe(net.JoinHostPort(HOST, PORT), nil); err != nil {
        log.Fatal("ListenAndServe", err.Error())
    }
}
