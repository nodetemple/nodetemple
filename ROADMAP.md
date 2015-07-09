= Roadmap

- nodectl: semi-automatic control plane for servers of selected cloud provider and many other helpful functions for easier life
- noded: initializes environment, sets up systemd and fleet units, watch and fix if anything gets wrong (service down, lack/excess of master/node servers, etc.)
- auto-check for updates of most important components (rolling update): node, etcd, fleet, flannel, rkt/docker, kubernetes, CoreOS
- rolling update for systemd and fleet units (probably zero-downtime)
- off-server usage of nodectl - connect to any server which is running etcd and/or noded and control your cluster
- encrypted and authenticated traffic between servers/services
- firewall included
- dockerize and/or rocketize all services (which are possible)
- private docker registry, highly-available one, with persistent storage, accessible only within cluster, git integration, auto-build and deploy with rolling update for pods, etc.
- web UI
- monitor resources, cluster-wide, CLI tool and web UI
- cluster-wide HA volumes
- API and/or socket
- cluster-wide logging
