package memory

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

func (*module) RunAsHost(h interfaces.Host) error {
	return nil
}

const (
	DefaultLimit = 64 * 1024 * 1024 //64Mb
)

func (m *module) RunAsChild(_ *rpc.Client) error {
	if err := syscall.Setrlimit(syscall.RLIMIT_AS, &syscall.Rlimit{
		Cur: *m.limit,
		Max: *m.limit,
	}); err != nil {
		return fmt.Errorf("setrlimit %v", err)
	}
	return nil
}

func (m *module) Settings(flagSet *flag.FlagSet) {
	m.limit = flagSet.Uint64("memory-limit", DefaultLimit, "Memory limit in bytes")
}

func (*module) GetName() string {
	return "memory"
}

func (*module) GetDescription() string {
	return "using unlimited memory"
}
