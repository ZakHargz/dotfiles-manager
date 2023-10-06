package cli

import (
	"fmt"
	"os"
	"runtime"

	"github.com/make-go-great/color-go"
	"github.com/urfave/cli/v2"
)

const (
	AppName  = "dottie"
	AppUsage = "Manage personal dotfiles"

	VerboseFlag  = "verbose"
	VerboseUsage = "Show what is going to happen"
	DryRunFlag   = "dry-run"
	DryRunUsage  = "Demo of the changes without changing any config files"

	InstallCommand = "install"
	InstallUsage   = "Install user configs from the data directory"

	UpdateCommand = "update"
	UpdateUsage   = "Update dotfiles from changes made in the user configs"

	CleanCommand = "clean"
	CleanUsage   = "Clean unused dotfiles from the system"

	DiffCommand = "diff"
	DiffUsage   = "Perform a diff between dotfiles and user configs"

	CurrentDirectory = "."
)

var DenyOperatingSystems = map[string]struct{}{
	"windows": {},
}

type DotfilesApp struct{ cliApp *cli.App }

func NewDotfilesManager() *DotfilesApp {
	a := &Action{}

	cliApp := &cli.App{
		Name:  AppName,
		Usage: AppUsage,
		Commands: []*cli.Command{
			{
				Name:  InstallCommand,
				Usage: InstallUsage,
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:  VerboseFlag,
						Usage: VerboseUsage,
					},
				},
				Action: a.RunInstall,
			},
			{
				Name:  UpdateCommand,
				Usage: UpdateUsage,
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:  VerboseFlag,
						Usage: VerboseUsage,
					},
				},
				Action: a.RunUpdate,
			},
			{
				Name:  CleanCommand,
				Usage: CleanUsage,
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:  VerboseFlag,
						Usage: VerboseUsage,
					},
				},
				Action: a.RunClean,
			},
			{
				Name:  DiffCommand,
				Usage: DiffUsage,
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:  VerboseFlag,
						Usage: VerboseUsage,
					},
				},
				Action: a.RunDiff,
			},
		},
	}

	return &DotfilesApp{
		cliApp: cliApp,
	}
}

func (a *DotfilesApp) Run() {
	if _, ok := DenyOperatingSystems[runtime.GOOS]; ok {
		color.PrintAppError(AppName, fmt.Sprintf("Sorry, %s is not supported right now", runtime.GOOS))
		return
	}

	if err := a.cliApp.Run(os.Args); err != nil {
		color.PrintAppError(AppName, err.Error())
	}
}
