package logic

import (
	"fmt"
	"github.com/hedzr/cmdr"
	"syreclabs.com/go/faker"
)

func genAddr(root *cmdr.RootCmdOpt) {
	oo := faker.Address()
	//sb.WriteString(fmt.Sprintf("  City                : %v\n", oo.City()))                // => "North Dessie"
	//sb.WriteString(fmt.Sprintf("  StreetName          : %v\n", oo.StreetName()))          // => "Buckridge Lakes"
	//sb.WriteString(fmt.Sprintf("  StreetAddress       : %v\n", oo.StreetAddress()))       // => "586 Sylvester Turnpike"
	//sb.WriteString(fmt.Sprintf("  SecondaryAddress    : %v\n", oo.SecondaryAddress()))    // => "Apt. 411"
	//sb.WriteString(fmt.Sprintf("  BuildingNumber      : %v\n", oo.BuildingNumber()))      // => "754"
	//sb.WriteString(fmt.Sprintf("  Postcode            : %v\n", oo.Postcode()))            // => "31340"
	//sb.WriteString(fmt.Sprintf("  PostcodeByState     : %v\n", oo.PostcodeByState("IN"))) // => "46511"
	//sb.WriteString(fmt.Sprintf("  ZipCode             : %v\n", oo.ZipCode()))             // ZipCode is an alias for Postcode.
	//sb.WriteString(fmt.Sprintf("  ZipCodeByState      : %v\n", oo.ZipCodeByState("IN")))  // ZipCodeByState is an alias for PostcodeByState.
	//sb.WriteString(fmt.Sprintf("  TimeZone            : %v\n", oo.TimeZone()))            // => "Asia/Taipei"
	//sb.WriteString(fmt.Sprintf("  CityPrefix          : %v\n", oo.CityPrefix()))          // => "East"
	//sb.WriteString(fmt.Sprintf("  CitySuffix          : %v\n", oo.CitySuffix()))          // => "town"
	//sb.WriteString(fmt.Sprintf("  StreetSuffix        : %v\n", oo.StreetSuffix()))        // => "Square"
	//sb.WriteString(fmt.Sprintf("  State               : %v\n", oo.State()))               // => "Maryland"
	//sb.WriteString(fmt.Sprintf("  StateAbbr           : %v\n", oo.StateAbbr()))           // => "IL"
	//sb.WriteString(fmt.Sprintf("  Country             : %v\n", oo.Country()))             // => "Uruguay"
	//sb.WriteString(fmt.Sprintf("  CountryCode         : %v\n", oo.CountryCode()))         // => "JP"
	//sb.WriteString(fmt.Sprintf("  Latitude            : %v\n", oo.Latitude()))            // => (float32) -38.811367
	//sb.WriteString(fmt.Sprintf("  Longitude           : %v\n", oo.Longitude()))           // => (float32) 89.2171
	//sb.WriteString(fmt.Sprintf("  String              : %v\n", oo.String()))              // => "6071 Heaney Island Suite 553, Ebbaville Texas 37307"
	m := map[string]struct {
		label   string
		fn      func() string
		short   string
		aliases []string
	}{
		"city":              {"City", oo.City, "c", []string{}},
		"street-name":       {"StreetName", oo.StreetName, "sn", []string{}},
		"street-address":    {"StreetAddress", oo.StreetAddress, "", []string{}},
		"secondary-address": {"SecondaryAddress", oo.SecondaryAddress, "", []string{}},
		"BuildingNumber":    {"BuildingNumber", oo.BuildingNumber, "", []string{}},
		"Postcode":          {"Postcode", oo.Postcode, "", []string{}},
		"PostcodeByState":   {"PostcodeByState", func() string { return fmt.Sprintf("%v", oo.PostcodeByState("IN")) }, "", []string{}},
		"ZipCode":           {"ZipCode", oo.ZipCode, "", []string{}},
		"ZipCodeByState":    {"ZipCodeByState", func() string { return fmt.Sprintf("%v", oo.ZipCodeByState("IN")) }, "", []string{}},
		"TimeZone":          {"TimeZone", oo.TimeZone, "", []string{}},
		"CityPrefix":        {"CityPrefix", oo.CityPrefix, "", []string{}},
		"CitySuffix":        {"CitySuffix", oo.CitySuffix, "", []string{}},
		"StreetSuffix":      {"StreetSuffix", oo.StreetSuffix, "", []string{}},
		"State":             {"State", oo.State, "", []string{}},
		"StateAbbr":         {"StateAbbr", oo.StateAbbr, "", []string{}},
		"Country":           {"Country", oo.Country, "", []string{}},
		"CountryCode":       {"CountryCode", oo.CountryCode, "", []string{}},
		"Latitude":          {"Latitude", func() string { return fmt.Sprintf("%v", oo.Latitude()) }, "", []string{}},
		"Longitude":         {"Longitude", func() string { return fmt.Sprintf("%v", oo.Longitude()) }, "", []string{}},
		"String":            {"String", oo.String, "", []string{}},
	}

	cc := root.NewSubCommand("addr", "a", "address").
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
			str := dumpIt(cmd, m)
			outputWithFormat(str, "Address")
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
