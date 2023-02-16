package cmdrrel

import (
	"fmt"

	"github.com/hedzr/cmdr"
	"github.com/hedzr/log"
	"github.com/hedzr/logex/build"
	"gopkg.in/hedzr/errors.v3"

	"github.com/hedzr/go-faker/cli/app"
	"github.com/hedzr/go-faker/internal/logic"
)

const (
	// appName   = "go-faker"
	copyright    = "go-faker - A faked records generator - cmdr series"
	desc         = "faker is an faked records generator"
	longDesc     = "faker is an faked records generator. It makes an demo application for `cmdr`"
	examples     = ``
	examplesLong = `
$ {{.AppName}} gen shell [--bash|--zsh|--auto]
  generate bash/shell completion scripts
$ {{.AppName}} gen man
  generate linux man page 1
$ {{.AppName}} --help
  show help screen.
`
	overview = ``

	zero = 0

	defaultTraceEnabled  = true
	defaultDebugEnabled  = true
	defaultLoggerBackend = "logrus"
	defaultLoggerLevel   = "debug"
)

func Entry() {

	if err := cmdr.Exec(buildRootCmd(),
		// trace.WithTraceEnable(defaultTraceEnabled),
		cmdr.WithUnhandledErrorHandler(onUnhandledErrorHandler),
		cmdr.WithLogx(build.New(build.NewLoggerConfigWith(defaultDebugEnabled, defaultLoggerBackend, defaultLoggerLevel))),
		cmdr.WithHelpTailLine(`
# Type '-h'/'-?' or '--help' to get command help screen.
# Star me if it's helpful: https://github.com/hedzr/go-hacker
		`),

		// dex.WithDaemon(
		//	svr.NewDaemon(svr.WithRouterImpl(sth.NewGinMux())),
		//	dex.WithCommandsModifier(modifier),
		//	dex.WithLoggerForward(true),
		// ),
		// server.WithCmdrDaemonSupport(),
		// server.WithCmdrHook(),

		optHideGenerateCmd,
		optAddTraceOption,
		// optAddServerExtOption,
		// pprof.GetCmdrProfilingOptions(),
	); err != nil {
		log.Fatalf("error occurs in app running: %+v\n", err)
	}
}

func onUnhandledErrorHandler(err interface{}) {
	if cmdr.GetBoolR("enable-ueh") {
		dumpStacks()
		// return
	}

	// internal.App().Close()

	panic(err)
}

func dumpStacks() {
	fmt.Printf("\n\n=== BEGIN goroutine stack dump ===\n%s\n=== END goroutine stack dump ===\n\n", errors.DumpStacksAsString(true))
}

func buildRootCmd() (rootCmd *cmdr.RootCommand) {
	root := cmdr.Root(app.AppName, app.Version).
		// AddGlobalPreAction(func(cmd *cmdr.Command, args []string) (err error) {
		//	// fmt.Printf("# global pre-action 1, curr-dir: %v\n", cmdr.GetCurrentDir())
		//	cmdr.Set("enable-ueh", true)
		//	return
		// }).
		// AddGlobalPreAction(func(cmd *cmdr.Command, args []string) (err error) {
		//	//fmt.Printf("# global pre-action 2, exe-path: %v\n", cmdr.GetExecutablePath())
		//	return
		// }).
		// AddGlobalPostAction(func(cmd *cmdr.Command, args []string) {
		//	//fmt.Println("# global post-action 1")
		// }).
		// AddGlobalPostAction(func(cmd *cmdr.Command, args []string) {
		//	//fmt.Println("# global post-action 2")
		// }).
		Copyright(copyright, "hedzr").
		Description(desc, longDesc).
		Examples(examples)
	rootCmd = root.RootCommand()

	// cmdr.NewBool(false).
	//	Titles("enable-ueh", "ueh").
	//	Description("Enables the unhandled exception handler?").
	//	AttachTo(root)

	// core.AttachToCmdr(root.RootCmdOpt())

	logic.AttachToCmdr(root.RootCmdOpt())

	// soundex(root)
	// panicTest(root)

	// pprof.AttachToCmdr(root.RootCmdOpt())
	return
}
