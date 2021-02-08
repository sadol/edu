// Prints the content found at a URL;
// need to use protocol as a part of argument, for example `fetch http://...'
package main

import (
    "fmt"
    "net/http"
    "os"
    "io"
)

func main() {
    for _, url := range os.Args[1:] {
        if resp, err := http.Get(url); err != nil {                    // error
            fmt.Fprintf(os.Stderr, "%s: %v.\n", os.Args[0], err)
            os.Exit(1)
        } else {                                                          // OK
            defer resp.Body.Close()
            if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
                fmt.Fprintf(os.Stderr, "%s reader error: %v.\n", os.Args[0], err)
                os.Exit(2)
            }
        }
    }
}
