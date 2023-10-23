# Docker 2
Continued working through Tech World with Nana's Docker course.

### Difference between Docker and VM
OS have 2 layers - Kernel and application layer

---------
Applications
---------
Kernel (communicates with the hardware)
---------
Hardware
---------

Docker and VM are both virtualization tools (virtualization means creating a
virtual representation of something).
Docker virtualizes the application layer and uses the Kernel of the host.
VM virtualizes the complete OS with application layer and Kernel.
- Docker images are smaller than VM images, because they just have to implement on layer.
- Docker containers can be started faster, because the Kernel doesn't have to be booted.
- Compatibility: VM of any OS can run on any host OS, this is not always possible
with Docker, so not every host OS can run each type of Docker images.

### Some Docker commands
docker pull - download an image from a registry (i.e. from the Docker Hub registry or a different registry)
docker start [OPTIONS] CONTAINER [CONTAINER...] - starts one or more stopped containers
docker stop - stops one or more containers
docker run [OPTIONS] IMAGE [COMMAND] [ARG...] - combines docker pull and docker start 
(-d to run detached). Takes an image id to create a new container
docker ps - list containers (to show both running and stopped containers, use -a flag)
docker images - list images

## Resources
[Tech World with Nana - Docker for Beginners](https://www.youtube.com/watch?v=3c-iBn73dDE)
[Docker Docs - Commands](https://docs.docker.com/engine/reference/commandline/docker/)
