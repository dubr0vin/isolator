package network

import (
	"flag"
	"github.com/dubr0vin/isolator/interfaces"
	"net/rpc"
	"syscall"
)

func NewModule() *module {
	return &module{}
}

type module struct {
}

func (*module) RunAsHost(h interfaces.Host) error {
	h.GetSysProcAttrPtr().Cloneflags |= syscall.CLONE_NEWNET
	return nil
}

func (*module) RunAsChild(_ *rpc.Client) error {
	return nil
}

func (*module) Settings(_ *flag.FlagSet) {

}

func (*module) GetName() string {
	return "network"
}

func (*module) GetDescription() string {
	return "using network"
}
