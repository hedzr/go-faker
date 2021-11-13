package logic

import (
	"fmt"
	"github.com/hedzr/cmdr"
	"syreclabs.com/go/faker"
)

func genAvatar(root *cmdr.RootCmdOpt) {
	oo := faker.Avatar()
	//sb.WriteString(fmt.Sprintf("    Jpeg                : %v\n", oo.Url("jpg", 100, 200))) // => "Alphazap"
	//sb.WriteString(fmt.Sprintf("    String              : %v\n", oo.String()))             // => "Tempsoft 4.51"
	m := map[string]struct {
		label   string
		fn      func() string
		short   string
		aliases []string
	}{
		"jpeg": {"JPEG", func() string {
			return oo.Url("jpg", cmdr.GetIntR("avatar.width"), cmdr.GetIntR("avatar.height"))
		}, "jpg", []string{}},
		"png": {"PNG", func() string {
			return oo.Url("png", cmdr.GetIntR("avatar.width"), cmdr.GetIntR("avatar.height"))
		}, "png", []string{}},
		"string": {"String", oo.String, "ss", []string{}},
	}

	cc := root.NewSubCommand("avatar", "av").
		Description("generate Avatar names").
		Group("").
		TailPlaceholder("[text1, text2, ...]").
		Action(func(cmd *cmdr.Command, remainArgs []string) (err error) {
			str := dumpIt(cmd, m)
			outputWithFormat(str, "Avatar")
			return
		})

	//cmdr.NewBool().
	//	Titles("jpeg", "j", "jpg").
	//	Description("generates JPEG field", "").
	//	ToggleGroup("Type").
	//	AttachTo(cc)
	//cmdr.NewBool().
	//	Titles("png", "p").
	//	Description("generates PNG field", "").
	//	ToggleGroup("Type").
	//	AttachTo(cc)
	//cmdr.NewBool().
	//	Titles("string", "ss").
	//	Description("generates String (PNG) field", "").
	//	ToggleGroup("Type").
	//	AttachTo(cc)

	for k, v := range m {
		cmdr.NewBool().
			Titles(k, v.short, v.aliases...).
			Description(fmt.Sprintf("generates %v field", v.label), "").
			ToggleGroup("Type").
			AttachTo(cc)
	}

	cmdr.NewInt(320).
		Titles("width", "W").
		Group("").
		Placeholder("Num").
		//Range(1,64).
		Description("image width").
		AttachTo(cc)
	cmdr.NewInt(240).
		Titles("height", "H").
		Group("").
		Placeholder("Num").
		//Range(1,64).
		Description("image height").
		AttachTo(cc)
}

//func dumpAvatar(oo faker.FakeAvatar, cmd *cmdr.Command) string {
//	var sb strings.Builder
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
//}
