package logic

import (
	"fmt"
	"github.com/hedzr/cmdr"
	"strings"
	"syreclabs.com/go/faker"
)

func genCompany(root *cmdr.RootCmdOpt) {
	root.NewSubCommand("company", "co", "corp").
		Description("generate Company names").
		Group("").
		TailPlaceholder("[text1, text2, ...]").
		Action(func(cmd *cmdr.Command, remainArgs []string) (err error) {
			oo := faker.Company()
			str := dumpCompany(oo)
			outputWithFormat(str, "Company")
			return
		})
}

func dumpCompany(oo faker.FakeCompany) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("    Name               : %v\n", oo.Name()))        // => "Aufderhar LLC"
	sb.WriteString(fmt.Sprintf("    Suffix             : %v\n", oo.Suffix()))      // => "Inc"
	sb.WriteString(fmt.Sprintf("    CatchPhrase        : %v\n", oo.CatchPhrase())) // => "Universal logistical artificial intelligence"
	sb.WriteString(fmt.Sprintf("    Bs                 : %v\n", oo.Bs()))          // => "engage distributed applications"
	sb.WriteString(fmt.Sprintf("    Ein                : %v\n", oo.Ein()))         // => "58-6520513"
	sb.WriteString(fmt.Sprintf("    DunsNumber         : %v\n", oo.DunsNumber()))  // => "16-708-2968"
	sb.WriteString(fmt.Sprintf("    Logo               : %v\n", oo.Logo()))        // => "http://www.biz-logo.com/examples/015.gif"
	sb.WriteString(fmt.Sprintf("    String             : %v\n", oo.String()))      // String is an alias for Name.
	return sb.String()
}
