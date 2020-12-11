# Continuous integration

For this project I've used Travis CI and GitHub Actions. The different configuration in both systems done can be seen below.

## Travis CI

To setup the CI in Travis the first step is to create an account in [travis-ci.com/](https://travis-ci.com/). I already have an account (linked with GitHub) so I ommited this step. After that, I've allowed Travis to access all the repositories in my GitHub account as show in the picture below:

![travis_gh](./imgs/travis_gh.png)

After that I've created the following `.travis.yml` file in the root of the project:

```yml
language: go

go:
  - 1.15

services:
  - mongodb

script:
  - make test
```

Its content is super simple. First of all I've set the language to test and the version to use, in this case `Golang v1.15`. Since the software makes use of a database connection to a `MongoDB` server, in order to run the unit tests related to that code I've set the `mongodb` service to be used in every build. Finally, I run the testing process by using the task manager.

The latest Travis builds for HarvestCCode can be seen [here](https://travis-ci.com/github/harvestcore/HarvestCCode/builds) and the status badge can be seen in the [README](../../README.md) of this repo.

## GitHub Actions

My main CI system is GitHub Actions, and I've configured three pipelines that do various things.

### Basic setup

To setup GitHub actions the only procedure needed is to create a `.github/workflows` folder in the root of the repo.

### Go Lint

This first pipeline runs a linter that analyses the code to check for stylistic, syntactic and any other static code error.

The content of that file is the following:

```yml
name: Go linter

on:
  push:
    paths:
      - '.github/**'
      - 'src/**'
  pull_request:
    branches:
      - master

jobs:
  build:
    runs-on: ubuntu-latest
    env:
      GOPATH: ${{ github.workspace }}
    defaults:
      run:
        working-directory: ${{ env.GOPATH }}/src/github.com/${{ github.repository }}
    
    steps:
    - name: Git checkout
      uses: actions/checkout@v2
      with:
        path: ${{ env.GOPATH }}/src/github.com/${{ github.repository }}

    - name: Setup Go v1.15.4
      uses: actions/setup-go@v2
      with:
        go-version: '1.15.4'

    - name: Go linter
      run: make lint
```

This pipeline is only run if its content changes or the source code of the software changes. First of all setups the Golang environment and then runs the linter using the [task manager](../../Makefile) (`make`).

```yml
lint: deps testdeps
    go vet ./src...
```

### Unit tests

This second pipeline runs the unit tests.

The content of that file is the following:

```yml
name: Tests

on:
  push:
    paths:
      - '.github/**'
      - 'src/**'
  pull_request:
    branches:
      - master

jobs:
  build:
    runs-on: ubuntu-latest
    env:
      GOPATH: ${{ github.workspace }}
    defaults:
      run:
        working-directory: ${{ env.GOPATH }}/src/github.com/${{ github.repository }}
    
    steps:
    - name: Git checkout
      uses: actions/checkout@v2
      with:
        path: ${{ env.GOPATH }}/src/github.com/${{ github.repository }}

    - name: Setup Go v1.15.4
      uses: actions/setup-go@v2
      with:
        go-version: '1.15.4'

    - name: Run mongoDB
      uses: wbari/start-mongoDB@v0.2
      with:
        mongoDBVersion: 4.4.2

    - name: Run tests
      run: make test

```

Its content is pretty similar to the linter one, but in this case a new step is added. As mentioned in the Travis CI pipeline: _Since the software makes use of a database connection to a `MongoDB` server, in order to run the unit tests related to that code I've set the `mongodb` service to be used in every build. Finally, I run the testing process by using the task manager._

```yml
test: testdeps
    go test ./src... -v
```

### Testing Docker image

This last pipeline builds and pushes a Docker image that can be used to run the unit tests.

More info about it [here](../dockerf.tests.md).

This image could have been used in the Travis CI pipeline for example, but I have put that idea aside for a while since I found a bunch of problems/errors when trying to setup the MongoDB service and the connection with the Docker image. I'll try to get it working in the future, though.
