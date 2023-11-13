## Docker 4
Continued working through Tech World with Nana's Docker course.
According to the docs, "Docker Compose is a tool for defining and running multi-container
Docker applications. With Compose, you use a YAML file to configure your application's
services. Then, with a single command, you create and start all the services from your configuration."

What I found noteworthy: Docker Compose automatically creates a network for the
defined containers to run in, so no network has to be created and specified manually.
- docker compose up/down
- Docker can build images automatically by reading the instructions from a Dockerfile.
- a Dockerfile is a text document that contains all the commands a user could call on the command line
to assemble an image -> a Dockerfile is a blueprint for building images
- a Dockerfile always is based on another image (i.e. NodeJs application needs NodeJs base image) 
CMD is the entrypoint command - in a Dockerfile there is just on CMD command (all
other RUN etc. commands are executed before)

- some stuff on private Docker registries (AWS ECR)

### Docker layers
A Docker image consists of read-only layers, each of which represents a Dockerfile instruction.
These layers are stacked and each layer builds upon the previous one.
When you run an image and generate a container, you add a new writable layer, also
called the container layer, on top of the underlying layers.
All changes made to the running container, like writing files, are written to this layer.
For each command in a Dockerfile Docker first tries to use an existing cached layer
if possible. So if you for example always copy all files and then run "npm install",
the "npm install" will run each time any of the files has changed.
It is more effective add a line to copy only the package.json and then run "npm install".
This way "npm install" only runs if the package.json changed, otherwise Docker will
use the cache and skip that part. After this, all the other files can be copied.

## Resources
[Tech World with Nana - Docker for Beginners](https://www.youtube.com/watch?v=3c-iBn73dDE)
[Docker Compose Docs](https://docs.docker.com/compose/)
[Docker Tutorial by Mosh](https://www.youtube.com/watch?v=pTFZFxd4hOI)
[Dockerfile Docs](https://docs.docker.com/engine/reference/builder/)
[Docker Layers](https://bitjudo.com/blog/2014/03/13/building-efficient-dockerfiles-node-dot-js/)
[Dockerfile best practices](https://docs.docker.com/develop/develop-images/dockerfile_best-practices/)
