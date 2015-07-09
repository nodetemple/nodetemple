# Nodetemple

CoreOS and Kubernetes cluster orchestration tool
> **Warning!** This software is not production ready - use it at your own risk!

Nodetemple is an open source system for easier orchestration of multiple CoreOS hosts on a cluster, it provides automated and semi-automated mechanisms for deployment, maintenance, and scaling of servers, clusters and cluster-wide services, such as etcd, fleet, flannel, docker/rocket, kubernetes, etc.

## Getting started

### Before you dive in

Nodetemple is designed to run on [CoreOS](https://coreos.com), basically because of [native systemd support](https://coreos.com/using-coreos/systemd/), [security](https://coreos.com/security/), [automatic updates engine](https://coreos.com/using-coreos/updates/), etc.

Supported IaaS providers:
- [DigitalOcean](https://www.digitalocean.com/)
- *Comming soon...* ~~GCE, AWS, Rackspace, Azure, Wmware, Vagrant, bare-metal, etc.~~

### Building

Nodetemple must be built with Go 1.4+ on a Linux machine. Simply run `./build` and then copy the binaries out of bin/ onto each of your machines.

## Project details

### Release notes

See the [releases tab](https://github.com/nodetemple/nodetemple/releases) for more information on each release.

### Contributing

Got great ideas? Awesome! Contribute! Feel free to join us! Open source software rules!

By contributing to this project you agree to the Developer Certificate of Origin (DCO). This document was created by the Linux Kernel community and is a simple statement that you, as a contributor, have the legal right to make the contribution. See the [DCO](DCO) file for details.

### License

Nodetemple is released under the Apache 2.0 license. See the [LICENSE](LICENSE) file for details.

Specific components of Nodetemple use code derivative from software distributed under other licenses; in those cases the appropriate licenses are stipulated alongside the code.

### Community, support and information

- Nodetemple website: [nodetemple.com](https://nodetemple.com)
- Nodetemple on Twitter [@nodetemple](https://twitter.com/nodetemple)
- Nodetemple repository short URL: [git.io/nodetemple](https://git.io/nodetemple)
