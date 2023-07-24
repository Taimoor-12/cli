package main

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"github.com/ipinfo/cli/lib"
	"github.com/ipinfo/cli/lib/complete"
	"github.com/ipinfo/cli/lib/complete/predict"
	"github.com/spf13/pflag"
	"os"
	"strings"
)

var completionsN2IP6 = &complete.Command{
	Flags: map[string]complete.Predictor{
		"--nocolor": predict.Nothing,
		"-h":        predict.Nothing,
		"--help":    predict.Nothing,
	},
}

func printHelpN2IP6() {
	fmt.Printf(
		`Usage: %s n2ip6 [<opts>] <expr>

Example:
  %[1]s n2ip6 "190.87.89.1"
  %[1]s n2ip6 "2001:0db8:85a3:0000:0000:8a2e:0370:7334
  %[1]s n2ip6 "2001:0db8:85a3::8a2e:0370:7334
  %[1]s n2ip6 "::7334
  %[1]s n2ip6 "7334::""
	

Options:
  General:
    --nocolor
      disable colored output.
    --help, -h
      show help.
`, progBase)
}

func cmdN2IP6() error {
	pflag.BoolVarP(&fHelp, "help", "h", false, "show help.")
	pflag.BoolVar(&fNoColor, "nocolor", false, "disable colored output.")
	pflag.Parse()

	if fNoColor {
		color.NoColor = true
	}

	if fHelp {
		printHelpDefault()
		return nil
	}

	var err error

	cmd := ""
	if len(os.Args) > 2 {
		cmd = os.Args[2]
	}

	// If no argument is given, print help.
	if strings.TrimSpace(cmd) == "" {
		printHelpN2IP6()
		return nil
	}

	if lib.IsInvalid(cmd) {
		return errors.New("invalid expression")
	}
	tokens, err := lib.TokeinzeExp(cmd)
	if err != nil {
		return err
	}

	postfix := lib.InfixToPostfix(tokens)

	result, err := lib.EvaluatePostfix(postfix)
	if err != nil {
		return err
	}

	res := lib.DecimalToIP(result.String(), true)
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v\n", err)
		printHelpN2IP6()
		return nil
	}

	fmt.Println(res)
	return nil
}
