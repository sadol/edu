package main
import (
    "fmt"
    "os"
    "net/http"
    "log"
    "io/ioutil"
    "flag"
)

func main() {
    flag.Parse()
    if flag.NArg() != 1 {
        fmt.Println("Usage: client <URL> .")
        os.Exit(1)
    }

    url := fmt.Sprintf("%s", flag.Arg(0))
    res, err := http.Get(url)
    checkError(err)
    data, err := ioutil.ReadAll(res.Body)
    checkError(err)
    fmt.Printf("Got: %q.", string(data))
}

func checkError(err error) {
    if err != nil {
        log.Fatalf("Get: %v.", err)
    }
}
