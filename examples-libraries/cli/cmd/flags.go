package cmd

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func Start() *cli.Command {
	return &cli.Command{
		Name:    "add",
		Aliases: []string{"a"},
		Usage:   "add a task to the list",
		Flags:   CommonFlags,
		Action: func(*cli.Context) error {
			fmt.Println("added task: ", os.Args)
			return nil
		},
	}
}

func Complete() *cli.Command {
	return &cli.Command{
		Name:    "complete",
		Aliases: []string{"c"},
		Usage:   "complete a task on the list",
		Flags:   CommonFlags,
		Action: func(*cli.Context) error {
			fmt.Println("completed task: ", os.Args)
			return nil
		},
	}
}

func Main() func(cCtx *cli.Context) error {
	return func(cCtx *cli.Context) error {
		name := "Nefertiti"
		if cCtx.NArg() > 0 {
			name = cCtx.Args().Get(0)
			fmt.Println("Name: ", name)
		}

		if cCtx.Bool("g") {
			fmt.Println("Argument g does exist!")
			os.Exit(0)
		}

		if cCtx.String("lang") == "spanish" {
			fmt.Println("Hola", name)
		} else {
			fmt.Println("Hello", name)
		}
		return nil
	}
}

var CommonFlags []cli.Flag = []cli.Flag{
	&cli.StringFlag{
		Name:    "config",
		Usage:   "Specify a path to a edgevpn config file",
		EnvVars: []string{"EDGEVPNCONFIG"},
	},
	&cli.BoolFlag{
		Name:  "g",
		Usage: "Generates a new configuration and prints it on screen",
	},
	&cli.StringFlag{
		Name:     "lang",
		Value:    "english",
		Usage:    "language for the greeting",
		EnvVars:  []string{"LEGACY_COMPAT_LANG", "APP_LANG", "LANG"},
		Required: true,
	},
	&cli.IntFlag{
		Name:        "port",
		Usage:       "Use a randomized port",
		Value:       0,
		DefaultText: "random",
		Required:    true,
		Action: func(ctx *cli.Context, v int) error {
			if v >= 65536 {
				return fmt.Errorf("Flag port value %v out of range[0-65535]", v)
			}
			return nil
		},
	},
}
