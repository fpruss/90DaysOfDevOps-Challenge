#Docker 1
Now skipping day 19 to day 41 of Michaels days to continue with the docker days.
Will do the skipped days later on.

## My notes on Michael's [corresponding day](https://www.90daysofdevops.com/2022/day42/)
Introduction to containers, again written very sloppy, there are better resources
on this topic.

## My notes on the resources below
I think the attached resources are better than Michael's blog to learn or refresh docker knowledge.

### What is Docker?
A platform for consistently build, run and ship applications.
The goal is to be intependend from infrastructure, so the application can run the
same way on different machines.

### What is a container?
A container is a way to package an application with their entire runtime enviroment,
including all the necessary dependencies and configuration data. This makes the contained
application portable and development and deployment more efficient.
A docker container consists of layers of images, a Linux image at the base (Alpine
most of the time, because of its small footprint), some intermediate images and an
application image at the top and then configuration data on top of that.

### What is a docker image?
A docker image is a read-only template that contains instructions for creating a
container. It's a snapshot of the libraries and dependencies required inside a
container for an application to run. This typically includes
- a cut-down OS
- a runtime environment like node
- application files
- third-party libraries
- environment variables
It's the actual package, the artifact that can be moved around.
When the image is run, the container is created. The Image is just an artifact, that
if run, becomes a container.

### Difference between Container and Image
Docker containers are runtime instances of Docker images, whether running or stopped.
Containers have a writable layer and the container runs the software.
Image is the blueprint of a container.

### Advantages of Containers
Without containers, dependencies like MySql would have to be installed on all machines and
it has to be taken care its always the correct version. Also the machines the application
should run on has to be configured always the same way (which could go wrong).
A container in comparison has it's own filesystem, dependencies, a start script and configuration.
This container is then deployed on all machines the application should run.
Also starting an application is then done with the same command for all applications.
Another advantage of using docker containers, is the possibility to run different
versions of an application on the same machine without conflicts.
Also deployment is easier, because no environmental configuration is needed on the
server, only the docker runtime has to be installed.

## Resources
[Tech World with Nana - Docker for Beginners](https://www.youtube.com/watch?v=3c-iBn73dDE)
[Red Head - Containers](https://www.youtube.com/watch?v=3c-iBn73dDE)
[AWS - Containers and Images](https://aws.amazon.com/compare/the-difference-between-docker-images-and-containers/)
[Docker Tutorial by Mosh](https://www.youtube.com/watch?v=pTFZFxd4hOI)
[Builduing Docker images](https://stackify.com/docker-build-a-beginners-guide-to-building-docker-images/)
