package logic

import (
	"fmt"
	"github.com/hedzr/cmdr"
	"syreclabs.com/go/faker"
)

func genBitcoin(root *cmdr.RootCmdOpt) {
	oo := faker.Bitcoin()
	m := map[string]struct {
		label   string
		fn      func() string
		short   string
		aliases []string
	}{
		"address": {"Address", oo.Address, "a", []string{"addr"}},
		"string":  {"String", oo.String, "ss", []string{}},
	}
	//sb.WriteString(fmt.Sprintf("    Address             : %v\n", oo.Address())) // => "Alphazap"
	//sb.WriteString(fmt.Sprintf("    String              : %v\n", oo.String()))  // => "Tempsoft 4.51"

	cc := root.NewSubCommand("bitcoin", "btc").
		Description("generate Bitcoin (BTC) names").
		Group("").
		TailPlaceholder("[text1, text2, ...]").
		Action(func(cmd *cmdr.Command, remainArgs []string) (err error) {
			str := dumpIt(cmd, m)
			outputWithFormat(str, "Bitcoin")
			return
		})
	for k, v := range m {
		cmdr.NewBool().
			Titles(k, v.short, v.aliases...).
			Description(fmt.Sprintf("generates %v field", v.label), "").
			ToggleGroup("Type").
			AttachTo(cc)
	}
}
