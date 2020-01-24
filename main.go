package main

import (
	"fmt"
	"os"

	"github.com/ggiamarchi/mac6/ipv6"
	cli "github.com/jawher/mow.cli"
)

func main() {
	app := cli.App("mac6", "Compute IPv6 link-local address from MAC address")

	app.Spec = "[ -i=<interface> ] MAC"

	outInterface := app.StringOpt("i interface", "", "Out interface to build IPv6 link local string")
	macAddress := app.StringArg("MAC", "", "Mac address")

	app.Action = func() {
		ip, err := ipv6.ComputeLinkLocalAddress(*macAddress, *outInterface)

		if err != nil {
			fmt.Printf("Error : %s\n", err)
			os.Exit(1)
		}
		fmt.Println(ip)
	}

	app.Run(os.Args)
}
