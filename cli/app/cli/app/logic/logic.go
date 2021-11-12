// Copyright © 2021 Hedzr Yeh.
// All Rights Reserved.
// These codes and documentations are reserved for
// non-commercial and private purpose.

package logic

import (
	"fmt"
	"github.com/hedzr/cmdr"
	"math"
	"syreclabs.com/go/faker"
)

// see also: https://github.com/danistefanovic/build-your-own-x#build-your-own-command-line-tool
// see also: https://flaviocopes.com/go-tutorial-lolcat/
// see also: https://github.com/dmgk/faker

func AttachToCmdr(root *cmdr.RootCmdOpt) {

	root.NewSubCommand("addr", "a", "address").
		Description("generate address").
		Group("").
		TailPlaceholder("[text1, text2, ...]").
		//PreAction(func(cmd *cmdr.Command, remainArgs []string) (err error) {
		//	fmt.Printf("[PRE] DebugMode=%v, TraceMode=%v. InDebugging/IsDebuggerAttached=%v\n",
		//		cmdr.GetDebugMode(), logex.GetTraceMode(), cmdr.InDebugging())
		//	for ix, s := range remainArgs {
		//		fmt.Printf("[PRE] %5d. %s\n", ix, s)
		//	}
		//
		//	fmt.Printf("[PRE] Debug=%v, Trace=%v\n", cmdr.GetDebugMode(), cmdr.GetTraceMode())
		//
		//	// return nil to be continue,
		//	// return cmdr.ErrShouldBeStopException to stop the following actions without error
		//	// return other errors for application purpose
		//	return
		//}).
		//PostAction(func(cmd *cmdr.Command, remainArgs []string) {
		//	for ix, s := range remainArgs {
		//		fmt.Printf("[POST] %5d. %s\n", ix, s)
		//	}
		//}).
		Action(func(cmd *cmdr.Command, remainArgs []string) (err error) {
			addr := faker.Address()
			fmt.Printf("    City                : %v\n", addr.City())                // => "North Dessie"
			fmt.Printf("    StreetName          : %v\n", addr.StreetName())          // => "Buckridge Lakes"
			fmt.Printf("    StreetAddress       : %v\n", addr.StreetAddress())       // => "586 Sylvester Turnpike"
			fmt.Printf("    SecondaryAddress    : %v\n", addr.SecondaryAddress())    // => "Apt. 411"
			fmt.Printf("    BuildingNumber      : %v\n", addr.BuildingNumber())      // => "754"
			fmt.Printf("    BuildingNumber      : %v\n", addr.BuildingNumber())      // => "31340"
			fmt.Printf("    PostcodeByState     : %v\n", addr.PostcodeByState("IN")) // => "46511"
			fmt.Printf("    ZipCode             : %v\n", addr.ZipCode())             // ZipCode is an alias for Postcode.
			fmt.Printf("    ZipCodeByState      : %v\n", addr.ZipCodeByState("IN"))  // ZipCodeByState is an alias for PostcodeByState.
			fmt.Printf("    TimeZone            : %v\n", addr.TimeZone())            // => "Asia/Taipei"
			fmt.Printf("    CityPrefix          : %v\n", addr.CityPrefix())          // => "East"
			fmt.Printf("    CitySuffix          : %v\n", addr.CitySuffix())          // => "town"
			fmt.Printf("    StreetSuffix        : %v\n", addr.StreetSuffix())        // => "Square"
			fmt.Printf("    State               : %v\n", addr.State())               // => "Maryland"
			fmt.Printf("    StateAbbr           : %v\n", addr.StateAbbr())           // => "IL"
			fmt.Printf("    Country             : %v\n", addr.Country())             // => "Uruguay"
			fmt.Printf("    CountryCode         : %v\n", addr.CountryCode())         // => "JP"
			fmt.Printf("    Latitude            : %v\n", addr.Latitude())            // => (float32) -38.811367
			fmt.Printf("    Longitude           : %v\n", addr.Longitude())           // => (float32) 89.2171
			fmt.Printf("    String              : %v\n", addr.String())              // => "6071 Heaney Island Suite 553, Ebbaville Texas 37307"
			return
		})

	root.NewSubCommand("app", "app").
		Description("generate app names").
		Group("").
		TailPlaceholder("[text1, text2, ...]").
		Action(func(cmd *cmdr.Command, remainArgs []string) (err error) {
			addr := faker.App()
			fmt.Printf("    Name                : %v\n", addr.Name())    // => "Alphazap"
			fmt.Printf("    Version             : %v\n", addr.Version()) // => "2.6.0"
			fmt.Printf("    Author              : %v\n", addr.Author())  // => "Dorian Shields"
			fmt.Printf("    String              : %v\n", addr.String())  // => "Tempsoft 4.51"
			return
		})

	root.NewSubCommand("avatar", "avatar").
		Description("generate Avatar names").
		Group("").
		TailPlaceholder("[text1, text2, ...]").
		Action(func(cmd *cmdr.Command, remainArgs []string) (err error) {
			addr := faker.Avatar()
			fmt.Printf("    Jpeg                : %v\n", addr.Url("jpg", 100, 200)) // => "Alphazap"
			fmt.Printf("    String              : %v\n", addr.String())             // => "Tempsoft 4.51"
			return
		})

	root.NewSubCommand("bitcoin", "b", "btc").
		Description("generate Bitcoin names").
		Group("").
		TailPlaceholder("[text1, text2, ...]").
		Action(func(cmd *cmdr.Command, remainArgs []string) (err error) {
			addr := faker.Bitcoin()
			fmt.Printf("    Address             : %v\n", addr.Address()) // => "Alphazap"
			fmt.Printf("    String              : %v\n", addr.String())  // => "Tempsoft 4.51"
			return
		})

	root.NewSubCommand("business", "bz", "credit-card").
		Description("generate Bitcoin names").
		Group("").
		TailPlaceholder("[text1, text2, ...]").
		Action(func(cmd *cmdr.Command, remainArgs []string) (err error) {
			addr := faker.Business()
			fmt.Printf("    CreditCardNumber            : %v\n", addr.CreditCardNumber())
			fmt.Printf("    CreditCardExpiryDate        : %v\n", addr.CreditCardExpiryDate())
			fmt.Printf("    CreditCardType              : %v\n", addr.CreditCardType())
			return
		})

	genCode(root)

	root.NewSubCommand("commerce", "cc").
		Description("generate Commerce names").
		Group("").
		TailPlaceholder("[text1, text2, ...]").
		Action(func(cmd *cmdr.Command, remainArgs []string) (err error) {
			addr := faker.Commerce()
			fmt.Printf("    ProductName        : %v\n", addr.ProductName())
			fmt.Printf("    Department         : %v\n", addr.Department())
			fmt.Printf("    Price              : %v\n", addr.Price())
			fmt.Printf("    Color              : %v\n", addr.Color())
			return
		})

	root.NewSubCommand("company", "corp").
		Description("generate Company names").
		Group("").
		TailPlaceholder("[text1, text2, ...]").
		Action(func(cmd *cmdr.Command, remainArgs []string) (err error) {
			addr := faker.Company()
			fmt.Printf("    Name               : %v\n", addr.Name())        // => "Aufderhar LLC"
			fmt.Printf("    Suffix             : %v\n", addr.Suffix())      // => "Inc"
			fmt.Printf("    CatchPhrase        : %v\n", addr.CatchPhrase()) // => "Universal logistical artificial intelligence"
			fmt.Printf("    Bs                 : %v\n", addr.Bs())          // => "engage distributed applications"
			fmt.Printf("    Ein                : %v\n", addr.Ein())         // => "58-6520513"
			fmt.Printf("    DunsNumber         : %v\n", addr.DunsNumber())  // => "16-708-2968"
			fmt.Printf("    Logo               : %v\n", addr.Logo())        // => "http://www.biz-logo.com/examples/015.gif"
			fmt.Printf("    String             : %v\n", addr.String())      // String is an alias for Name.
			return
		})

	genFinance(root)
	genHackers(root)

	root.NewSubCommand("internet", "i").
		Description("generate Internet names").
		Group("").
		TailPlaceholder("[text1, text2, ...]").
		Action(func(cmd *cmdr.Command, remainArgs []string) (err error) {
			addr := faker.Internet()

			fmt.Printf("    Email               : %v\n", addr.Email())         // => "maritza@farrell.org"
			fmt.Printf("    FreeEmail           : %v\n", addr.FreeEmail())     // => "sven_rice@hotmail.com"
			fmt.Printf("    SafeEmail           : %v\n", addr.SafeEmail())     // => "theron.nikolaus@example.net"
			fmt.Printf("    UserName            : %v\n", addr.UserName())      // => "micah_pfeffer"
			fmt.Printf("    Password            : %v\n", addr.Password(8, 14)) // => "s5CzvVp6Ye"
			fmt.Printf("    DomainName          : %v\n", addr.DomainName())    // => "rolfson.info"
			fmt.Printf("    DomainWord          : %v\n", addr.DomainWord())    // => "heller"
			fmt.Printf("    DomainSuffix        : %v\n", addr.DomainSuffix())  // => "net"
			fmt.Printf("    MacAddress          : %v\n", addr.MacAddress())    // => "15:a9:83:29:76:26"
			fmt.Printf("    IpV4Address         : %v\n", addr.IpV4Address())   // => "121.204.82.227"
			fmt.Printf("    IpV6Address         : %v\n", addr.IpV6Address())   // => "c697:392f:6a0e:bf6d:77e1:714a:10ab:0dbc"
			fmt.Printf("    Url                 : %v\n", addr.Url())           // => "http://sporerhamill.net/kyla.schmitt"
			fmt.Printf("    Slug                : %v\n", addr.Slug())          // => "officiis-commodi"
			return
		})

	// Date
	// Lorem
	// Name
	// Number
	// PhoneNumber
	// Team
	// Time

	// https://github.com/dmgk/faker
}

func rgb(i int) (int, int, int) {
	var f = 0.1
	return int(math.Sin(f*float64(i)+0)*127 + 128),
		int(math.Sin(f*float64(i)+2*math.Pi/3)*127 + 128),
		int(math.Sin(f*float64(i)+4*math.Pi/3)*127 + 128)
}

func printZ(output []rune) {
	for j := 0; j < len(output); j++ {
		r, g, b := rgb(j)
		fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m", r, g, b, output[j])
	}
	fmt.Println()
}
