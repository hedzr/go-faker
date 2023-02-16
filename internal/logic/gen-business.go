package logic

import (
	"fmt"

	"github.com/hedzr/cmdr"
	"syreclabs.com/go/faker"
)

func genBusiness(root *cmdr.RootCmdOpt) {
	oo := faker.Business()
	m := map[string]struct {
		label   string
		fn      func() string
		short   string
		aliases []string
	}{
		"credit-card-number":      {"CreditCardNumber", oo.CreditCardNumber, "ccn", []string{}},
		"credit-card-expiry-date": {"CreditCardExpiryDate", oo.CreditCardExpiryDate, "cced", []string{}},
		"credit-card-type":        {"CreditCardType", oo.CreditCardType, "cct", []string{}},
	}
	// sb.WriteString(fmt.Sprintf("    CreditCardNumber            : %v\n", oo.CreditCardNumber()))
	// sb.WriteString(fmt.Sprintf("    CreditCardExpiryDate        : %v\n", oo.CreditCardExpiryDate()))
	// sb.WriteString(fmt.Sprintf("    CreditCardType              : %v\n", oo.CreditCardType()))

	cc := cmdr.NewSubCmd().Titles("business", "bz", "biz", "credit-card").
		Description("generate Business names").
		Group("").
		TailPlaceholder("[text1, text2, ...]").
		Action(func(cmd *cmdr.Command, remainArgs []string) (err error) {
			str := dumpIt(cmd, m)
			outputWithFormat(str, "Business")
			return
		}).
		AttachTo(root)
	for k, v := range m {
		cmdr.NewBool().
			Titles(k, v.short, v.aliases...).
			Description(fmt.Sprintf("generates %v field", v.label), "").
			ToggleGroup("Type").
			AttachTo(cc)
	}
}
