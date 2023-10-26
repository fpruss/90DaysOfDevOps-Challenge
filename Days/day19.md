## Docker 4
Continued working through Tech World with Nana's Docker course.
According to the docs, "Docker Compose is a tool for defining and running multi-container
Docker applications. With Compose, you use a YAML file to configure your application's
services. Then, with a single command, you create and start all the services from your configuration."

What I found noteworthy: Docker Compose automatically creates a network for the
defined containers to run in, so no network has to be created and specified manually.
- docker compose up/down
- Dockerfile is a blueprint for building images 
- Dockerfile always is based on another image (i.e. NodeJs application needs NodeJs base image) 
CMD is the entrypoint command - in a Dockerfile there is just on CMD command (all
other RUN etc. commands are executed before)

- some stuff on private Docker registries (AWS ECR)

## Resources
[Tech World with Nana - Docker for Beginners](https://www.youtube.com/watch?v=3c-iBn73dDE)
[Docker Compose Docs](https://docs.docker.com/compose/)
