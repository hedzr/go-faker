module github.com/hedzr/go-faker

go 1.16

//replace github.com/hedzr/log => ../../go-cmdr/10.log

//replace github.com/hedzr/logex => ../../go-cmdr/15.logex

//replace github.com/hedzr/cmdr => ../../go-cmdr/50.cmdr

//replace github.com/hedzr/cmdr-addons => ../../go-cmdr/53.cmdr-addons

require (
	github.com/hedzr/cmdr v1.9.6
	github.com/hedzr/log v1.3.23
	github.com/hedzr/logex v1.3.23
	github.com/kr/text v0.2.0 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/crypto v0.0.0-20210513164829-c07d793c2f9a // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/hedzr/errors.v2 v2.1.5
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
	syreclabs.com/go/faker v1.2.3
)
