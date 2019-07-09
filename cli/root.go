package cli

import (
	"github.com/DataDrake/cli-ng/cmd"
)

//GlobalFlags contains the flags for commands.
type GlobalFlags struct{}

// Root is the main command.
var Root *cmd.RootCMD

// init creates the command interface and registers the possible commands.
func init() {
	Root = &cmd.RootCMD{
		Name:  "lookout",
		Short: "Lookout is an Upstream Update Watcher",
		Flags: &GlobalFlags{},
	}
	Root.RegisterCMD(&cmd.Help)
	Root.RegisterCMD(&Add)
	Root.RegisterCMD(&Search)
	Root.RegisterCMD(&Run)
	Root.RegisterCMD(&Import)
}
