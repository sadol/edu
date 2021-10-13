package main

import (
    "fmt"
    "crypto/sha512"
    "crypto/sha256"
    "flag"
    "bufio"
    "os"
)

func main() {
    var strToHash []string                             // to raw string to hash
    var noArgs int
    fShaType := flag.String("t", "256",
                   "Choose type of the SHA hash sum function: 256,384 or 512.")
    flag.Parse()

    if flag.NArg() > 0 {   // number of non optional args sent via cli directly
        noArgs = flag.NArg()
        strToHash = flag.Args()                   // non option args themselves
    } else {
        scanner := bufio.NewScanner(os.Stdin)
        for scanner.Scan() {                            // line scan by default
            noArgs++
            strToHash = append(strToHash, scanner.Text())
        }
        if err := scanner.Err(); err != nil {
            fmt.Println(err)
            os.Exit(1)
        }
    }

    switch *fShaType {
    case "256":
        for _, val := range strToHash {
            outputValue := sha256.Sum256([]byte(val))
            if _, err := fmt.Fprintf(os.Stdout, "%x\n", outputValue); err != nil {
                fmt.Println(err)
                os.Exit(1)
            }
        }
    case "384":
        for _, val := range strToHash {
            outputValue := sha512.Sum384([]byte(val))
            if _, err := fmt.Fprintf(os.Stdout, "%x\n", outputValue); err != nil {
                fmt.Println(err)
                os.Exit(1)
            }
        }
    case "512":
        for _, val := range strToHash {
            outputValue := sha512.Sum512([]byte(val))
            if _, err := fmt.Fprintf(os.Stdout, "%x\n", outputValue); err != nil {
                fmt.Println(err)
                os.Exit(1)
            }
        }
    default:                                                            //error
        flag.PrintDefaults()
        os.Exit(1)
    }
}
