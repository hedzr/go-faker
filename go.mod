module github.com/hedzr/go-faker

go 1.16

//replace github.com/hedzr/log => ../10.log

//replace github.com/hedzr/logex => ../15.logex

//replace github.com/hedzr/cmdr => ../50.cmdr

//replace github.com/hedzr/cmdr-addons => ../53.cmdr-addons

require (
	github.com/hedzr/cmdr v1.9.0
	github.com/hedzr/log v1.3.22
	github.com/hedzr/logex v1.3.22
	github.com/kr/text v0.2.0 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9 // indirect
	golang.org/x/sys v0.0.0-20210119212857-b64e53b001e4 // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/hedzr/errors.v2 v2.1.5
	syreclabs.com/go/faker v1.2.3
)
