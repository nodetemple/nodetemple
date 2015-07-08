= Roadmap

- nodectl: semi-automatic control plane for servers of selected cloud provider and many other helpful functions for easier life
- noded: initializes environment, sets up systemd and fleet units, watch and fix if anything gets wrong (service down, lack/excess of master/node servers, etc.)
- auto-check for updates of most important components (rolling update): node, etcd, fleet, flannel, rkt/docker, kubernetes, CoreOS
- rolling update for systemd and fleet units (probably zero-downtime)
- off-server usage of nodectl - connect to any server which is running etcd and/or noded and control your cluster
- encrypted and authenticated traffic between servers/services
- firewall included
- dockerize and/or rocketize all services (which are possible)
