package logic

import (
	"fmt"
	"github.com/hedzr/cmdr"
	"syreclabs.com/go/faker"
)

func genNumber(root *cmdr.RootCmdOpt) {
	oo := faker.Number()
	//faker.Number().Number(5)          // => "43202"
	//faker.Number().NumberInt(3)       // => 213
	//faker.Number().NumberInt32(5)     // => 92938
	//faker.Number().NumberInt64(19)    // => 1689541633257139096
	//faker.Number().Decimal(8, 2)      // => "879420.60"
	//faker.Number().Digit()            // => "7"
	//faker.Number().Hexadecimal(4)     // => "e7f3"
	//faker.Number().Between(-100, 100) // => "-47"
	//faker.Number().Positive(100)      // => "3"
	//faker.Number().Negative(-100)     // => "-16"
	m := map[string]struct {
		label   string
		fn      func() string
		short   string
		aliases []string
	}{
		"number":       {"Number", func() string { return fmt.Sprintf("%v", oo.Number(cmdr.GetIntR("number.count"))) }, "n", []string{}},
		"number-int":   {"NumberInt", func() string { return fmt.Sprintf("%v", oo.NumberInt(cmdr.GetIntR("number.count"))) }, "ni", []string{}},
		"number-int32": {"NumberInt32", func() string { return fmt.Sprintf("%v", oo.NumberInt32(cmdr.GetIntR("number.count"))) }, "ni32", []string{}},
		"number-int64": {"NumberInt64", func() string { return fmt.Sprintf("%v", oo.NumberInt64(cmdr.GetIntR("number.count"))) }, "ni64", []string{}},
		"decimal": {"Decimal", func() string {
			return fmt.Sprintf("%v", oo.Decimal(cmdr.GetIntR("number.precision"), cmdr.GetIntR("number.scale")))
		}, "dec", []string{}},
		"digit":       {"Digit", func() string { return fmt.Sprintf("%v", oo.Digit()) }, "d", []string{}},
		"hexadecimal": {"Hexadecimal", func() string { return fmt.Sprintf("%v", oo.Hexadecimal(cmdr.GetIntR("number.count"))) }, "h", []string{"hex"}},
		"between": {"Between", func() string {
			return fmt.Sprintf("%v", oo.Between(cmdr.GetIntR("number.lower-bound"), cmdr.GetIntR("number.upper-bound")))
		}, "be", []string{}},
		"positive": {"Positive", func() string { return fmt.Sprintf("%v", oo.Positive(cmdr.GetIntR("number.upper-bound"))) }, "pp", []string{}},
		"negative": {"Negative", func() string { return fmt.Sprintf("%v", oo.Negative(cmdr.GetIntR("number.lower-bound"))) }, "nn", []string{}},
	}

	cc := root.NewSubCommand("number", "num", "num").
		Description("generate address").
		Group("").
		TailPlaceholder("[text1, text2, ...]").
		//PreAction(func(cmd *cmdr.Command, remainArgs []string) (err error) {
		//	fmt.Printf("[PRE] DebugMode=%v, TraceMode=%v. InDebugging/IsDebuggerAttached=%v\n",
		//		cmdr.GetDebugMode(), logex.GetTraceMode(), cmdr.InDebugging())
		//	for ix, s := range remainArgs {
		//		fmt.Printf("[PRE] %5d. %s\n", ix, s)
		//	}
		//
		//	fmt.Printf("[PRE] Debug=%v, Trace=%v\n", cmdr.GetDebugMode(), cmdr.GetTraceMode())
		//
		//	// return nil to be continue,
		//	// return cmdr.ErrShouldBeStopException to stop the following actions without error
		//	// return other errors for application purpose
		//	return
		//}).
		//PostAction(func(cmd *cmdr.Command, remainArgs []string) {
		//	for ix, s := range remainArgs {
		//		fmt.Printf("[POST] %5d. %s\n", ix, s)
		//	}
		//}).
		Action(func(cmd *cmdr.Command, remainArgs []string) (err error) {
			str := dumpIt(cmd, m)
			outputWithFormat(str, "Number")
			return
		})

	cmdr.NewInt(5).
		Titles("count", "C", "cnt").
		Group("").
		Placeholder("N").
		//Range(1,64).
		Description("number count for these fields: NumberXXX, Hexadecimal").
		AttachTo(cc)

	cmdr.NewInt(8).
		Titles("precision", "P").
		Group("").
		Placeholder("N").
		//Range(1,64).
		Description("float number precision for Decimal field").
		AttachTo(cc)

	cmdr.NewInt(2).
		Titles("scale", "S").
		Group("").
		Placeholder("N").
		//Range(1,64).
		Description("float number scale for Decimal field").
		AttachTo(cc)

	cmdr.NewInt(-100).
		Titles("lower-bound", "lb").
		Group("").
		Placeholder("N").
		//Range(1,64).
		Description("lower bound for these fields: Between, Positive, Negative").
		AttachTo(cc)

	cmdr.NewInt(100).
		Titles("upper-bound", "ub").
		Group("").
		Placeholder("N").
		//Range(1,64).
		Description("upper bound for these fields: Between, Positive, Negative").
		AttachTo(cc)

	for k, v := range m {
		cmdr.NewBool().
			Titles(k, v.short, v.aliases...).
			Description(fmt.Sprintf("generates %v field", v.label), "").
			ToggleGroup("Type").
			AttachTo(cc)
	}
}
