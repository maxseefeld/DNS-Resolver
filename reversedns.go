package main

import (
    "fmt"
    "net"
)

func main() {
    ip := net.ParseIP("8.8.8.8")
    if ip == nil {
        fmt.Println("Invalid IP address")
        return
    }

    domains, err := net.LookupAddr(ip.String())
    if err != nil {
        fmt.Println("Reverse DNS lookup failed:", err)
    } else {
        for _, domain := range domains {
            fmt.Println("Domain name:", domain)
        }
    }
}
