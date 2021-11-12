package cmdrrel

import (
	"github.com/%REPOSITORY%/cli/app"
	"github.com/%REPOSITORY%/internal"
	"fmt"
	"github.com/hedzr/cmdr"
	"github.com/hedzr/cmdr/conf"
	"io/ioutil"
	"net"
	"os"
	"time"
)

func onAppStart(cmd *cmdr.Command, args []string) (err error) {
	cmdr.Logger.Debugf("onAppStart")
	return
}

func onAppExit(cmd *cmdr.Command, args []string) {
	cmdr.Logger.Debugf("onAppExit")
}

func onServerPostStop(cmd *cmdr.Command, args []string) {
	cmdr.Logger.Debugf("onServerPostStop")
	internal.App().Close()
}

func host2ip(host string) {
	addr, err := net.LookupIP(host)
	if err != nil {
		fmt.Println("  ? Unknown host")
	} else {
		fmt.Printf("  - IP address of %q: %q\n", host, addr)
	}
}

// onServerPreStart is earlier than onAppStart.
func onServerPreStart(cmd *cmdr.Command, args []string) (err error) {
	fmt.Printf(`


********************************************************************
        SERVER STARTUP AT: %s
                 BUILT AT: %s
              APP VERSION: %s
               GIT COMMIT: %s

`,
		time.Now().UTC().Format(time.RFC1123Z), // or: RFC3339
		conf.Buildstamp, conf.Version, conf.Githash)

	var v []byte
	v, err = ioutil.ReadFile("/etc/hosts")
	fmt.Printf("/etc/hosts:\n%v\n\n", string(v))
	v, err = ioutil.ReadFile("/etc/resolv.conf")
	fmt.Printf("/etc/resolv.conf:\n%v\n\n", string(v))

	//fmt.Printf("Environment:\n\n")
	//_, _, err = exec.RunCommand("printenv", false)
	//// fmt.Printf("env (%v):\n%v\n\n", err, str)
	//fmt.Println()
	fmt.Printf("Environment:\n\n")
	for _, l := range os.Environ() {
		fmt.Printf("    %v\n", l)
	}
	//_, _, err = exec.RunCommand("printenv", false)
	//// fmt.Printf("env (%v):\n%v\n\n", err, str)
	fmt.Println()

	fmt.Println("IPs:")
	host2ip("db-master-host")
	host2ip("cache-master-host")
	host2ip("etcd-master-host")
	host2ip("registrar")
	host2ip("cache")
	host2ip("db")
	fmt.Println()

	cmdr.Logger.Debugf("onServerPreStart")

	cmdr.Set("app-module-name", app.AppModuleName)
	cmdr.Set("app-title", app.AppTitle)
	err = internal.App().Init(cmd, args)
	return
}
