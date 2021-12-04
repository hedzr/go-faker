package cmdrrel

import "github.com/hedzr/cmdr"

var optHideGenerateCmd, optAddTraceOption cmdr.ExecOption

func init() {

	// hide generate command
	optHideGenerateCmd = cmdr.WithXrefBuildingHooks(nil, func(root *cmdr.RootCommand, args []string) {
		if cc := cmdr.FindSubCommand("generate", &root.Command); cc != nil {
			cc.Hidden = true
		}
	})

	// attaches `--trace` to root command
	// deprecated: instead of `trace.WithTraceEnable(true)`
	optAddTraceOption = cmdr.WithXrefBuildingHooks(func(root *cmdr.RootCommand, args []string) {
		cmdr.NewBool(false).
			Titles("trace", "tr").
			Description("enable trace mode for tcp/mqtt send/recv data dump", "").
			AttachToRoot(root)
	}, nil)

	//// the following statements show you how to attach an option to a sub-command
	//optAddServerExtOption = cmdr.WithXrefBuildingHooks(func(root *cmdr.RootCommand, args []string) {
	//	serverCmd := cmdr.FindSubCommandRecursive("server", nil)
	//	serverStartCmd := cmdr.FindSubCommand("start", serverCmd)
	//	cmdr.NewInt(5100).
	//		Titles("vnc-server", "vnc").
	//		Description("start as a vnc server (just a faked demo)", "").
	//		Placeholder("PORT").
	//		AttachTo(cmdr.NewCmdFrom(serverStartCmd))
	//}, nil)
}
