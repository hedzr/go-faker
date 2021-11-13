package logic

import (
	"fmt"
	"github.com/hedzr/cmdr"
	"syreclabs.com/go/faker"
)

func genCompany(root *cmdr.RootCmdOpt) {
	oo := faker.Company()
	m := map[string]struct {
		label   string
		fn      func() string
		short   string
		aliases []string
	}{
		"name":         {"Name", oo.Name, "n", []string{}},
		"suffix":       {"Suffix", oo.Suffix, "s", []string{}},
		"catch-phrase": {"CatchPhrase", oo.CatchPhrase, "cp", []string{}},
		"bs":           {"Bs", oo.Bs, "bs", []string{}},
		"ein":          {"Ein", oo.Ein, "ein", []string{}},
		"duns-number":  {"DunsNumber", oo.DunsNumber, "dn", []string{}},
		"logo":         {"Logo", oo.Logo, "l", []string{}},
		"string":       {"String", oo.String, "ss", []string{}},
	}

	//typ := cmdr.GetStringRP(cmd.GetDottedNamePath(), "Type")
	//sb.WriteString(fmt.Sprintf("    Name               : %v\n", oo.Name()))        // => "Aufderhar LLC"
	//sb.WriteString(fmt.Sprintf("    Suffix             : %v\n", oo.Suffix()))      // => "Inc"
	//sb.WriteString(fmt.Sprintf("    CatchPhrase        : %v\n", oo.CatchPhrase())) // => "Universal logistical artificial intelligence"
	//sb.WriteString(fmt.Sprintf("    Bs                 : %v\n", oo.Bs()))          // => "engage distributed applications"
	//sb.WriteString(fmt.Sprintf("    Ein                : %v\n", oo.Ein()))         // => "58-6520513"
	//sb.WriteString(fmt.Sprintf("    DunsNumber         : %v\n", oo.DunsNumber()))  // => "16-708-2968"
	//sb.WriteString(fmt.Sprintf("    Logo               : %v\n", oo.Logo()))        // => "http://www.biz-logo.com/examples/015.gif"
	//sb.WriteString(fmt.Sprintf("    String             : %v\n", oo.String()))      // String is an alias for Name.

	cc := root.NewSubCommand("company", "co", "corp").
		Description("generate Company names").
		Group("").
		TailPlaceholder("[text1, text2, ...]").
		Action(func(cmd *cmdr.Command, remainArgs []string) (err error) {
			str := dumpIt(cmd, m)
			outputWithFormat(str, "Company")
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
