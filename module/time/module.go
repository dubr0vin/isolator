package time

import (
	"flag"
	"fmt"
	"github.com/dubr0vin/isolator/interfaces"
	"net/rpc"
	"syscall"
	"time"
)

func NewModule() *module {
	return &module{}
}

type module struct {
	cpuLimit  *time.Duration
	realLimit *time.Duration
}

func (m *module) RunAsHost(h interfaces.Host) error {
	return nil
}

func (m *module) RunAsChild(client *rpc.Client) error {
	limit := uint64(m.cpuLimit.Seconds())
	if err := syscall.Setrlimit(syscall.RLIMIT_CPU, &syscall.Rlimit{
		Cur: limit,
		Max: limit,
	}); err != nil {
		return fmt.Errorf("setrlimit %v", err)
	}

	if err := client.Call("Host.StartTimer", *m.realLimit, nil); err != nil {
		return fmt.Errorf("Host.StartTimer %v", err)
	}

	return nil
}

func (m *module) Settings(flagSet *flag.FlagSet) {
	m.cpuLimit = flagSet.Duration("cpu-time-limit", 2*time.Second, "CPU time limit")
	m.realLimit = flagSet.Duration("real-time-limit", 5*time.Second, "Real time limit")
}

func (*module) GetName() string {
	return "time"
}

func (*module) GetDescription() string {
	return "using unlimited cpu time"
}
