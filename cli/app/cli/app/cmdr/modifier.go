package cmdrrel

import "github.com/hedzr/cmdr"

func modifier(daemonServerCommands *cmdr.Command) *cmdr.Command {
	if startCmd := daemonServerCommands.FindSubCommand("start"); startCmd != nil {
		startCmd.PreAction = onServerPreStart
		startCmd.PostAction = onServerPostStop
	}

	return daemonServerCommands
}
