# Dockerfile.tests

> This document contains all the information related to the confection of the [`Dockerfile.tests`](../Dockerfile.tests) file used to create an image to run the unit tests of this project.

The content of the testing Dockerfile is the following:

```Dockerfile
FROM golang:1.15.5-alpine

WORKDIR /go/src/github.com/harvestcore/HarvestCCode
ENV CGO_ENABLED 0

RUN apk update --no-cache; apk add --no-cache make git

CMD cp -R --symbolic-link /app/test/* /go/src/github.com/harvestcore/HarvestCCode; \
    make test
```

## Base image

The base image selected is [`golang:1.15.5-alpine`](https://github.com/docker-library/golang/blob/071e264f53e89ea75f1a38f6c1c33641685d8560/1.15/alpine3.12/Dockerfile). It is a lightweigh image that already includes Golang v1.15.5 installed, without any extra unnecesary software. Since I'm only going to run some unit tests in the container I don't need an image that includes extra software, like [`golang:1.15.5-buster`](https://github.com/docker-library/golang/blob/071e264f53e89ea75f1a38f6c1c33641685d8560/1.15/buster/Dockerfile) (based on debian) or [`golang:1.15.5-windowsservercore-1809`](https://github.com/docker-library/golang/blob/071e264f53e89ea75f1a38f6c1c33641685d8560/1.15/windows/windowsservercore-1809/Dockerfile) (based on Windows).

On the other hand, having a base image with Golang already installed saves me some time.

## Dockerfile explanation

> In this section I'll explain each line of the Dockerfile.

```Dockerfile
FROM golang:1.15.5-alpine
```

Base image, explained above.

```Dockerfile
WORKDIR /go/src/github.com/harvestcore/HarvestCCode
```

`WORKDIR` sets the working directory in which all commands (like `RUN`, `CMD` or `COPY`) are executed. In my case I've set it to `/go/src/github.com/harvestcore/HarvestCCode` since `/go/src` is the usual directory in which all Golang projects are stored. `github.com/harvestcore/HarvestCCode` is also the name of the package of my project.

```Dockerfile
ENV CGO_ENABLED 0
```

Golang is able to run C code, and Golang by itself sometimes makes use of this ([`cgo` package](https://golang.org/cmd/cgo/)). Since I'm using an Alpine image which does not include the GCC compiler, it is possible to avoid this code execution by setting the environment variable `CGO_ENABLED` to `0`. By doing this I'm force a build using just Golang. This has no impact at all and allows me to save some bytes in the image size.

```Dockerfile
RUN apk update --no-cache; apk add --no-cache make git
```

Simple run command that updates the index of packages and after that installs Git (used by Golang to fetch the needed packages) and Make (since it is the task manager I'm using). I've added the `--no-cache` flag to avoid storing cache files when installing the needed packages (this also saves me some bytes).

```Dockerfile
CMD cp -R --symbolic-link /app/test/* /go/src/github.com/harvestcore/HarvestCCode; make test
```

This image will be tested using the command `docker run -t -v /some/path:/app/test nick-estudiante/nombre-del-repo` this means that a volume will be mounted in `/app/test` and all the project files will be there. The command `cp -R --symbolic-link /app/test/* /go/src/github.com/harvestcore/HarvestCCode` creates symbolic links recursively (from all the files and directories in `/app/test`) in the directory I've set as `WORKDIR`. The project files are not available in build time, so this command must be run in runtime. By creating symbolic links I can avoid copying the raw files into the desired directory. After that I just execute `make test`, which installs all the needed testing dependencies and runs the unit tests.

## DockerHub

First of all I've created a new repository in DockerHub ([this one](https://hub.docker.com/r/harvestcore/harvestccode)). After that I've configured the automated build as shown below:

- I've linked my GitHub account with DockerHub.
- I've set the build rules:
  - Branch: `master`
  - Dockerfile location: `Dockerfile.tests`
  - Autobuild: enabled
  - Build cache: enabled

![DockerHub Build Tests](imgs/dockerhub_build_tests.png)
