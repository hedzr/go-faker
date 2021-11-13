package logic

import (
	"fmt"
	"github.com/hedzr/cmdr"
	"strings"
	"syreclabs.com/go/faker"
)

func genBitcoin(root *cmdr.RootCmdOpt) {
	root.NewSubCommand("bitcoin", "btc").
		Description("generate Bitcoin (BTC) names").
		Group("").
		TailPlaceholder("[text1, text2, ...]").
		Action(func(cmd *cmdr.Command, remainArgs []string) (err error) {
			oo := faker.Bitcoin()
			str := dumpBitcoin(oo)
			outputWithFormat(str, "Bitcoin")
			return
		})

}

func dumpBitcoin(oo faker.FakeBitcoin) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("    Address             : %v\n", oo.Address())) // => "Alphazap"
	sb.WriteString(fmt.Sprintf("    String              : %v\n", oo.String()))  // => "Tempsoft 4.51"
	return sb.String()
}
