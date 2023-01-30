package main

import (
	"flag"
	"fmt"
	"github.com/dubr0vin/isolator/interfaces"
	"github.com/dubr0vin/isolator/module/pid"
	"github.com/dubr0vin/isolator/module/uts"
	"os"
)

var allModules = []interfaces.NamedModule{
	uts.NewUTSModule(),
	pid.NewPidModule(),
}

func getEnabledModules(args []string) ([]interfaces.NamedModule, int, *flag.FlagSet) {
	flagSet := flag.NewFlagSet(args[0], flag.ExitOnError)

	disabledModules := make(map[string]*bool)
	for _, module := range allModules {
		disabledModules[module.GetName()] = flagSet.Bool("disable-"+module.GetName(), false, "Allow "+module.GetDescription()+" for isolated process")
		module.Settings(flagSet)
	}
	port := flagSet.Int("rpc-port", 1234, "Port for client rpc server")
	if err := flagSet.Parse(args[1:]); err != nil {
		fmt.Printf("Error due to parse args: %s\n", err.Error())
		os.Exit(1)
	}

	enabledModules := make([]interfaces.NamedModule, 0, len(allModules))
	for _, module := range allModules {
		if *disabledModules[module.GetName()] {
			continue
		}
		enabledModules = append(enabledModules, module)
	}
	return allModules, *port, flagSet
}
