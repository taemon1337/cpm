package pkg

import (
)

type PackageBuild struct {
  Dockerfile          string              `yaml:"dockerfile"`
  GitPath             string              `yaml:"gitpath"`
  BuildArgs           []string            `yaml:"build_args"`
}

type PackageVolume struct {
  HostPath            string              `yaml:"hostpath"`
  Path                string              `yaml:"path"`
}

type PackageConfig struct {
  Name                string              `yaml:"name"`
  Version             string              `yaml:"version"`
  Registry            string              `yaml:"registry"`
  Build               PackageBuild        `yaml:"build"`
  CapAdd              []string            `yaml:"cap_add"`
  CapDrop             []string            `yaml:"cap_drop"`
  Privileged          bool                `yaml:"privileged"`
  Volumes             []PackageVolume     `yaml:"volumes"`
  Devices             []string            `yaml:"devices"`
  Network             string              `yaml:"network"`
  Ports               []string            `yaml:"ports"`
  Pid                 string              `yaml:"pid"`
  Restart             string              `yaml:"restart"`
}

