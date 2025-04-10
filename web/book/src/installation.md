# Installation

Installing the tool can currently only be done via Go and Docker.

## Go

You need to have [Go](https://go.dev/dl/) installed. You can then install using:

```bash
go install github.com/kkrypt0nn/overscry@latest
```

You can then generate a query with the basic examples in the [examples](_examples) folder with

```bash
overscry gen --settings _examples/housenumber.yml
```

## Docker

You can run the tool from the published [Docker image](https://hub.docker.com/r/kkrypt0nn/overscry) using:

```bash
docker run -v .:/data -it kkrypt0nn/overscry gen --settings /data/_examples/housenumber.yml
```
