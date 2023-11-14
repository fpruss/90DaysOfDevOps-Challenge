## Docker Security
### Kernel Namespaces
Docker creates a set of namespaces and control groups for each container.
Processes running in a container cannot see or affect processes in another container
or in the host system.
Each container has its own network stack, so containers only can interact with other
containers or external hosts through their respective network interfaces. IP traffic
is allowed when public ports are specified for a container.
All Containers on a host are connected through a bridge interface, that behaves
like a virtual network switch.

### Control Groups
Control groups implement resource accounting and limiting, they provide metrics,
but also help to ensure that each container gets the resources it needs, but
also that a single container can't exhaust one of those resources (memory, CPU,
disk I/O, ...).
They are essential to mitigate DoS attacks.

### Docker Daemon Attack Surface
The docker daemon requires root privileges unless the rootless mode is activated
(see resource below).
Of course only trusted users should be allowed to control the Docker daemon.
That's because Docker allows for example to share a directory between the host
a guest container. If i.e. a container is started where the /host directory is the
root directory on the host, the container can alter the host filesystem without
restrictions.
If Docker is called through an API to provision containers, parameters have to
be checked thoroughly, to make sure a malicious user cannot pass crafted parameters
that cause Docker to create arbitrary containers.
API endpoints have to be secured with HTTPS and certificates, because even when
a firewall limits access to the API to other hosts in the network, access could
be possible from other containers, which could result in privilege escalation.
Also, such an API only should be accessible from a trusted network or VPN.

If Docker is run on a server, it is recommended to run Docker exclusively on the
server, and to dockerize all other services. But of course admin tools like an
SSH server and monitoring processes, like NRPE and collectd, are allowed (and
needed) to stay undockerized.

### Linux kernel capabilities
In most cases, containers do not need "real" root privileges, because processes
that normally run as root are handled by the infrastructure around the container.
For example:
- SSH access is typically managed by a single server running on the Docker host
- log management is typically handed to Docker or a 3rd party service
- no hardware management is needed inside containers
- network management happens outside of containers

That means, containers can run with a reduced capability set, so that "root"
within a container has much less privileges than the real "root"
For instances, it is possible to:
- deny all "mount" operations
- deny access to raw sockets (to prevent packet spoofing)
- restrict access to filesystem operations, like changing the owner of files

**Important**
This means, that even if an attacker escalates to root within a container,
damage is mitigated through the limited capabilities.
But of course, the set of capabilitites and mounts given to a container may
provide incomplete isolation.
The best practice is to remove all capabilities (from the defaults) except
those explicitly required for their processes.

### Docker Content Trust Signature Verification
The Docker Enginge can be configured to only run signed images.
When enabled, only repositories signed with a user-specified root key can be
pulled and run.

### Other Kernel security features
The following is directly taken from the docs:
While Docker currently only enables capabilities, it doesn't interfere with the other systems. This means that there are many different ways to harden a Docker host. Here are a few examples.

    You can run a kernel with GRSEC and PAX. This adds many safety checks, both at compile-time and run-time; it also defeats many exploits, thanks to techniques like address randomization. It doesn't require Docker-specific configuration, since those security features apply system-wide, independent of containers.
    If your distribution comes with security model templates for Docker containers, you can use them out of the box. For instance, we ship a template that works with AppArmor and Red Hat comes with SELinux policies for Docker. These templates provide an extra safety net (even though it overlaps greatly with capabilities).
    You can define your own policies using your favorite access control mechanism.

Just as you can use third-party tools to augment Docker containers, including special network topologies or shared filesystems, tools exist to harden Docker containers without the need to modify Docker itself.

As of Docker 1.10 User Namespaces are supported directly by the docker daemon. This feature allows for the root user in a container to be mapped to a non uid-0 user outside the container, which can help to mitigate the risks of container breakout. This facility is available but not enabled by default.

## Resources
[Docker Docs - Security](https://docs.docker.com/engine/security/)
[Arch Wiki - Network Bridge](https://wiki.archlinux.org/title/network_bridge)
[Baeldung - Bridging Network Interfaces](https://www.baeldung.com/linux/bridging-network-interfaces)
[Docker Docs - Rootless mode](https://docs.docker.com/engine/security/rootless/)
