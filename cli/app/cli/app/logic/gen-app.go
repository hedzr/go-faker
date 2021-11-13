package logic

import (
	"fmt"
	"github.com/hedzr/cmdr"
	"strings"
	"syreclabs.com/go/faker"
)

func genApp(root *cmdr.RootCmdOpt) {
	root.NewSubCommand("app", "app").
		Description("generate app names").
		Group("").
		TailPlaceholder("[text1, text2, ...]").
		Action(func(cmd *cmdr.Command, remainArgs []string) (err error) {
			oo := faker.App()
			str := dumpApp(oo)
			outputWithFormat(str, "App")
			return
		})

}

func dumpApp(oo faker.FakeApp) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("    Name                : %v\n", oo.Name()))    // => "Alphazap"
	sb.WriteString(fmt.Sprintf("    Version             : %v\n", oo.Version())) // => "2.6.0"
	sb.WriteString(fmt.Sprintf("    Author              : %v\n", oo.Author()))  // => "Dorian Shields"
	sb.WriteString(fmt.Sprintf("    String              : %v\n", oo.String()))  // => "Tempsoft 4.51"
	return sb.String()
}
