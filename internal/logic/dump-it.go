package logic

import (
	"fmt"
	"sort"
	"strings"

	"github.com/hedzr/cmdr"
)

func dumpIt(cmd *cmdr.Command, m map[string]struct {
	label   string
	fn      func() string
	short   string
	aliases []string
}) string {
	var sb strings.Builder
	typ := cmdr.GetStringRP(cmd.GetDottedNamePath(), "Type")
	if typ != "" {
		sb.WriteString(fmt.Sprintf("    %20s : %v\n", m[typ].label, m[typ].fn()))
	} else {
		keys := make([]string, 0, len(m))
		for k := range m {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, k := range keys {
			sb.WriteString(fmt.Sprintf("    %20s : %v\n", m[k].label, m[k].fn()))
		}
	}
	return sb.String()
}
