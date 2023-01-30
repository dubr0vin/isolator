package uts

import (
	"flag"
	"github.com/dubr0vin/isolator/interfaces"
	"syscall"
)

func NewModule() *module {
	return &module{}
}

type module struct {
}

func (*module) RunAsHost(h interfaces.Host) error {
	h.GetSysProcAttrPtr().Cloneflags |= syscall.CLONE_NEWUTS
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
