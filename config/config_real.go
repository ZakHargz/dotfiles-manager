package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/make-go-great/copy-go"
	"github.com/make-go-great/diff-go"
)

type ConfigReal struct{ ConfigApps }

var _ Config = (*ConfigReal)(nil)

func (c *ConfigReal) Install() error {
	for _, app := range c.Apps {
		for _, p := range app.Paths {
			if err := copy.Replace(p.Internal, p.External); err != nil {
				return fmt.Errorf("failed to replace %s -> %s: %w", p.Internal, p.External, err)
			}
		}
	}

	return nil
}

func (c *ConfigReal) Update() error {
	for _, app := range c.Apps {
		for _, p := range app.Paths {
			if err := copy.Replace(p.External, p.Internal); err != nil {
				return fmt.Errorf("failed to replace %s -> %s: %w", p.External, p.Internal, err)
			}
		}
	}

	return nil
}

func (c *ConfigReal) Clean() error {
	unusedDirs, err := GetUnusedDirectories(c.Apps)
	if err != nil {
		return err
	}

	// Delete ununsed dirs to save some space
	for dir := range unusedDirs {
		dirPath := filepath.Join(ConfigDirectoryPath, dir)
		if err := os.RemoveAll(dirPath); err != nil {
			return fmt.Errorf("failed to remove %s: %w", dir, err)
		}
	}

	return nil
}

func (c *ConfigReal) Diff() error {
	for _, app := range c.Apps {
		for _, p := range app.Paths {
			if err := diff.Diff(p.Internal, p.External); err != nil {
				return fmt.Errorf("failed to compare %s with %s: %w", p.Internal, p.External, err)
			}
		}
	}

	return nil
}

func GetUnusedDirectories(apps map[string]App) (map[string]struct{}, error) {
	files, err := os.ReadDir(ConfigDirectoryPath)
	if err != nil {
		return nil, fmt.Errorf("Failed to read directory %s: %w", ConfigDirectoryPath, err)
	}

	unusedDirectories := make(map[string]struct{})
	for _, file := range files {
		if file.Name() == ConfigFile {
			continue
		}
		unusedDirectories[file.Name()] = struct{}{}
	}

	for name := range apps {
		delete(unusedDirectories, name)
	}

	return unusedDirectories, nil
}
