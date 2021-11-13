package logic

import (
	"encoding/json"
	"fmt"
	"github.com/hedzr/cmdr"
	"github.com/hedzr/log"
	"gopkg.in/yaml.v3"
	"regexp"
	"strings"
)

// egrepReplace compile pattern with POSIX ERE (egrep) syntax
func egrepReplace(source, pattern, repl string) (result string) {
	r := regexp.MustCompilePOSIX(pattern)
	result = r.ReplaceAllString(source, repl)
	return
}

func fromYAMLString(s string) (target interface{}) {
	s2 := egrepReplace(s, `^([ \t]+)([a-zA-Z0-9_-]+) +: +(.+?)$`, "  $2: $3")
	err := yaml.Unmarshal([]byte(s2), &target)
	if err != nil {
		log.Errorf("parse yaml string failed: %v", err)
		return
	}
	return
}

func toYAML(target interface{}) (b []byte) {
	var err error
	b, err = yaml.Marshal(target)
	if err != nil {
		log.Errorf("generate yaml string failed: %v", err)
		return
	}
	return
}

func toJSON(target interface{}) (b []byte) {
	var err error
	b, err = json.Marshal(target)
	if err != nil {
		log.Errorf("generate yaml string failed: %v", err)
		return
	}
	return
}

func toJSONPretty(target interface{}) (b []byte) {
	var err error
	b, err = json.MarshalIndent(target, "", "  ")
	if err != nil {
		log.Errorf("generate yaml string failed: %v", err)
		return
	}
	return
}

func outputWithFormat(str, header string) {
	format := cmdr.GetStringR("faker.Format", "plain")
	switch format {
	case "yaml", "json":
		var sb strings.Builder
		sb.WriteString(header + ":\n")
		sb.WriteString(str)
		s := sb.String()
		//fmt.Println(s)
		obj := fromYAMLString(s)
		switch format {
		case "yaml":
			fmt.Println(string(toYAML(obj)))
		case "json":
			fmt.Println(string(toJSONPretty(obj)))
		case "json-compact":
			fmt.Println(string(toJSON(obj)))
		}
	default:
		fmt.Println(str)
	}
}
