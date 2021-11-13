package logic

import (
	"fmt"
	"github.com/hedzr/cmdr"
	"strings"
	"syreclabs.com/go/faker"
)

func genAddr(root *cmdr.RootCmdOpt) {
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
			oo := faker.Address()
			str := dumpAddr(oo)
			outputWithFormat(str, "Address")
			return
		})

}

func dumpAddr(oo faker.FakeAddress) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("  City                : %v\n", oo.City()))                // => "North Dessie"
	sb.WriteString(fmt.Sprintf("  StreetName          : %v\n", oo.StreetName()))          // => "Buckridge Lakes"
	sb.WriteString(fmt.Sprintf("  StreetAddress       : %v\n", oo.StreetAddress()))       // => "586 Sylvester Turnpike"
	sb.WriteString(fmt.Sprintf("  SecondaryAddress    : %v\n", oo.SecondaryAddress()))    // => "Apt. 411"
	sb.WriteString(fmt.Sprintf("  BuildingNumber      : %v\n", oo.BuildingNumber()))      // => "754"
	sb.WriteString(fmt.Sprintf("  Postcode            : %v\n", oo.Postcode()))            // => "31340"
	sb.WriteString(fmt.Sprintf("  PostcodeByState     : %v\n", oo.PostcodeByState("IN"))) // => "46511"
	sb.WriteString(fmt.Sprintf("  ZipCode             : %v\n", oo.ZipCode()))             // ZipCode is an alias for Postcode.
	sb.WriteString(fmt.Sprintf("  ZipCodeByState      : %v\n", oo.ZipCodeByState("IN")))  // ZipCodeByState is an alias for PostcodeByState.
	sb.WriteString(fmt.Sprintf("  TimeZone            : %v\n", oo.TimeZone()))            // => "Asia/Taipei"
	sb.WriteString(fmt.Sprintf("  CityPrefix          : %v\n", oo.CityPrefix()))          // => "East"
	sb.WriteString(fmt.Sprintf("  CitySuffix          : %v\n", oo.CitySuffix()))          // => "town"
	sb.WriteString(fmt.Sprintf("  StreetSuffix        : %v\n", oo.StreetSuffix()))        // => "Square"
	sb.WriteString(fmt.Sprintf("  State               : %v\n", oo.State()))               // => "Maryland"
	sb.WriteString(fmt.Sprintf("  StateAbbr           : %v\n", oo.StateAbbr()))           // => "IL"
	sb.WriteString(fmt.Sprintf("  Country             : %v\n", oo.Country()))             // => "Uruguay"
	sb.WriteString(fmt.Sprintf("  CountryCode         : %v\n", oo.CountryCode()))         // => "JP"
	sb.WriteString(fmt.Sprintf("  Latitude            : %v\n", oo.Latitude()))            // => (float32) -38.811367
	sb.WriteString(fmt.Sprintf("  Longitude           : %v\n", oo.Longitude()))           // => (float32) 89.2171
	sb.WriteString(fmt.Sprintf("  String              : %v\n", oo.String()))              // => "6071 Heaney Island Suite 553, Ebbaville Texas 37307"
	return sb.String()
}
