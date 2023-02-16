package logic

import (
	"fmt"

	"github.com/hedzr/cmdr"
	"syreclabs.com/go/faker"
)

func genApp(root *cmdr.RootCmdOpt) {
	oo := faker.App()

	m := map[string]struct {
		label   string
		fn      func() string
		short   string
		aliases []string
	}{
		"name":    {"Name", oo.Name, "n", []string{}},
		"version": {"Version", oo.Version, "vv", []string{"ver"}},
		"author":  {"Author", oo.Author, "a", []string{}},
		"string":  {"String", oo.String, "ss", []string{}},
	}
	// sb.WriteString(fmt.Sprintf("    Name                : %v\n", oo.Name()))    // => "Alphazap"
	// sb.WriteString(fmt.Sprintf("    Version             : %v\n", oo.Version())) // => "2.6.0"
	// sb.WriteString(fmt.Sprintf("    Author              : %v\n", oo.Author()))  // => "Dorian Shields"
	// sb.WriteString(fmt.Sprintf("    String              : %v\n", oo.String()))  // => "Tempsoft 4.51"

	cc := cmdr.NewSubCmd().Titles("app", "app").
		Description("generate app names").
		Group("").
		TailPlaceholder("[text1, text2, ...]").
		Action(func(cmd *cmdr.Command, remainArgs []string) (err error) {
			str := dumpIt(cmd, m)
			outputWithFormat(str, "App")
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

	// cmdr.NewBool().
	//	Titles("name", "n").
	//	Description("generates Name field", "").
	//	ToggleGroup("Type").
	//	AttachTo(cc)
	// cmdr.NewBool().
	//	Titles("version", "v", "ver").
	//	Description("generates Version field", "").
	//	ToggleGroup("Type").
	//	AttachTo(cc)
	// cmdr.NewBool().
	//	Titles("Author", "a").
	//	Description("generates Author field", "").
	//	ToggleGroup("Type").
	//	AttachTo(cc)
	// cmdr.NewBool().
	//	Titles("string", "ss").
	//	Description("generates String field", "").
	//	ToggleGroup("Type").
	//	AttachTo(cc)

}

// func dumpApp(oo faker.FakeApp, cmd *cmdr.Command) string {
//	var sb strings.Builder
//	m := map[string]struct {
//		label string
//		fn    func() string
//	}{
//		"name":    {"Name", oo.Name},
//		"version": {"Version", oo.Version},
//		"author":  {"Author", oo.Author},
//		"string":  {"String", oo.String},
//	}
//	//sb.WriteString(fmt.Sprintf("    Name                : %v\n", oo.Name()))    // => "Alphazap"
//	//sb.WriteString(fmt.Sprintf("    Version             : %v\n", oo.Version())) // => "2.6.0"
//	//sb.WriteString(fmt.Sprintf("    Author              : %v\n", oo.Author()))  // => "Dorian Shields"
//	//sb.WriteString(fmt.Sprintf("    String              : %v\n", oo.String()))  // => "Tempsoft 4.51"
//	typ := cmdr.GetStringRP(cmd.GetDottedNamePath(), "Type")
//	if typ != "" {
//		sb.WriteString(fmt.Sprintf("    %20s : %v\n", m[typ].label, m[typ].fn()))
//	} else {
//		keys := make([]string, 0, len(m))
//		for k := range m {
//			keys = append(keys, k)
//		}
//		sort.Strings(keys)
//
//		for _, k := range keys {
//			sb.WriteString(fmt.Sprintf("    %20s : %v\n", m[k].label, m[k].fn()))
//		}
//	}
//	return sb.String()
// }
