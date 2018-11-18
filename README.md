# Container Package Manager (cpm)
Container Package Management for the Desktop.

## How is this different then Docker Hub?
Docker Hub is simply for distributing Docker Images, where CPM is for distributing desktop applications as containers.

## How is this different then RPM, YUM, NPM, GEM, PIP, etc?
The main difference is CPM is the equivalent of all these package managers but for containers as desktop applications.

## Help and Usage

### List of locally installed packages
cpm list

### Search the configured package registries for a package
cpm search <query>

### Install a new container package
cpm install <name>

### Update a container package
cpm update <name>

### Build a container package from source
cpm build <name>

### Remove a container package
cpm remove <name>

### List currently configured container package registries
cpm registry list

### Add a new cpm registry
cpm registry add <registry-url>

### Install a container runtime
cpm install docker
cpm install cri-o
cpm install rkt
cpm install runc

### Configure a container runtime to use
cpm install docker --default


### Configure a github repo as a cpm registry
cpm looks for a .package.yaml file in the base folder of a configured git repo.
```yaml
name: <package-name>
build:
  path: ./docker-build
  command: "docker build -it <package-name>"
volumes:
- name: local
  path: $HOME/.local/$package
```

