package pkg

import (
)

type PackageBuild struct {
  GitPath             string              `yaml:"gitpath"`
  BuildArgs           []string            `yaml:"build_args"`
}

type PackageVolume struct {
  HostPath            string              `yaml:"hostpath"`
  Path                string              `yaml:"path"`
}

type PackageCapability struct {

}

type PackageConfig struct {
  Name                string              `yaml:"name"`
  Build               PackageBuild        `yaml:"build"`
  Volumes             []PackageVolume     `yaml:"volumes"`
  Capabilities        []PackageCapability `yaml:"capabilities"`
}

