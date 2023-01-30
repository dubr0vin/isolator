package interfaces

import (
	"flag"
	"net/rpc"
	"syscall"
)

type Module interface {
	RunAsHost(Host) error
	RunAsChild(*rpc.Client) error
	Settings(*flag.FlagSet)
}

type NamedModule interface {
	Module
	GetName() string
	GetDescription() string
}

type Host interface {
	GetSysProcAttrPtr() *syscall.SysProcAttr
}
