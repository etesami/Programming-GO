package main

import (
	"fmt"
	"myapp/cmd"
	"os"
	"sort"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:        "MyApp",
		Usage:       "myapp --config config.yaml",
		Description: "This is a sample description of this app.",
		Flags:       cmd.CommonFlags,
		Commands: []*cli.Command{
			cmd.Start(),
			cmd.Complete(),
		},
		Action: cmd.Main(),
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
