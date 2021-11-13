module github.com/hedzr/go-faker

go 1.16

//replace github.com/hedzr/log => ../../go-cmdr/10.log

//replace github.com/hedzr/logex => ../../go-cmdr/15.logex

//replace github.com/hedzr/cmdr => ../../go-cmdr/50.cmdr

//replace github.com/hedzr/cmdr-addons => ../../go-cmdr/53.cmdr-addons

require (
	github.com/hedzr/cmdr v1.9.1
	github.com/hedzr/log v1.3.22
	github.com/hedzr/logex v1.3.22
	gopkg.in/hedzr/errors.v2 v2.1.5
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
	syreclabs.com/go/faker v1.2.3
)
