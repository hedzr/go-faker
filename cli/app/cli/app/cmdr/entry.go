package cmdrrel

import (
	"fmt"
	"github.com/hedzr/cmdr"
	"github.com/hedzr/cmdr/tool"
	"github.com/hedzr/go-faker/cli/app"
	"github.com/hedzr/go-faker/cli/app/cli/app/logic"
	"github.com/hedzr/go-faker/internal"
	"github.com/hedzr/log"
	"github.com/hedzr/logex"
	"github.com/hedzr/logex/build"
	"gopkg.in/hedzr/errors.v2"
)

const (
	// appName   = "go-faker"
	copyright = "go-faker - A faked records generator - cmdr series"
	desc      = "faker is an faked records generator"
	longDesc  = "faker is an faked records generator. It makes an demo application for `cmdr`"
	examples  = `
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
		//trace.WithTraceEnable(defaultTraceEnabled),
		cmdr.WithUnhandledErrorHandler(onUnhandledErrorHandler),
		cmdr.WithLogx(build.New(build.NewLoggerConfigWith(defaultDebugEnabled, defaultLoggerBackend, defaultLoggerLevel))),
		//cmdr.WithHelpTailLine(`
		//
		//`),

		//dex.WithDaemon(
		//	svr.NewDaemon(svr.WithRouterImpl(sth.NewGinMux())),
		//	dex.WithCommandsModifier(modifier),
		//	dex.WithLoggerForward(true),
		//),
		// server.WithCmdrDaemonSupport(),
		// server.WithCmdrHook(),

		optHideGenerateCmd,
		//optAddTraceOption,
		//optAddServerExtOption,
		//pprof.GetCmdrProfilingOptions(),
	); err != nil {
		log.Fatalf("error occurs in app running: %+v\n", err)
	}
}

func onUnhandledErrorHandler(err interface{}) {
	if cmdr.GetBoolR("enable-ueh") {
		dumpStacks()
		// return
	}

	internal.App().Close()

	panic(err)
}

func dumpStacks() {
	fmt.Printf("\n\n=== BEGIN goroutine stack dump ===\n%s\n=== END goroutine stack dump ===\n\n", errors.DumpStacksAsString(true))
}

func buildRootCmd() (rootCmd *cmdr.RootCommand) {
	root := cmdr.Root(app.AppName, app.Version).
		//AddGlobalPreAction(func(cmd *cmdr.Command, args []string) (err error) {
		//	// fmt.Printf("# global pre-action 1, curr-dir: %v\n", cmdr.GetCurrentDir())
		//	cmdr.Set("enable-ueh", true)
		//	return
		//}).
		//AddGlobalPreAction(func(cmd *cmdr.Command, args []string) (err error) {
		//	//fmt.Printf("# global pre-action 2, exe-path: %v\n", cmdr.GetExecutablePath())
		//	return
		//}).
		//AddGlobalPostAction(func(cmd *cmdr.Command, args []string) {
		//	//fmt.Println("# global post-action 1")
		//}).
		//AddGlobalPostAction(func(cmd *cmdr.Command, args []string) {
		//	//fmt.Println("# global post-action 2")
		//}).
		Copyright(copyright, "hedzr").
		Description(desc, longDesc).
		Examples(examples)
	rootCmd = root.RootCommand()

	cmdr.NewBool(false).
		Titles("enable-ueh", "ueh").
		Description("Enables the unhandled exception handler?").
		AttachTo(root)

	// core.AttachToCmdr(root.RootCmdOpt())

	logic.AttachToCmdr(root.RootCmdOpt())

	//soundex(root)
	//panicTest(root)

	//pprof.AttachToCmdr(root.RootCmdOpt())
	return
}

func prd(key string, val interface{}, format string, params ...interface{}) {
	fmt.Printf("         [--%v] %v, %v\n", key, val, fmt.Sprintf(format, params...))
}

