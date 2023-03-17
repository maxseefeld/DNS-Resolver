package main

import (
    "fmt"
    "net"
)

func main() {
    cname, err := net.LookupCNAME("www.example.com")
    if err != nil {
        fmt.Println("CNAME lookup failed:", err)
    } else {
        fmt.Println("CNAME:", cname)
    }

    ips, err := net.LookupIP("www.example.com")
    if err != nil {
        fmt.Println("IP lookup failed:", err)
    } else {
        for _, ip := range ips {
            fmt.Println("IP address:", ip.String())
        }
    }
}
