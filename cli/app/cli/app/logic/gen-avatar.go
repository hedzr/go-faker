package logic

import (
	"fmt"
	"github.com/hedzr/cmdr"
	"strings"
	"syreclabs.com/go/faker"
)

func genAvatar(root *cmdr.RootCmdOpt) {
	root.NewSubCommand("avatar", "avatar").
		Description("generate Avatar names").
		Group("").
		TailPlaceholder("[text1, text2, ...]").
		Action(func(cmd *cmdr.Command, remainArgs []string) (err error) {
			oo := faker.Avatar()
			str := dumpAvatar(oo)
			outputWithFormat(str, "Avatar")
			return
		})

}

func dumpAvatar(oo faker.FakeAvatar) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("    Jpeg                : %v\n", oo.Url("jpg", 100, 200))) // => "Alphazap"
	sb.WriteString(fmt.Sprintf("    String              : %v\n", oo.String()))             // => "Tempsoft 4.51"
	return sb.String()
}
