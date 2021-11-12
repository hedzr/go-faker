package logic

import (
	"fmt"
	"github.com/hedzr/cmdr"
	"syreclabs.com/go/faker"
)

func genFinance(root *cmdr.RootCmdOpt) {

	cc := root.NewSubCommand("finance", "f").
		Description("generate Finance names (visa, mastercard, ...)").
		Group("").
		TailPlaceholder("[text1, text2, ...]").
		Action(func(cmd *cmdr.Command, remainArgs []string) (err error) {
			code := faker.Finance()
			typ := cmdr.GetStringRP(cmd.GetDottedNamePath(), "Type") // get 'Type' from ToggleGroup
			fmt.Printf("    %8s : %v\n", typ, code.CreditCard(typ))
			return
		})

	cmdr.NewBool(true).
		Titles("visa", "v").
		Description("generates a visa code", "").
		ToggleGroup("Type").
		Group("").
		EnvKeys("").
		AttachTo(cc)
	cmdr.NewBool().
		Titles("mastercard", "m").
		Description("generates a mastercard code", "").
		ToggleGroup("Type").
		AttachTo(cc)
	cmdr.NewBool().
		Titles("american_express", "a", "ae", "american-express").
		Description("generates a american_express code", "").
		ToggleGroup("Type").
		AttachTo(cc)
	cmdr.NewBool().
		Titles("diners_club", "d", "dc", "diners-club").
		Description("generates a diners_club code", "").
		ToggleGroup("Type").
		AttachTo(cc)
	cmdr.NewBool().
		Titles("discover", "di").
		Description("generates a discover code", "").
		ToggleGroup("Type").
		AttachTo(cc)
	cmdr.NewBool().
		Titles("maestro", "ma").
		Description("generates a maestro code", "").
		ToggleGroup("Type").
		AttachTo(cc)
	cmdr.NewBool().
		Titles("switch", "s").
		Description("generates a switch code", "").
		ToggleGroup("Type").
		AttachTo(cc)
	cmdr.NewBool().
		Titles("solo", "so").
		Description("generates a solo code", "").
		ToggleGroup("Type").
		AttachTo(cc)
	cmdr.NewBool().
		Titles("forbrugsforeningen", "f").
		Description("generates a forbrugsforeningen code", "").
		ToggleGroup("Type").
		AttachTo(cc)
	cmdr.NewBool().
		Titles("dankort", "dk").
		Description("generates a dankort code", "").
		ToggleGroup("Type").
		AttachTo(cc)
	cmdr.NewBool().
		Titles("laser", "l").
		Description("generates a laser code", "").
		ToggleGroup("Type").
		AttachTo(cc)

}
