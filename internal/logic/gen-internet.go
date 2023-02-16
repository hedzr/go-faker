package logic

import (
	"fmt"

	"github.com/hedzr/cmdr"
	"syreclabs.com/go/faker"
)

func genInternet(root *cmdr.RootCmdOpt) {
	oo := faker.Internet()
	m := map[string]struct {
		label   string
		fn      func() string
		short   string
		aliases []string
	}{
		"email":         {"Email", oo.Email, "em", []string{}},
		"free-email":    {"FreeEmail", oo.FreeEmail, "fe", []string{}},
		"safe-email":    {"SafeEmail", oo.SafeEmail, "se", []string{}},
		"user-name":     {"UserName", oo.UserName, "un", []string{}},
		"password":      {"Password", func() string { return oo.Password(cmdr.GetIntR("internet.min"), cmdr.GetIntR("internet.max")) }, "p", []string{"pwd", "pass"}},
		"domain-name":   {"DomainName", oo.DomainName, "dn", []string{}},
		"domain-word":   {"DomainWord", oo.DomainWord, "dw", []string{}},
		"domain-suffix": {"DomainSuffix", oo.DomainSuffix, "ds", []string{}},
		"mac-address":   {"MacAddress", oo.MacAddress, "m", []string{"mac"}},
		"ip-v4-address": {"IpV4Address", oo.IpV4Address, "ip", []string{}},
		"ip-v6-address": {"IpV6Address", oo.IpV6Address, "ip6", []string{}},
		"url":           {"Url", oo.Url, "u", []string{}},
		"slug":          {"Slug", oo.Slug, "ss", []string{}},
	}
	// sb.WriteString(fmt.Sprintf("    Email               : %v\n", oo.Email()))         // => "maritza@farrell.org"
	// sb.WriteString(fmt.Sprintf("    FreeEmail           : %v\n", oo.FreeEmail()))     // => "sven_rice@hotmail.com"
	// sb.WriteString(fmt.Sprintf("    SafeEmail           : %v\n", oo.SafeEmail()))     // => "theron.nikolaus@example.net"
	// sb.WriteString(fmt.Sprintf("    UserName            : %v\n", oo.UserName()))      // => "micah_pfeffer"
	// sb.WriteString(fmt.Sprintf("    Password            : %v\n", oo.Password(8, 14))) // => "s5CzvVp6Ye"
	// sb.WriteString(fmt.Sprintf("    DomainName          : %v\n", oo.DomainName()))    // => "rolfson.info"
	// sb.WriteString(fmt.Sprintf("    DomainWord          : %v\n", oo.DomainWord()))    // => "heller"
	// sb.WriteString(fmt.Sprintf("    DomainSuffix        : %v\n", oo.DomainSuffix()))  // => "net"
	// sb.WriteString(fmt.Sprintf("    MacAddress          : %v\n", oo.MacAddress()))    // => "15:a9:83:29:76:26"
	// sb.WriteString(fmt.Sprintf("    IpV4Address         : %v\n", oo.IpV4Address()))   // => "121.204.82.227"
	// sb.WriteString(fmt.Sprintf("    IpV6Address         : %v\n", oo.IpV6Address()))   // => "c697:392f:6a0e:bf6d:77e1:714a:10ab:0dbc"
	// sb.WriteString(fmt.Sprintf("    Url                 : %v\n", oo.Url()))           // => "http://sporerhamill.net/kyla.schmitt"
	// sb.WriteString(fmt.Sprintf("    Slug                : %v\n", oo.Slug()))          // => "officiis-commodi"

	cc := cmdr.NewSubCmd().Titles("internet", "i").
		Description("generate Internet names").
		Group("").
		TailPlaceholder("[text1, text2, ...]").
		Action(func(cmd *cmdr.Command, remainArgs []string) (err error) {
			str := dumpIt(cmd, m)
			outputWithFormat(str, "Internet")
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
	cmdr.NewInt(8).
		Titles("min", "min").
		Group("").
		Placeholder("Num").
		// Range(1,64).
		Description("minimal length of Password field").
		AttachTo(cc)
	cmdr.NewInt(32).
		Titles("max", "max").
		Group("").
		Placeholder("Num").
		// Range(1,64).
		Description("maximal length of Password field").
		AttachTo(cc)
}
