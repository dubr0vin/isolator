package main

import (
	"fmt"
	"net/rpc"
	"os"
	"syscall"
)

func childMain() {
	client, err := rpc.DialHTTP("tcp", os.Args[2])
	if err != nil {
		fmt.Printf("Error due to setup client: %s\n", err.Error())
		os.Exit(1)
	}
	ping(client)
	var args []string
	if err := client.Call("Host.GetArgs", 0, &args); err != nil {
		fmt.Printf("Error due to setup Host.GetArgs: %s\n", err.Error())
		os.Exit(1)
	}
	enabledModules, _, flagSet := getEnabledModules(args)
	for _, module := range enabledModules {
		if err := module.RunAsChild(); err != nil {
			fmt.Printf("Error due to %s.RunAsChild: %s\n", module.GetName(), err.Error())
			os.Exit(1)
		}
	}
	if err := syscall.Exec(flagSet.Arg(0), flagSet.Args()[0:], os.Environ()); err != nil {
		fmt.Printf("Error due to exec: %s", err.Error())
		os.Exit(1)
	}
}

func ping(client *rpc.Client) {
	ping := "test"
	var result string
	if err := client.Call("Host.Ping", ping, &result); err != nil {
		fmt.Printf("Error due to ping host: %s\n", err.Error())
		os.Exit(1)
	}
	if result != ping {
		fmt.Printf("Error due to ping host: \"%s\" != \"%s\"\n", result, ping)
		os.Exit(1)
	}
}
