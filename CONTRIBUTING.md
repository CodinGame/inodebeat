# Contributing to Inodebeat

Welcome to Inodebeat.

This beat has been created thanks to the official [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).

## Getting Started with Inodebeat

### Requirements

* [Golang](https://golang.org/dl/) >= 1.7
* [Glide](https://glide.sh/)
* Virtualenv

### Build

Ensure that this folder is at the following location: `${GOPATH}/github.com/codingame/inodebeat`

To build the binary for Inodebeat run the command below. This will generate a binary
in the same directory with the name inodebeat.

```
# Install dependencies and generate config/template files
make setup

# Create the inodebeat binary
make
```

### Run

To run Inodebeat with debugging output enabled, run:

```
./inodebeat -c inodebeat.yml -e -d "*"
```

### Update

Each beat has a template for the mapping in elasticsearch and a documentation for the fields
which is automatically generated based on `_meta/fields.yml`.
To generate etc/inodebeat.template.json and etc/inodebeat.asciidoc

```
make update
```

### Cleanup

To clean up the build directory and generated artifacts, run:

```
make clean
```

## Packaging

To create an `inodebeat` Docker image, run:

```
docker build -t inodebeat .
```

You can run it with:

```
docker run inodebeat
```

The official Docker images for `codingame/inodebeat` are built by [Docker Hub](https://hub.docker.com/r/codingame/inodebeat/).
