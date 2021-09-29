// Code generated by goa v3.5.2, DO NOT EDIT.
//
// sdk-utilities gRPC client CLI support package
//
// Command:
// $ goa gen github.com/allinbits/sdk-service-meta

package cli

import (
	"flag"
	"fmt"
	"os"

	sdkutilitiesc "github.com/allinbits/sdk-service-meta/gen/grpc/sdk_utilities/client"
	goa "goa.design/goa/v3/pkg"
	grpc "google.golang.org/grpc"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `sdk-utilities supply
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` sdk-utilities supply --message '{
      "chainName": "Ea eos.",
      "port": 2673035619671438831
   }'` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(cc *grpc.ClientConn, opts ...grpc.CallOption) (goa.Endpoint, interface{}, error) {
	var (
		sdkUtilitiesFlags = flag.NewFlagSet("sdk-utilities", flag.ContinueOnError)

		sdkUtilitiesSupplyFlags       = flag.NewFlagSet("supply", flag.ExitOnError)
		sdkUtilitiesSupplyMessageFlag = sdkUtilitiesSupplyFlags.String("message", "", "")
	)
	sdkUtilitiesFlags.Usage = sdkUtilitiesUsage
	sdkUtilitiesSupplyFlags.Usage = sdkUtilitiesSupplyUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "sdk-utilities":
			svcf = sdkUtilitiesFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "sdk-utilities":
			switch epn {
			case "supply":
				epf = sdkUtilitiesSupplyFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "sdk-utilities":
			c := sdkutilitiesc.NewClient(cc, opts...)
			switch epn {
			case "supply":
				endpoint = c.Supply()
				data, err = sdkutilitiesc.BuildSupplyPayload(*sdkUtilitiesSupplyMessageFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// sdk-utilitiesUsage displays the usage of the sdk-utilities command and its
// subcommands.
func sdkUtilitiesUsage() {
	fmt.Fprintf(os.Stderr, `sdk-utilities performs Cosmos SDK-related operations
Usage:
    %[1]s [globalflags] sdk-utilities COMMAND [flags]

COMMAND:
    supply: Supply implements supply.

Additional help:
    %[1]s sdk-utilities COMMAND --help
`, os.Args[0])
}
func sdkUtilitiesSupplyUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] sdk-utilities supply -message JSON

Supply implements supply.
    -message JSON: 

Example:
    %[1]s sdk-utilities supply --message '{
      "chainName": "Ea eos.",
      "port": 2673035619671438831
   }'
`, os.Args[0])
}
