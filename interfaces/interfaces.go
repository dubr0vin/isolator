package interfaces

import "flag"

type Module interface {
	RunAsHost(Host) error
	RunAsChild() error
	Settings(*flag.FlagSet)
}

type NamedModule interface {
	Module
	GetName() string
	GetDescription() string
}

type Host interface {
	AppendCloneFlag(uintptr)
}
