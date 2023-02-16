package logic

import (
	"fmt"

	"github.com/hedzr/cmdr"
	"syreclabs.com/go/faker"
)

func genTeam(root *cmdr.RootCmdOpt) {
	oo := faker.Team()
	// faker.Team().Name()     // => "Colorado cats"
	// faker.Team().Creature() // => "cats"
	// faker.Team().State()    // => "Oregon"
	// faker.Team().String()   // String is an alias for Name.
	m := map[string]struct {
		label   string
		fn      func() string
		short   string
		aliases []string
	}{
		"name":     {"Name", oo.Name, "n", []string{}},
		"creature": {"Creature", oo.Creature, "c", []string{}},
		"state":    {"State", oo.State, "st", []string{}},
		"string":   {"String", oo.String, "ss", []string{}},
	}

	cc := cmdr.NewSubCmd().Titles("team", "t").
		Description("generate Team record").
		Group("").
		TailPlaceholder("[text1, text2, ...]").
		Action(func(cmd *cmdr.Command, remainArgs []string) (err error) {
			str := dumpIt(cmd, m)
			outputWithFormat(str, "PhoneNumber")
			return
		}).
		AttachTo(root)

	cmdr.NewInt(5).
		Titles("count", "C", "cnt").
		Group("").
		Placeholder("N").
		// Range(1,64).
		Description("number count for these fields: NumberXXX, Hexadecimal").
		AttachTo(cc)

	for k, v := range m {
		cmdr.NewBool().
			Titles(k, v.short, v.aliases...).
			Description(fmt.Sprintf("generates %v field", v.label), "").
			ToggleGroup("Type").
			AttachTo(cc)
	}
}
