package logic

import (
	"fmt"
	"strings"

	"github.com/hedzr/cmdr"
	"syreclabs.com/go/faker"
)

func genCode(root *cmdr.RootCmdOpt) {

	cc := cmdr.NewSubCmd().Titles("code", "c", "codes").
		Description("generate Codes (ISBN10, ISBN13, EAN13, EAN8, RUT, ABN)").
		Group("").
		TailPlaceholder("[text1, text2, ...]").
		Action(func(cmd *cmdr.Command, remainArgs []string) (err error) {
			oo := faker.Code()
			str := dumpCode(oo, cmd)
			outputWithFormat(str, "Code")

			// fmt.Printf("    %8s : %v\n", bdrs[typ].label, bdrs[typ].fn())
			// fmt.Printf("    ISBN13        : %v\n", oo.Isbn13())
			// fmt.Printf("    Type          : %v / %v / %v\n", cmdr.GetStringRP(cmd.GetDottedNamePath(), "Type"), cmd.GetDottedNamePath(), cmdr.GetStringR("oo.Type"))
			return
		}).
		AttachTo(root)

	cmdr.NewBool().
		Titles("isbn10", "10").
		Description("generates a ISBN10 code", "").
		ToggleGroup("Type").
		Group("").
		EnvKeys("").
		AttachTo(cc)
	cmdr.NewBool().
		Titles("isbn13", "13").
		Description("generates a ISBN13 code", "").
		ToggleGroup("Type").
		AttachTo(cc)
	cmdr.NewBool().
		Titles("ean13", "e13").
		Description("generates a EAN13 code", "").
		ToggleGroup("Type").
		AttachTo(cc)
	cmdr.NewBool().
		Titles("ean8", "e8").
		Description("generates a EAN8 code", "").
		ToggleGroup("Type").
		AttachTo(cc)
	cmdr.NewBool().
		Titles("rut", "r").
		Description("generates a RUT code", "").
		ToggleGroup("Type").
		AttachTo(cc)
	cmdr.NewBool().
		Titles("abn", "a").
		Description("generates a ABN code", "").
		ToggleGroup("Type").
		AttachTo(cc)

}

func dumpCode(oo faker.FakeCode, cmd *cmdr.Command) string {
	var sb strings.Builder
	typ := cmdr.GetStringRP(cmd.GetDottedNamePath(), "Type") // get 'Type' from ToggleGroup
	bdrs := map[string]struct {
		label string
		fn    func() string
	}{
		"isbn10": {"ISBN10", oo.Isbn10},
		"isbn13": {"ISBN13", oo.Isbn13},
		"ean13":  {"EAN13", oo.Ean13},
		"ean8":   {"EAN8", oo.Ean8},
		"rut":    {"RUT", oo.Rut},
		"abn":    {"ABN", oo.Abn},
		"":       {"Unknown", func() string { return "??" }},
	}
	sb.WriteString(fmt.Sprintf("    %8s : %v\n", bdrs[typ].label, bdrs[typ].fn()))
	return sb.String()
}
