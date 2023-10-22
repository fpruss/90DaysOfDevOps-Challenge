Now skipping day 19 to day 41 of Michaels days to continue with the docker days.
Will do the skipped days later on.

## My notes on Michael's [corresponding day](https://www.90daysofdevops.com/2022/day42/)
Introduction to containers, again written very sloppy, there are better resources
on this topic.

## My notes on the resources below
I think the attached resources are better than Michael's blog to learn or refresh docker knowledge.

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
container for an application to run.
It's the actual package, the artifact that can be moved around.
When the image is run, the container is created. The Image is just an artifact, that
if run becomes a container.

### Advantages of Comtainers
Advantages for developer machines: Instead of having to install applications like
MySql, and having to do it differently for each OS, developers just have to check
out the container for MySql, which has it's own OS, a start script and configuration.
Also starting an application is then done with the same command for all applications.
Another advantage of using docker containers, is the possibility to run different
versions of an application on the same machine without conflicts.
Also deployment is easier, because no environmental configuration is needed on the
server, only the docker runtime has to be installed.

## Resources
[Tech World with Nana - Docker for Beginners](https://www.youtube.com/watch?v=3c-iBn73dDE)
[Red Head - Containers](https://www.youtube.com/watch?v=3c-iBn73dDE)
[AWS - Containers and Images](https://aws.amazon.com/compare/the-difference-between-docker-images-and-containers/)
