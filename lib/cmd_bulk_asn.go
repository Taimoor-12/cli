package lib

import (
	"github.com/ipinfo/go/v2/ipinfo"
	"github.com/spf13/pflag"
	"strings"
)

// CmdBulkASNFlags are flags expected by CmdBulkASN
type CmdBulkASNFlags struct {
	Token   string
	nocache bool
	help    bool
	Field   []string
	json    bool
	Yaml    bool
}

// Init initializes the common flags available to CmdBulkASN with sensible
// defaults.
// pflag.Parse() must be called to actually use the final flag values.
func (f *CmdBulkASNFlags) Init() {
	_h := "see description in --help"
	pflag.StringVarP(
		&f.Token,
		"token", "t", "",
		_h,
	)
	pflag.BoolVarP(
		&f.nocache,
		"nocache", "", false,
		_h,
	)
	pflag.BoolVarP(
		&f.help,
		"help", "h", false,
		_h,
	)
	pflag.StringSliceVarP(
		&f.Field,
		"field", "f", []string{},
		_h,
	)
	pflag.BoolVarP(
		&f.json,
		"json", "j", false,
		_h,
	)
	pflag.BoolVarP(
		&f.Yaml,
		"yaml", "y", false,
		_h,
	)
}

func CmdBulkASN(ii *ipinfo.Client, args []string) (ipinfo.BatchASNDetails, error) {
	var asns []string

	if !validateWithFunctions(args, []func(string) bool{StrIsASNStr}) {
		return nil, ErrInvalidInput
	}

	for _, arg := range args {
		asns = append(asns, strings.ToUpper(arg))
	}

	return ii.GetASNDetailsBatch(asns, ipinfo.BatchReqOpts{
		TimeoutPerBatch:              60 * 30, // 30min
		ConcurrentBatchRequestsLimit: 20,
	})

}
