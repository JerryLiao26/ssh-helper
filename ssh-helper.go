package main

import (
	"flag"
	"fmt"

	"github.com/JerryLiao26/ssh-helper/cli"
	"github.com/JerryLiao26/ssh-helper/helper"
)

/* Commands */
var (
	envComm   bool
	helpComm  bool
	killComm  bool
	tidyComm  bool
	startComm bool

	addComm string
)

func main() {
	// Load config
	helper.LoadConf()

	// Validate config
	helper.ValidateConf()

	// Handle command-line flags
	flag.Parse()

	if helpComm {
		flag.Usage()
	} else if envComm {
		cli.EnvHandler()
	} else if killComm {
		cli.KillHandler()
	} else if tidyComm {
		cli.TidyHandler()
	} else if startComm {
		cli.StartHandler()
	} else if addComm != "" {
		cli.AddHandler(addComm)
	} else {
		flag.Usage()
	}
}

func init() {
	flag.BoolVar(&helpComm, "h", false, "Print help text")
	flag.BoolVar(&envComm, "e", false, "Print ssh-agent environment variables")
	flag.BoolVar(&killComm, "k", false, "Kill all ssh-agent processes")
	flag.BoolVar(&tidyComm, "t", false, "Kill all ssh-agent processes except managed one")
	flag.BoolVar(&startComm, "s", false, "Start a managed ssh-agent process, or show PID if one is already running")
	flag.StringVar(&addComm, "a", "", "Wrapper of ssh-add, for detailed help of this command execute \"man ssh-add\"")

	flag.Usage = usage
}

func usage() {
	_, _ = fmt.Println(`SSH-Helper Usage:`)

	flag.PrintDefaults()
}
