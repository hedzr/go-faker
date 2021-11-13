package logic

import (
	"fmt"
	"github.com/hedzr/cmdr"
	"strings"
	"syreclabs.com/go/faker"
)

func genCommerce(root *cmdr.RootCmdOpt) {
	root.NewSubCommand("commerce", "cc").
		Description("generate Commerce names").
		Group("").
		TailPlaceholder("[text1, text2, ...]").
		Action(func(cmd *cmdr.Command, remainArgs []string) (err error) {
			oo := faker.Commerce()
			str := dumpCommerce(oo)
			outputWithFormat(str, "Commerce")
			return
		})
}

func dumpCommerce(oo faker.FakeCommerce) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("    ProductName        : %v\n", oo.ProductName()))
	sb.WriteString(fmt.Sprintf("    Department         : %v\n", oo.Department()))
	sb.WriteString(fmt.Sprintf("    Price              : %v\n", oo.Price()))
	sb.WriteString(fmt.Sprintf("    Color              : %v\n", oo.Color()))
	return sb.String()
}
