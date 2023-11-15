## Kubernetes
I found a very good Udemy course on Kubernetes (see below), but also did the
Tech with Nana Kubernetes full course and worked through Michael cades stuff.
I will summarize my learnings in the following days.
Docker compose is good to start multiple containers together, but in production,
automated scaling, monitoring, load balancing and such is needed.
This is called orchestration and Kubernetes is the most popular a container
orchestration platform. Others are Docker swarm (easy to setup, but lacks
advanced features) and Mesos (difficult to setup).
Kubernetes provides a lot of options to customize deployments and supports
deployments of complex architectures.

### Advantages of Kubernetes
- application is highly available
- service discovery and load balancing
- storage orchestration (K8s allows to mount local storage, cloud storage and more)
- automated scaling on service level
- easy configuration with code
- automated rollouts and rollbacks
- self-healing (K8s restarts failed containers, replaces containers, kills containers
that are unresponsive and doesn't advertise containers that are not ready to serve
- secret and configuration management (secrets and application configuration can
be deployed and updated without rebuildung container images and without exposing
secrets in the stack configuration

### Architecture
#### Node
- is a machine, physical or virtual, where K8s is installed and where K8s will
launch containers
- more than one node is needed to account for machine failure

#### Cluster
- set of nodes grouped together, to keep the application available, if one node
fails and for sharing load

#### Master
- responsible for managing the cluster, for monitoring, for moving the workload
of a failed node to another node
- stores information about the cluster members
- watches the worker nodes of the cluster and is responsible for the actual orchestration
of containers on the worker nodes
- is a node that is configured as master

#### Components
- API server (to interact with a K8s cluster)
- etcd service (key value store to store all data used to manage the cluster; is
responsible for implementing locks within the cluster to prevent conflicts between
multiple masters)
- kublet service (agent that runs on each node, ensures that the containers are
running on the nodes as expected)
- container runtime (i.e. Docker)
- controllers (responsible for noticing and responding when nodes, containers or
endpoints are down)
- schedulers (distributes containers across multiple nodes)

#### Master vs Worker Nodes
- worker node has the container runtime and the kublet agent, that interacts with
the master's kube API server. The master also has the controller and scheduler and
a key value store.

#### kubectl
- cli tool to deploy and manage applications on a K8s cluster.
- i.e. kubectl run (deploy an application on a cluster,
kubectl cluster-info, kubectl get nodes (to list all nodes of the cluster)

## Resources
[Michael Cades corresponding day](https://www.90daysofdevops.com/2022/day49/)
[Udemy Course](https://www.udemy.com/course/learn-kubernetes/)
[Tech with Nana - Kubernetes Full Course](https://www.youtube.com/watch?v=X48VuDVv0do)
