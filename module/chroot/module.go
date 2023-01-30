package chroot

import (
	"flag"
	"fmt"
	"github.com/dubr0vin/isolator/interfaces"
	"net/rpc"
	"os"
	"syscall"
)

func NewModule() *module {
	return &module{}
}

type module struct {
	rootfs *string
}

func (*module) RunAsHost(h interfaces.Host) error {
	h.GetSysProcAttrPtr().Cloneflags |= syscall.CLONE_NEWNS
	return nil
}

func (m *module) RunAsChild(_ *rpc.Client) error {
	if err := syscall.Mount(*m.rootfs, *m.rootfs, "", syscall.MS_BIND, ""); err != nil {
		return fmt.Errorf("mount %v", err)
	}
	if err := syscall.Chdir(*m.rootfs); err != nil {
		return fmt.Errorf("chdir %v", err)
	}
	if err := os.MkdirAll("oldrootfs", 0700); err != nil {
		return fmt.Errorf("mkdir %v", err)
	}
	if err := syscall.PivotRoot(".", "oldrootfs"); err != nil {
		return fmt.Errorf("pivot_root %v", err)
	}
	if err := syscall.Chdir("/"); err != nil {
		return fmt.Errorf("chdir %v", err)
	}
	if err := syscall.Unmount("/oldrootfs", syscall.MNT_DETACH); err != nil {
		return fmt.Errorf("unmount %v", err)
	}
	if err := os.RemoveAll("/oldrootfs"); err != nil {
		return err
	}

	return nil
}

func (m *module) Settings(flagSet *flag.FlagSet) {
	m.rootfs = flagSet.String("chroot-dir", "rootfs", "Path to new rootfs")
}

func (*module) GetName() string {
	return "chroot"
}

func (*module) GetDescription() string {
	return "using host rootfs without chroot"
}
