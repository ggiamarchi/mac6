package main

import (
	"fmt"
	"os"

	"github.com/ggiamarchi/mac6/ipv6"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println()
		fmt.Println("mac6 compute IPv6 link-local address from MAC address")
		fmt.Println()
		fmt.Println("Usage : mac6 <MAC Address>")
		fmt.Println()
		os.Exit(1)
	}

	ip, err := ipv6.Compute(os.Args[1])

	if err != nil {
		fmt.Printf("Error : %s\n", err)
		os.Exit(1)
	}

	fmt.Println(ip)
}
