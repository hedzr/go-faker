package logic

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
	"time"

	"github.com/hedzr/cmdr"
	"syreclabs.com/go/faker"
)

func dumpHacker(oo faker.FakeHacker) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("    SaySomethingSmart   : %v\n", oo.SaySomethingSmart())) // => "If we connect the bus, we can get to the XML microchip through the digital TCP sensor!"
	sb.WriteString(fmt.Sprintf("    Abbreviation        : %v\n", oo.Abbreviation()))      // => "HTTP"
	sb.WriteString(fmt.Sprintf("    Adjective           : %v\n", oo.Adjective()))         // => "cross-platform"
	sb.WriteString(fmt.Sprintf("    Noun                : %v\n", oo.Noun()))              // => "interface"
	sb.WriteString(fmt.Sprintf("    Verb                : %v\n", oo.Verb()))              // => "bypass"
	sb.WriteString(fmt.Sprintf("    IngVerb             : %v\n", oo.IngVerb()))           // => "parsing"
	sb.WriteString(fmt.Sprintf("    Phrases             : \n%v\n", oo.Phrases()))         // => []string{
	return sb.String()
}

func genHackers(root *cmdr.RootCmdOpt) {

	cmdr.NewSubCmd().Titles("hacker", "hh", "hack").
		Description("generate Hacker names").
		Group("").
		TailPlaceholder("[text1, text2, ...]").
		Action(func(cmd *cmdr.Command, remainArgs []string) (err error) {
			oo := faker.Hacker()
			str := dumpHacker(oo)
			outputWithFormat(str, "Bitcoin")
			//        "If we bypass the program, we can get to the AGP protocol through the optical SDD alarm!",
			//        "We need to calculate the back-end XML microchip!",
			//        "Try to generate the GB bus, maybe it will hack the neural panel!",
			//        "You can't navigate the transmitter without synthesizing the optical SMS bus!",
			//        "Use the optical THX application, then you can override the mobile port!",
			//        "The CSS monitor is down, quantify the multi-byte bus so we can calculate the XSS bandwidth!",
			//        "Connecting the card won't do anything, we need to back up the multi-byte RSS card!",
			//        "I'll reboot the primary SMTP feed, that should monitor the XML protocol!`",
			//    }
			return
		}).
		AttachTo(root)

	cmdr.NewSubCmd().Titles("hacker-phrases", "hp").
		Description("generate Hacker Phrases names").
		Group("").
		TailPlaceholder("[text1, text2, ...]").
		Action(func(cmd *cmdr.Command, remainArgs []string) (err error) {
			var phrases []string

			for i := 1; i < 5; i++ {
				phrases = append(phrases, faker.Hacker().Phrases()...)
			}

			fmt.Println(strings.Join(phrases[:], "; "))
			return
		}).
		AttachTo(root)

	cmdr.NewSubCmd().Titles("hacker-phrases-colored", "hhc").
		Description("generate Colored Hacker Phrases names").
		Group("").
		TailPlaceholder("[text1, text2, ...]").
		Action(func(cmd *cmdr.Command, remainArgs []string) (err error) {
			var phrases []string

			for i := 1; i < 5; i++ {
				phrases = append(phrases, faker.Hacker().Phrases()...)
			}

			output := strings.Join(phrases[:], "; ")

			// r, g, b := 255, 215, 0 //gold color
			//
			// for j := 0; j < len(output); j++ {
			//	fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m", r, g, b, output[j])
			// }

			for j := 0; j < len(output); j++ {
				r, g, b := rgb(j)
				fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m", r, g, b, output[j])
			}
			fmt.Println()
			return
		}).
		AttachTo(root)

	hhp := cmdr.NewSubCmd().Titles("hacker-phrases-colored-piped", "hhp").
		Description("generate Colored Hacker Phrases names, let's work as a pipe").
		Group("").
		TailPlaceholder("[text1, text2, ...]").
		Action(func(cmd *cmdr.Command, remainArgs []string) (err error) {

			info, _ := os.Stdin.Stat()
			if info.Mode()&os.ModeCharDevice != 0 {
				fmt.Println("The command is intended to work with pipes.")
				fmt.Println("Usage: fortune | gorainbow")
				fmt.Println("Usage: fortune | go run ./cli/faker/cli/faker hhp [-j]")
				fmt.Println("Usage: fortune -n 300 | cowsay | go run ./cli/faker/cli/faker hhp [-j]")
				return
			}

			if cmdr.GetBoolRP(cmd.GetDottedNamePath(), ".just-in-time") {
				reader := bufio.NewReader(os.Stdin)
				j := 0
				for {
					input, _, err := reader.ReadRune()
					if err != nil && err == io.EOF {
						break
					}
					r, g, b := rgb(j)
					fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m", r, g, b, input)
					time.Sleep(40 * time.Millisecond)
					j++
				}
				return
			}

			var output []rune

			reader := bufio.NewReader(os.Stdin)
			for {
				input, _, err := reader.ReadRune()
				if err != nil && err == io.EOF {
					break
				}
				output = append(output, input)
			}

			printZ(output)
			return
		}).
		AttachTo(root)

	cmdr.NewBool().
		Titles("just-in-time", "j").
		Description("render the char from pipe one-by-one (just-in-time)").
		AttachTo(hhp)

}

func rgb(i int) (int, int, int) {
	var f = 0.1
	return int(math.Sin(f*float64(i)+0)*127 + 128),
		int(math.Sin(f*float64(i)+2*math.Pi/3)*127 + 128),
		int(math.Sin(f*float64(i)+4*math.Pi/3)*127 + 128)
}

func printZ(output []rune) {
	for j := 0; j < len(output); j++ {
		r, g, b := rgb(j)
		fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m", r, g, b, output[j])
	}
	fmt.Println()
}
