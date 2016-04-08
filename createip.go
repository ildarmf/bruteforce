package main

import (
    "fmt"
)

func main() {
    for i:=1; i<256; i++ {
        fmt.Printf("192.168.102.%d\n", i)
    }
}
