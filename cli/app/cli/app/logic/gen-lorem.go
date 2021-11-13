package logic

import (
	"fmt"
	"github.com/hedzr/cmdr"
	"syreclabs.com/go/faker"
)

func genLorem(root *cmdr.RootCmdOpt) {
	oo := faker.Lorem()
	m := map[string]struct {
		label   string
		fn      func() string
		short   string
		aliases []string
	}{
		"character": {"Character", oo.Character, "c", []string{"char"}},
		"characters": {"Characters", func() string {
			length := cmdr.GetIntRP("lorem", "length")
			return oo.Characters(length)
		}, "cc", []string{"chars"}},
		"word": {"word", oo.Word, "w", []string{}},
		"words": {"Words", func() string {
			count := cmdr.GetIntRP("lorem", "count")
			return fmt.Sprintf("%v", oo.Words(count))
		}, "ww", []string{}},
		"sentence": {"Sentence", func() string {
			count := cmdr.GetIntRP("lorem", "count")
			return fmt.Sprintf("%v", oo.Sentence(count))
		}, "s", []string{}},
		"sentences": {"Sentences", func() string {
			count := cmdr.GetIntRP("lorem", "count")
			return fmt.Sprintf("%v", oo.Sentences(count))
		}, "ss", []string{}},
	}

	//faker.Lorem().Character()    // => "c"
	//faker.Lorem().Characters(17) // => "wqFyJIrXYfVP7cL9M"
	//faker.Lorem().Word()         // => "veritatis"
	//faker.Lorem().Words(3)       // => []string{"omnis", "libero", "neque"}
	//faker.Lorem().Sentence(3)    // => "Necessitatibus sit autem."

	cc := root.NewSubCommand("lorem", "l").
		Description("generate Lorem text").
		Group("").
		TailPlaceholder("[text1, text2, ...]").
		Action(func(cmd *cmdr.Command, remainArgs []string) (err error) {
			str := dumpIt(cmd, m)
			outputWithFormat(str, "Lorem")
			return
		})

	cmdr.NewInt(17).
		Titles("length", "l", "len").
		Group("").
		Placeholder("N").
		//Range(1,64).
		Description("characters length").
		AttachTo(cc)
	cmdr.NewInt(3).
		Titles("count", "C", "cnt").
		Group("").
		Placeholder("N").
		//Range(1,64).
		Description("words count").
		AttachTo(cc)

	for k, v := range m {
		cmdr.NewBool().
			Titles(k, v.short, v.aliases...).
			Description(fmt.Sprintf("generates %v field", v.label), "").
			ToggleGroup("Type").
			AttachTo(cc)
	}

}
