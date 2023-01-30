package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"os/exec"
	"syscall"
)

func hostMain() {
	enabledModules, port, _ := getEnabledModules(os.Args)
	h := &host{}

	for _, module := range enabledModules {
		if err := module.RunAsHost(h); err != nil {
			fmt.Printf("Error due to %s.RunAsHost: %s\n", module.GetName(), err.Error())
			os.Exit(1)
		}
	}

	address := fmt.Sprintf("127.0.0.1:%d", port)

	startServer(address)
	cmd := exec.Command("/proc/self/exe", "child", address)
	cmd.SysProcAttr = &h.SysProcAttr

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error due to run child process: %s\n", err.Error())
		os.Exit(1)
	}

}

type host struct {
	SysProcAttr syscall.SysProcAttr
}

func (h *host) GetSysProcAttrPtr() *syscall.SysProcAttr {
	return &h.SysProcAttr
}

func startServer(address string) {
	serve := HostServer{}
	if err := rpc.RegisterName("Host", &serve); err != nil {
		fmt.Printf("Error due to register host rpc: %s\n", err.Error())
		os.Exit(1)
	}
	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Printf("Error due to listen: %s\n", err.Error())
		os.Exit(1)
	}
	go func() {
		err := http.Serve(listener, nil)
		if err != nil {
			fmt.Printf("Error due to start server: %s\n", err.Error())
			os.Exit(1)
		}
	}()
}

type HostServer struct {
}

func (h *HostServer) Ping(args string, response *string) error {
	*response = args
	return nil
}

func (h *HostServer) GetArgs(args int, response *[]string) error {
	*response = os.Args
	return nil
}
