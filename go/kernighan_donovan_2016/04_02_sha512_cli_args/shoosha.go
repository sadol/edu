package main

import (
    "fmt"
    "crypto/sha512"
    "flag"
    "bufio"
    "os"
)

func main() {
    var strToHash []string                             // to raw string to hash
    var noArgs int
    fShaType := flag.String("t", "512",
                       "Choose type of the SHA hash sum function: 384 or 512.")
    flag.Parse()

    if flag.NArg() > 0 {     // number of non option args sent via cli directly
        noArgs = flag.NArg()
        strToHash = flag.Args()                   // non option args themselves
    } else {                        // try to read from stdin (piping and such)
        scanner := bufio.NewScanner(os.Stdin)
        scanner.Split(bufio.ScanWords)
        for scanner.Scan() {
            noArgs++
            strToHash = append(strToHash, scanner.Text())
        }
        if err := scanner.Err(); err != nil {
            fmt.Println(err)
            os.Exit(1)
        }
    }

    switch *fShaType {
    case "384":
        outputValues := make([][sha512.Size384]byte, noArgs)
        for id, val := range strToHash {
            outputValues[id] = sha512.Sum384([]byte(val))
        }
        if _, err := fmt.Fprintf(os.Stdout, "%v", outputValues); err != nil {
            fmt.Println(err)
            os.Exit(1)
        }
    default:                                                           // "512"
        outputValues := make([][sha512.Size]byte, noArgs)
        for id, val := range strToHash {
            outputValues[id] = sha512.Sum512([]byte(val))
        }
        if _, err := fmt.Fprintf(os.Stdout, "%v", outputValues); err != nil {
            fmt.Println(err)
            os.Exit(1)
        }
    }
}
