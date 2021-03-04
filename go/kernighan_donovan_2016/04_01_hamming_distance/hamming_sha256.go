// measure of Hamming distance between two SHA256 hashes
package main

import (
    "fmt"
    "crypto/sha256"
    "errors"
    "strconv"
)

// WEGNER algo for counting set bits:
func distance(first, second [32]byte) (output int, err error) {
    if len(first) != len(second) {
        return -1, errors.New("Difference in lenghts of arguments is illegal.")
    }
    var intFirst, intSecond uint64
    step := len(first) / 4                               // to prevent overflow
    for i := 0; i < 4; i++ {
        strFirst := fmt.Sprintf("%x", string(first[(i * step):((i + 1) * step)]))
        strSecond := fmt.Sprintf("%x", string(second[(i * step):((i + 1) * step)]))
        if intFirst, err = strconv.ParseUint(strFirst, 16, 64); err != nil {
            return -1, err
        }
        if intSecond, err = strconv.ParseUint(strSecond, 16, 64); err != nil {
            return -1, err
        }
        for i := intFirst ^ intSecond; i > 0; i = i & (i - 1) { output++ }
    }
    return
}

func main() {
    inputA := []byte("Putin chuj złamany.")
    inputB := []byte("Putin chuj mały.")
    A := sha256.Sum256(inputA)
    B := sha256.Sum256(inputB)
    if dist, err := distance(A, B); err == nil {
        fmt.Printf("A: %x\t B: %x\t dist:%d\n", A, B, dist)
    } else {
        fmt.Println(err)
    }
}
