package logic

import (
	"fmt"
	"github.com/hedzr/cmdr"
	"syreclabs.com/go/faker"
)

func genCommerce(root *cmdr.RootCmdOpt) {
	oo := faker.Commerce()
	m := map[string]struct {
		label   string
		fn      func() string
		short   string
		aliases []string
	}{
		"product-name": {"ProductName", oo.ProductName, "pn", []string{}},
		"department":   {"Department", oo.Department, "d", []string{"dept"}},
		"price":        {"Price", func() string { return fmt.Sprintf("%v", oo.Price()) }, "pr", []string{}},
		"color":        {"Color", oo.Color, "c", []string{}},
	}

	//sb.WriteString(fmt.Sprintf("    ProductName        : %v\n", oo.ProductName()))
	//sb.WriteString(fmt.Sprintf("    Department         : %v\n", oo.Department()))
	//sb.WriteString(fmt.Sprintf("    Price              : %v\n", oo.Price()))
	//sb.WriteString(fmt.Sprintf("    Color              : %v\n", oo.Color()))

	cc := root.NewSubCommand("commerce", "cc").
		Description("generate Commerce names").
		Group("").
		TailPlaceholder("[text1, text2, ...]").
		Action(func(cmd *cmdr.Command, remainArgs []string) (err error) {
			str := dumpIt(cmd, m)
			outputWithFormat(str, "Commerce")
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
