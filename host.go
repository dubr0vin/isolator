package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"os/exec"
	"syscall"
	"time"
)

var childProcess *os.Process

func hostMain() {
	enabledModules, _ := getEnabledModules(os.Args)
	h := &host{}

	for _, module := range enabledModules {
		if err := module.RunAsHost(h); err != nil {
			fmt.Printf("Error due to %s.RunAsHost: %s\n", module.GetName(), err.Error())
			os.Exit(1)
		}
	}

	address := fmt.Sprintf("/tmp/isolator-%d.sock", os.Getpid())

	startServer(address)
	cmd := exec.Command("/proc/self/exe", "child", address)
	cmd.SysProcAttr = &h.SysProcAttr

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		fmt.Printf("Error due to start child process: %s\n", err.Error())
		os.Exit(1)
	}
	childProcess = cmd.Process
	if err := cmd.Wait(); err != nil {
		fmt.Printf("Error due to wair child process: %s\n", err.Error())
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

	listener, err := net.Listen("unix", address)
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

func (h *HostServer) StartTimer(args time.Duration, response *int) error {
	go func() {
		time.Sleep(args)
		if err := syscall.Kill(childProcess.Pid, syscall.SIGKILL); err != nil {
			fmt.Printf("Error due to kill %v\n", err)
		}
		os.Exit(2)
	}()
	return nil
}
