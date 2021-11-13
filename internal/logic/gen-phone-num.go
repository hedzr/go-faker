package logic

import (
	"fmt"
	"github.com/hedzr/cmdr"
	"syreclabs.com/go/faker"
)

func genPhoneNumber(root *cmdr.RootCmdOpt) {
	oo := faker.PhoneNumber()
	//faker.PhoneNumber().PhoneNumber()       // => "1-599-267-6597 x537"
	//faker.PhoneNumber().CellPhone()         // => "+49-131-0003060"
	//faker.PhoneNumber().AreaCode()          // => "903"
	//faker.PhoneNumber().ExchangeCode()      // => "574"
	//faker.PhoneNumber().SubscriberNumber(4) // => "1512"
	//faker.PhoneNumber().String()            // String is an alias for PhoneNumber.
	m := map[string]struct {
		label   string
		fn      func() string
		short   string
		aliases []string
	}{
		"phone-number":      {"PhoneNumber", oo.PhoneNumber, "pn", []string{}},
		"call-phone":        {"CellPhone", oo.CellPhone, "cp", []string{}},
		"area-code":         {"AreaCode", oo.AreaCode, "an", []string{}},
		"exchange-code":     {"ExchangeCode", oo.ExchangeCode, "ec", []string{}},
		"string":            {"String", oo.String, "ss", []string{}},
		"subscriber-number": {"SubscriberNumber", func() string { return fmt.Sprintf("%v", oo.SubscriberNumber(cmdr.GetIntR("phone-number.count"))) }, "sn", []string{}},
	}

	cc := root.NewSubCommand("phone-number", "pn").
		Description("generate PhoneNumber record").
		Group("").
		TailPlaceholder("[text1, text2, ...]").
		Action(func(cmd *cmdr.Command, remainArgs []string) (err error) {
			str := dumpIt(cmd, m)
			outputWithFormat(str, "PhoneNumber")
			return
		})

	cmdr.NewInt(5).
		Titles("count", "C", "cnt").
		Group("").
		Placeholder("N").
		//Range(1,64).
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
