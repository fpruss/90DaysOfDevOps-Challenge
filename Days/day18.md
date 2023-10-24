## Docker 3
Continued working through Tech World with Nana's Docker course.

### Container Port vs Host Port
Multiple containers can listen on the same port. A binding between host port and
container port is needed, so if you have for example 2 containers listening on
port 3000, then you need to bind a different host port to each container.
The port binding can be done on docker run, with the -p flag:
docker run -p<host-port>:<container-port> <image-name>

### Debugging
docker logs <container-id> - Fetch the logs of a container 

Its possible to give a container a name (otherwise each will get on of those funny random names).
This is done with the --name flag on docker run, like:
docker run --name <my-container-name> <image-name>

Then docker logs can be executed with specifying the container name instead the container id.

docker exec - Can be used to execute a command in a running container, i.e.:
docker exec -it <container-name> /bin/bash - executes an interactive bash shell on the container.
For more examples see second resource below.

## My notes on Michael's [corresponding day](https://www.90daysofdevops.com/2022/day43/)

## Resources
[Tech World with Nana - Docker for Beginners](https://www.youtube.com/watch?v=3c-iBn73dDE)
[Docker Exec Docs](https://docs.docker.com/engine/reference/commandline/exec/)
