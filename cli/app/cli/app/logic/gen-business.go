package logic

import (
	"fmt"
	"github.com/hedzr/cmdr"
	"strings"
	"syreclabs.com/go/faker"
)

func genBusiness(root *cmdr.RootCmdOpt) {
	root.NewSubCommand("business", "bz", "biz", "credit-card").
		Description("generate Business names").
		Group("").
		TailPlaceholder("[text1, text2, ...]").
		Action(func(cmd *cmdr.Command, remainArgs []string) (err error) {
			oo := faker.Business()
			str := dumpBusiness(oo)
			outputWithFormat(str, "Business")
			return
		})
}

func dumpBusiness(oo faker.FakeBusiness) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("    CreditCardNumber            : %v\n", oo.CreditCardNumber()))
	sb.WriteString(fmt.Sprintf("    CreditCardExpiryDate        : %v\n", oo.CreditCardExpiryDate()))
	sb.WriteString(fmt.Sprintf("    CreditCardType              : %v\n", oo.CreditCardType()))
	return sb.String()
}
