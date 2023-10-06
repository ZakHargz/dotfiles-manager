package config

import "fmt"

type ConfigDemo struct{ ConfigApps }

var _ Config = (*ConfigDemo)(nil)

func (c *ConfigDemo) Install() error {
	for _, app := range c.Apps {
		for _, p := range app.Paths {
			fmt.Printf("Replace %s -> %s\n", p.Internal, p.External)
		}
	}
	return nil
}

func (c *ConfigDemo) Update() error {
	for _, app := range c.Apps {
		for _, p := range app.Paths {
			fmt.Printf("Replace %s -> %s\n", p.Internal, p.External)
		}
	}
	return nil
}

func (c *ConfigDemo) Clean() error {
	unusedDirectories, err := GetUnusedDirectories(c.Apps)
	if err != nil {
		return err
	}

	for dir := range unusedDirectories {
		fmt.Printf("Remove %s\n", dir)
	}

	return nil
}

func (c *ConfigDemo) Diff() error {
	return nil
}
