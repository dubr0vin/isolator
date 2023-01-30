package memory

import (
	"flag"
	"fmt"
	"github.com/dubr0vin/isolator/interfaces"
	"os"
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
	RlimitRss    = 5
	DefaultLimit = 64 * 1024 * 1024 //64Mb
)

func (m *module) RunAsChild() error {
	limit := *m.limit / uint64(os.Getpagesize())
	if err := syscall.Setrlimit(RlimitRss, &syscall.Rlimit{
		Cur: *m.limit / limit,
		Max: *m.limit / limit,
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
