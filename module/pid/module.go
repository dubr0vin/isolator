package pid

import (
	"flag"
	"github.com/dubr0vin/isolator/interfaces"
	"syscall"
)

func NewPidModule() *module {
	return &module{}
}

type module struct {
}

func (*module) RunAsHost(h interfaces.Host) error {
	h.AppendCloneFlag(syscall.CLONE_NEWPID)
	return nil
}

func (*module) RunAsChild() error {
	return nil
}

func (*module) Settings(_ *flag.FlagSet) {

}

func (*module) GetName() string {
	return "PID"
}

func (*module) GetDescription() string {
	return "using host pid namespace"
}