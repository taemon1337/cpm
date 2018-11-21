package cmd

import (
  "os"
  "path"
)

var config_file = path.Join(os.Getenv("HOME"), ".cpm", "config")

var DefaultCommandOptions = &CommandOptions{
  Verbose:        "1",
  Branch:         "master",
  ConfigFile:     config_file,
}

type CommandOptions struct {
  Verbose         string
  Branch          string
  ConfigFile      string
}

func NewCommandOptions() *CommandOptions {
  return &CommandOptions{}
}
