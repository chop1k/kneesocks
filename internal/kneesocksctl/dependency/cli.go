package dependency

import (
	"github.com/sarulabs/di"
	"github.com/urfave/cli"
	"socks/internal/kneesocksctl/cli/actions"
)

func registerCli(builder di.Builder) {
	appDef := di.Def{
		Name:  "cli_app",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			flags := ctn.Get("cli_global_flags").([]cli.Flag)
			start := ctn.Get("cli_start_command").(cli.Command)
			stop := ctn.Get("cli_stop_command").(cli.Command)

			app := cli.App{
				Name:  "kneesocksctl",
				Usage: "Command line utility to control kneesocks daemon.",
				Flags: flags,
				Commands: []cli.Command{
					start,
					stop,
				},
				Version: "v1.1.0",
			}

			app.UseShortOptionHandling = true

			return app, nil
		},
	}

	flagsDef := di.Def{
		Name:  "cli_global_flags",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return []cli.Flag{
				cli.VersionFlag,
				cli.HelpFlag,
			}, nil
		},
	}

	startCommandDef := di.Def{
		Name:  "cli_start_command",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			flags := ctn.Get("cli_start_flags").([]cli.Flag)
			action := ctn.Get("cli_start_action").(func(ctx *cli.Context) error)

			return cli.Command{
				Name:   "start",
				Flags:  flags,
				Action: action,
				Usage:  "Start kneesocks server daemon.",
			}, nil
		},
	}

	startFlagsDef := di.Def{
		Name:  "cli_start_flags",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return []cli.Flag{
				cli.StringFlag{
					Name:  "pid-file",
					Value: "/var/run/kneesocks.pid",
					Usage: "Where to save daemon process id.",
				},
				cli.StringFlag{
					Name:   "config",
					EnvVar: "config_path",
					Usage:  "Override default config path of kneesocks daemon.",
				},
				cli.StringFlag{
					Name:  "binary",
					Value: "kneesocks",
					Usage: "Path to kneesocks server binary.",
				},
				cli.StringFlag{
					Name:  "chroot",
					Value: "/var/lib/kneesocks",
					Usage: "Where to chroot daemon.",
				},
			}, nil
		},
	}

	stopCommandDef := di.Def{
		Name:  "cli_stop_command",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			flags := ctn.Get("cli_stop_flags").([]cli.Flag)
			action := ctn.Get("cli_stop_action").(func(ctx *cli.Context) error)

			return cli.Command{
				Name:   "stop",
				Flags:  flags,
				Action: action,
				Usage:  "Send stop signal to kneesocks daemon.",
			}, nil
		},
	}

	stopFlagsDef := di.Def{
		Name:  "cli_stop_flags",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return []cli.Flag{
				cli.StringFlag{
					Name:  "pid-file",
					Value: "/var/run/kneesocks.pid",
				},
			}, nil
		},
	}

	err := builder.Add(
		appDef,
		flagsDef,
		startCommandDef,
		startFlagsDef,
		stopCommandDef,
		stopFlagsDef,
	)

	if err != nil {
		panic(err)
	}

	registerActions(builder)
}

func registerActions(builder di.Builder) {
	startDef := di.Def{
		Name:  "cli_start_action",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return actions.Start, nil
		},
	}

	stopDef := di.Def{
		Name:  "cli_stop_action",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return actions.Stop, nil
		},
	}

	err := builder.Add(
		startDef,
		stopDef,
	)

	if err != nil {
		panic(err)
	}
}
