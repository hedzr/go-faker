package logic

import (
	"fmt"
	"github.com/hedzr/cmdr"
	"syreclabs.com/go/faker"
)

func genCode(root *cmdr.RootCmdOpt) {

	codeCmd := root.NewSubCommand("code", "c", "codes").
		Description("generate codes (ISBN10, ISBN13, EAN13, EAN8, RUT, ABN)").
		Group("").
		TailPlaceholder("[text1, text2, ...]").
		Action(func(cmd *cmdr.Command, remainArgs []string) (err error) {
			code := faker.Code()

			typ := cmdr.GetStringRP(cmd.GetDottedNamePath(), "Type") // get 'Type' from ToggleGroup
			bdrs := map[string]struct {
				label string
				fn    func() string
			}{
				"isbn10": {"ISBN10", code.Isbn10},
				"isbn13": {"ISBN13", code.Isbn13},
				"ean13":  {"EAN13", code.Ean13},
				"ean8":   {"EAN8", code.Ean8},
				"rut":    {"RUT", code.Rut},
				"abn":    {"ABN", code.Abn},
				"":       {"Unknown", func() string { return "??" }},
			}

			fmt.Printf("    %8s : %v\n", bdrs[typ].label, bdrs[typ].fn())
			//fmt.Printf("    ISBN13        : %v\n", code.Isbn13())
			// fmt.Printf("    Type          : %v / %v / %v\n", cmdr.GetStringRP(cmd.GetDottedNamePath(), "Type"), cmd.GetDottedNamePath(), cmdr.GetStringR("code.Type"))
			return
		})

	cmdr.NewBool().
		Titles("isbn10", "10").
		Description("generates a ISBN10 code", "").
		ToggleGroup("Type").
		Group("").
		EnvKeys("").
		AttachTo(codeCmd)
	cmdr.NewBool().
		Titles("isbn13", "13").
		Description("generates a ISBN13 code", "").
		ToggleGroup("Type").
		AttachTo(codeCmd)
	cmdr.NewBool().
		Titles("ean13", "e13").
		Description("generates a EAN13 code", "").
		ToggleGroup("Type").
		AttachTo(codeCmd)
	cmdr.NewBool().
		Titles("ean8", "e8").
		Description("generates a EAN8 code", "").
		ToggleGroup("Type").
		AttachTo(codeCmd)
	cmdr.NewBool().
		Titles("rut", "r").
		Description("generates a RUT code", "").
		ToggleGroup("Type").
		AttachTo(codeCmd)
	cmdr.NewBool().
		Titles("abn", "a").
		Description("generates a ABN code", "").
		ToggleGroup("Type").
		AttachTo(codeCmd)

}
