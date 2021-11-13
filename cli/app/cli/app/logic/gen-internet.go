package logic

import (
	"fmt"
	"github.com/hedzr/cmdr"
	"strings"
	"syreclabs.com/go/faker"
)

func genInternet(root *cmdr.RootCmdOpt) {
	root.NewSubCommand("internet", "i").
		Description("generate Internet names").
		Group("").
		TailPlaceholder("[text1, text2, ...]").
		Action(func(cmd *cmdr.Command, remainArgs []string) (err error) {
			oo := faker.Internet()
			str := dumpInternet(oo)
			outputWithFormat(str, "Internet")
			return
		})
}

func dumpInternet(oo faker.FakeInternet) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("    Email               : %v\n", oo.Email()))         // => "maritza@farrell.org"
	sb.WriteString(fmt.Sprintf("    FreeEmail           : %v\n", oo.FreeEmail()))     // => "sven_rice@hotmail.com"
	sb.WriteString(fmt.Sprintf("    SafeEmail           : %v\n", oo.SafeEmail()))     // => "theron.nikolaus@example.net"
	sb.WriteString(fmt.Sprintf("    UserName            : %v\n", oo.UserName()))      // => "micah_pfeffer"
	sb.WriteString(fmt.Sprintf("    Password            : %v\n", oo.Password(8, 14))) // => "s5CzvVp6Ye"
	sb.WriteString(fmt.Sprintf("    DomainName          : %v\n", oo.DomainName()))    // => "rolfson.info"
	sb.WriteString(fmt.Sprintf("    DomainWord          : %v\n", oo.DomainWord()))    // => "heller"
	sb.WriteString(fmt.Sprintf("    DomainSuffix        : %v\n", oo.DomainSuffix()))  // => "net"
	sb.WriteString(fmt.Sprintf("    MacAddress          : %v\n", oo.MacAddress()))    // => "15:a9:83:29:76:26"
	sb.WriteString(fmt.Sprintf("    IpV4Address         : %v\n", oo.IpV4Address()))   // => "121.204.82.227"
	sb.WriteString(fmt.Sprintf("    IpV6Address         : %v\n", oo.IpV6Address()))   // => "c697:392f:6a0e:bf6d:77e1:714a:10ab:0dbc"
	sb.WriteString(fmt.Sprintf("    Url                 : %v\n", oo.Url()))           // => "http://sporerhamill.net/kyla.schmitt"
	sb.WriteString(fmt.Sprintf("    Slug                : %v\n", oo.Slug()))          // => "officiis-commodi"
	return sb.String()
}
