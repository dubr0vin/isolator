package user

import (
	"flag"
	"fmt"
	"github.com/dubr0vin/isolator/interfaces"
	"syscall"
)

func NewModule() *module {
	return &module{}
}

type module struct {
}

func (*module) RunAsHost(h interfaces.Host) error {
	h.GetSysProcAttrPtr().Cloneflags |= syscall.CLONE_NEWUSER
	h.GetSysProcAttrPtr().UidMappings = append(h.GetSysProcAttrPtr().UidMappings, syscall.SysProcIDMap{
		ContainerID: 0,
		HostID:      syscall.Getuid(),
		Size:        1,
	})
	h.GetSysProcAttrPtr().GidMappings = append(h.GetSysProcAttrPtr().GidMappings, syscall.SysProcIDMap{
		ContainerID: 0,
		HostID:      syscall.Getgid(),
		Size:        1,
	})
	return nil
}

func (*module) RunAsChild() error {
	if err := syscall.Setuid(0); err != nil {
		return fmt.Errorf("setuid %v", err)
	}
	if err := syscall.Setgid(0); err != nil {
		return fmt.Errorf("setgid %v", err)
	}
	return nil
}

func (*module) Settings(_ *flag.FlagSet) {

}

func (*module) GetName() string {
	return "uid"
}

func (*module) GetDescription() string {
	return "using hosts UIDs and GIDs"
}
