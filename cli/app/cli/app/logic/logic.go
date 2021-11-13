// Copyright © 2021 Hedzr Yeh.
// All Rights Reserved.
// These codes and documentations are reserved for
// non-commercial and private purpose.

package logic

import (
	"github.com/hedzr/cmdr"
)

// see also: https://github.com/danistefanovic/build-your-own-x#build-your-own-command-line-tool
// see also: https://flaviocopes.com/go-tutorial-lolcat/
// see also: https://github.com/dmgk/faker

func AttachToCmdr(root *cmdr.RootCmdOpt) {

	cmdr.NewBool().
		Titles("plain", "plain").
		Description("output with plain format", "").
		ToggleGroup("Format").
		Group("").
		EnvKeys("").
		AttachTo(root)
	cmdr.NewBool(true).
		Titles("json-compact", "json-compact").
		Description("output with json format", "").
		ToggleGroup("Format").
		Group("").
		EnvKeys("").
		AttachTo(root)
	cmdr.NewBool().
		Titles("json", "json").
		Description("output with json pretty format", "").
		ToggleGroup("Format").
		Group("").
		EnvKeys("").
		AttachTo(root)
	cmdr.NewBool().
		Titles("yaml", "yaml").
		Description("output with yaml format", "").
		ToggleGroup("Format").
		Group("").
		EnvKeys("").
		AttachTo(root)

	//

	genAddr(root)
	genApp(root)
	genAvatar(root)
	genBitcoin(root)
	genBusiness(root)
	genCode(root)
	genCommerce(root)
	genCompany(root)
	genFinance(root)
	genHackers(root)
	genInternet(root)

	// Date
	// Lorem
	// Name
	// Number
	// PhoneNumber
	// Team
	// Time

	// https://github.com/dmgk/faker
}
