package uts

import (
	"flag"
	"github.com/dubr0vin/isolator/interfaces"
	"syscall"
)

func NewUTSModule() *module {
	return &module{}
}

type module struct {
}

func (*module) RunAsHost(h interfaces.Host) error {
	h.AppendCloneFlag(syscall.CLONE_NEWUTS)
	return nil
}

func (*module) RunAsChild() error {
	return nil
}

func (*module) Settings(_ *flag.FlagSet) {

}

func (*module) GetName() string {
	return "UTS"
}

func (*module) GetDescription() string {
	return "changing hostname"
}
