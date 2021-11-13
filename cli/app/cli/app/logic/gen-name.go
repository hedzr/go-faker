package logic

import (
	"fmt"
	"github.com/hedzr/cmdr"
	"syreclabs.com/go/faker"
)

func genName(root *cmdr.RootCmdOpt) {
	oo := faker.Name()
	m := map[string]struct {
		label   string
		fn      func() string
		short   string
		aliases []string
	}{
		"name":       {"Name", oo.Name, "n", []string{}},
		"first-name": {"FirstName", oo.FirstName, "f", []string{}},
		"last-name":  {"LastName", oo.LastName, "l", []string{}},
		"prefix":     {"Prefix", oo.Prefix, "p", []string{}},
		"suffix":     {"Suffix", oo.Suffix, "s", []string{}},
		"title":      {"Title", oo.Title, "t", []string{}},
		"string":     {"String", oo.String, "ss", []string{}},
	}

	//faker.Name().Name()      // => "Natasha Hartmann"
	//faker.Name().FirstName() // => "Carolina"
	//faker.Name().LastName()  // => "Kohler"
	//faker.Name().Prefix()    // => "Dr."
	//faker.Name().Suffix()    // => "Jr."
	//faker.Name().Title()     // => "Chief Functionality Orchestrator"
	//faker.Name().String()    // String is an alias for Name.

	cc := root.NewSubCommand("name", "n").
		Description("generate Name").
		Group("").
		TailPlaceholder("[text1, text2, ...]").
		Action(func(cmd *cmdr.Command, remainArgs []string) (err error) {
			str := dumpIt(cmd, m)
			outputWithFormat(str, "Name")
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
