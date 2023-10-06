package cli

import (
	"fmt"
	"log"

	"github.com/ZakHargz/dotfiles_manager/config"
	"github.com/urfave/cli/v2"
)

type Action struct {
	flags struct {
		verbose bool
		dryRun  bool
	}
}

func (a *Action) ShowHelp(c *cli.Context) error {
	return cli.ShowAppHelp(c)
}

func (a *Action) RunInstall(c *cli.Context) error {
	a.GetFlags(c)
	a.Logs("Start %s\n", InstallCommand)

	cfg, err := a.LoadConfigs()
	if err != nil {
		return err
	}

	if err := cfg.Install(); err != nil {
		return fmt.Errorf("Failed to install configs: %w", err)
	}

	return nil
}

func (a *Action) RunUpdate(c *cli.Context) error {
	a.GetFlags(c)
	a.Logs("start %s\n", UpdateCommand)

	cfg, err := a.LoadConfigs()
	if err != nil {
		return err
	}

	if err := cfg.Update(); err != nil {
		return fmt.Errorf("failed to update config: %w", err)
	}

	return nil
}

func (a *Action) RunClean(c *cli.Context) error {
	a.GetFlags(c)
	a.Logs("start %s\n", CleanCommand)

	cfg, err := a.LoadConfigs()
	if err != nil {
		return err
	}

	if err := cfg.Clean(); err != nil {
		return fmt.Errorf("failed to clean config: %w", err)
	}

	return nil
}

func (a *Action) RunDiff(c *cli.Context) error {
	a.GetFlags(c)
	a.Logs("start %s\n", DiffCommand)

	cfg, err := a.LoadConfigs()
	if err != nil {
		return err
	}

	if err := cfg.Diff(); err != nil {
		return fmt.Errorf("failed to compare config: %w", err)
	}

	return nil
}

func (a *Action) LoadConfigs() (config.Config, error) {
	cfgReal, cfgDemo, err := config.LoadConfig(CurrentDirectory)
	if err != nil {
		return nil, fmt.Errorf("Failed to load dotfiles config: %w", err)
	}

	if a.flags.dryRun {
		return cfgDemo, nil
	}

	return cfgReal, nil
}

func (a *Action) GetFlags(c *cli.Context) {
	a.flags.verbose = c.Bool(VerboseFlag)
	a.flags.dryRun = c.Bool(DryRunFlag)
}

func (a *Action) Logs(format string, v ...interface{}) {
	if a.flags.verbose {
		log.Printf(format, v...)
	}
}
