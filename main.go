package main

import (
  "os"
  "log"
  "flag"

  "github.com/taemon1337/cpm/cmd"
  "github.com/taemon1337/cpm/pkg"
  "github.com/taemon1337/cpm/config"
)

func command(cpm *pkg.ContainerPackageManager, c string, args []string, options *cmd.CommandOptions) (pkg.Printable, error) {
  switch c {
    case "config":
      return cpm.Config(args, options)
    case "list":
      return cpm.List(args, options)
    case "search":
      return cpm.Search(args, options)
    case "build":
      return cpm.Build(args, options)
    case "install":
      return cpm.Install(args, options)
    case "update":
      return cpm.Update(args, options)
    case "remove":
      return cpm.Remove(args, options)
    case "start":
      return cpm.Start(args, options)
    case "stop":
      return cpm.Stop(args, options)
    case "restart":
      return cpm.Restart(args, options)
    case "pause":
      return cpm.Pause(args, options)
    case "registry":
      return cpm.Registry(args, options)
    default:
      return cpm.Help(args, options)
  }
}

func main() {
  defaults := cmd.DefaultCommandOptions
  options := cmd.NewCommandOptions()

  verbose := flag.String("verbose", defaults.Verbose, "Set the verbosity level (1-5)")
  branch := flag.String("branch", defaults.Branch, "Set the branch of the git repo")
  configfile := flag.String("config", defaults.ConfigFile, "Set a custom config file to use")
  flag.Parse()

  options.Branch = *branch
  options.ConfigFile = *configfile

  cfg, err := config.LoadPackageManagerConfig(options.ConfigFile)
  if err != nil {
    log.Fatal("Could not load or initialize package manager config: %s", err)
  }

  cfg.Verbose = *verbose
  cpm := pkg.NewContainerPackageManager(cfg)
  args := flag.Args()

  if len(args) > 0 {
    res, err := command(cpm, args[0], args[1:], options)

    if err != nil {
      log.Fatal(err)
    }

    res.Print()

    os.Exit(0)
  } else {
    os.Exit(1)
  }
  return
}
