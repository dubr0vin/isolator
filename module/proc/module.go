package proc

import (
	"flag"
	"fmt"
	"github.com/dubr0vin/isolator/interfaces"
	"net/rpc"
	"syscall"
)

func NewModule() *module {
	return &module{}
}

type module struct {
	limit *uint64
}

func (m *module) RunAsHost(h interfaces.Host) error {
	return nil
}

const (
	RlimitNproc = 6
)

func (m *module) RunAsChild(_ *rpc.Client) error {
	if err := syscall.Setrlimit(RlimitNproc, &syscall.Rlimit{
		Cur: *m.limit,
		Max: *m.limit,
	}); err != nil {
		return fmt.Errorf("setrlimit %v", err)
	}

	return nil
}

func (m *module) Settings(flagSet *flag.FlagSet) {
	m.limit = flagSet.Uint64("proc-limit", 1, "Maximum amount of processors")
}

func (*module) GetName() string {
	return "proc"
}

func (*module) GetDescription() string {
	return "using unlimited processes"
}
