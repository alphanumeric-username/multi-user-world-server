package main

import (
	"fmt"
	"os"

	"github.com/alphanumeric-username/muwsvr/cli"
)

func main() {

	command := os.Args[1]

	switch command {
	case "start":
		args, flagSet, err := cli.ParseStart(os.Args[2:])

		if err != nil {
			flagSet.Usage()
			os.Exit(1)
		}

		if args.Help {
			flagSet.Usage()
		} else {
			serverConfig, err := cli.LoadServerConfig(args.ConfigFile)

			if err != nil {
				fmt.Println("Error while loading server configuration file:")
				fmt.Printf("\t%s", err.Error())
			}

			fmt.Println("Openning the server with")
			fmt.Printf("\tip: %s\n", serverConfig.IP)
			fmt.Printf("\thttp_port: %d\n", serverConfig.HTTPPort)
			fmt.Printf("\tudp_port: %d\n", serverConfig.UDPPort)
			fmt.Printf("\tworld_file: %s\n", serverConfig.WorldFile)
			fmt.Printf("\tmax_users: %d\n", serverConfig.MaxUsers)
			fmt.Printf("\tenable_chat: %t\n", serverConfig.EnableChat)
			fmt.Printf("\tmax_messages: %d\n", serverConfig.MaxMessages)
			fmt.Println("Closing server...")

		}
		break
	case "version":
		args, flagSet, err := cli.ParseVersion(os.Args[2:])

		if err != nil {
			flagSet.Usage()
			os.Exit(1)
		}

		if args.Help {
			flagSet.Usage()
		} else {
			fmt.Println(args.Version)
		}
		break
	}
}