func soundex(root cmdr.OptCmd) {
	// soundex

	parent := root.NewSubCommand("soundex", "snd", "sndx", "sound").
		Description("soundex test").
		Group("Test").
		TailPlaceholder("[text1, text2, ...]").
		PreAction(func(cmd *cmdr.Command, remainArgs []string) (err error) {
			fmt.Printf("[PRE] DebugMode=%v, TraceMode=%v. InDebugging/IsDebuggerAttached=%v\n",
				cmdr.GetDebugMode(), logex.GetTraceMode(), cmdr.InDebugging())
			for ix, s := range remainArgs {
				fmt.Printf("[PRE] %5d. %s\n", ix, s)
			}

			fmt.Printf("[PRE] Debug=%v, Trace=%v\n", cmdr.GetDebugMode(), cmdr.GetTraceMode())

			// return nil to be continue,
			// return cmdr.ErrShouldBeStopException to stop the following actions without error
			// return other errors for application purpose
			return
		}).
		Action(func(cmd *cmdr.Command, remainArgs []string) (err error) {
			for ix, s := range remainArgs {
				// fmt.Printf("[ACTION] %5d. %s\n", ix, s)
				fmt.Printf("[ACTION] %5d. %s => %s\n", ix, s, tool.Soundex(s))
			}

			prd("bool", cmdr.GetBoolR("soundex.bool"), "")
			prd("int", cmdr.GetIntR("soundex.int"), "")
			prd("int64", cmdr.GetInt64R("soundex.int64"), "")
			prd("uint", cmdr.GetUintR("soundex.uint"), "")
			prd("uint64", cmdr.GetUint64R("soundex.uint64"), "")
			prd("float32", cmdr.GetFloat32R("soundex.float32"), "")
			prd("float64", cmdr.GetFloat64R("soundex.float64"), "")
			prd("complex64", cmdr.GetComplex64R("soundex.complex64"), "")
			prd("complex128", cmdr.GetComplex128R("soundex.complex128"), "")

			prd("single", cmdr.GetBoolR("soundex.single"), "")
			prd("double", cmdr.GetBoolR("soundex.double"), "")
			prd("norway", cmdr.GetBoolR("soundex.norway"), "")
			prd("mongo", cmdr.GetBoolR("soundex.mongo"), "")
			return
		}).
		PostAction(func(cmd *cmdr.Command, remainArgs []string) {
			for ix, s := range remainArgs {
				fmt.Printf("[POST] %5d. %s\n", ix, s)
			}
		})

	cmdr.NewBool(false).
		Titles("bool", "b").
		Description("A bool flag", "").
		Group("").
		EnvKeys("").
		AttachTo(parent)

	cmdr.NewInt(1).
		Titles("int", "i", "i32").
		Description("A int flag", "").
		Group("1000.Integer").
		EnvKeys("").
		AttachTo(parent)
	cmdr.NewInt64(2).
		Titles("int64", "i64").
		Description("A int64 flag", "").
		Group("1000.Integer").
		EnvKeys("").
		AttachTo(parent)
	cmdr.NewUint(3).
		Titles("uint", "u", "u32").
		Description("A uint flag", "").
		Group("1000.Integer").
		EnvKeys("").
		AttachTo(parent)
	cmdr.NewUint64(4).
		Titles("uint64", "u64").
		Description("A uint64 flag", "").
		Group("1000.Integer").
		EnvKeys("").
		AttachTo(parent)

	cmdr.NewFloat32(2.71828).
		Titles("float32", "f", "float", "f32").
		Description("A float32 flag with 'e' value", "").
		Group("2000.Float").
		EnvKeys("E", "E2").
		AttachTo(parent)
	cmdr.NewFloat64(3.14159265358979323846264338327950288419716939937510582097494459230781640628620899).
		Titles("float64", "f64").
		Description("A float64 flag with a `PI` value", "").
		Group("2000.Float").
		EnvKeys("PI").
		AttachTo(parent)
	cmdr.NewComplex64(3.14+9i).
		Titles("complex64", "c64").
		Description("A complex64 flag", "").
		Group("2010.Complex").
		EnvKeys("").
		AttachTo(parent)
	cmdr.NewComplex128(3.14+9i).
		Titles("complex128", "c128").
		Description("A complex128 flag", "").
		Group("2010.Complex").
		EnvKeys("").
		AttachTo(parent)

	cmdr.NewBool(false).
		Titles("single", "s").
		Description("A bool flag: single", "").
		Group("Boolean").
		EnvKeys("").
		AttachTo(parent)

	cmdr.NewBool(false).
		Titles("double", "d").
		Description("A bool flag: double", "").
		Group("Boolean").
		EnvKeys("").
		AttachTo(parent)

	cmdr.NewBool(false).
		Titles("norway", "n", "nw").
		Description("A bool flag: norway", "").
		Group("Boolean").
		EnvKeys("").
		AttachTo(parent)

	cmdr.NewBool(false).
		Titles("mongo", "m").
		Description("A bool flag: mongo", "").
		Group("Boolean").
		EnvKeys("").
		AttachTo(parent)

}

func panicTest(root cmdr.OptCmd) {
	// panic test

	pa := root.NewSubCommand("panic-test", "pa", "panic").
		Description("test panic inside cmdr actions", "").
		Group("Test")

	val := 9
	zeroVal := zero
	slice1 := []int{1, 2, 3}

	pa.NewSubCommand("slice-bound-out-of-range", "sb", "sboor").
		Description("").
		Group("Test").
		Action(func(cmd *cmdr.Command, args []string) (err error) {
			fmt.Println(slice1[100])
			return
		})

	pa.NewSubCommand("division-by-zero", "dz").
		Description("").
		Group("Test").
		Action(func(cmd *cmdr.Command, args []string) (err error) {
			fmt.Println(val / zeroVal)
			return
		})

	pa.NewSubCommand("panic", "pa").
		Description("").
		Group("Test").
		Action(func(cmd *cmdr.Command, args []string) (err error) {
			panic(9)
			return
		})
}
