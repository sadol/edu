package main

import (
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "os"
    "time"
)

const FILENAME = "http.times"

func main() {
    start := time.Now()
    ch := make(chan string)

    outputFile, err := os.Create(FILENAME)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer outputFile.Close()

    for _, url := range os.Args[1:] {
        go fetch(url, ch)
    }
    for range os.Args[1:] {
        outputFile.WriteString(<-ch)                 // no error check!
    }
    // no error check!
    outputFile.WriteString(fmt.Sprintf("%.2fs elapsed.\n", time.Since(start).Seconds()))
}

func fetch(url string, ch chan<- string) {
    start := time.Now()
    resp, err := http.Get(url)
    if err != nil {
        ch <- fmt.Sprint(err)
        return
    }

    nbytes, err:= io.Copy(ioutil.Discard, resp.Body)
    resp.Body.Close()
    if err != nil {
        ch <- fmt.Sprintf("While reading %s: %v.\n", url, err)
        return
    }

    ch <- fmt.Sprintf("%.2fs %7d %s.\n", time.Since(start).Seconds(), nbytes, url)
}
