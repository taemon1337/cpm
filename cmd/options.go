package cmd

import (

)

var DefaultCommandOptions = &CommandOptions{
  Verbose:        "1",
  Branch:         "master",
}

type CommandOptions struct {
  Verbose         string
  Branch          string
}

func NewCommandOptions() *CommandOptions {
  return &CommandOptions{}
}
