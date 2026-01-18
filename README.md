<div align="center">

<img alt="OverScry Logo" height="25%" width="25%" src="https://raw.githubusercontent.com/kkrypt0nn/OverScry/refs/heads/main/assets/logo.png" />

# OverScry

[![Discord Server Badge](https://img.shields.io/discord/1358456011316396295?logo=discord)](https://discord.gg/xj6y5ZaTMr)
[![Version Badge](https://img.shields.io/github/release/kkrypt0nn/overscry.svg)](https://github.com/kkrypt0nn/OverScry/releases)
[![Docker Badge](https://img.shields.io/docker/v/kkrypt0nn/overscry?logo=docker)](https://hub.docker.com/r/kkrypt0nn/overscry)

[![CI Badge](https://github.com/kkrypt0nn/OverScry/actions/workflows/ci.yml/badge.svg)](https://github.com/kkrypt0nn/OverScry/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/kkrypt0nn/overscry)](https://goreportcard.com/report/github.com/kkrypt0nn/overscry)
[![Last Commit Badge](https://img.shields.io/github/last-commit/kkrypt0nn/OverScry)](https://github.com/kkrypt0nn/OverScry/commits/main)

</div>

> [!NOTE]
> I've put the project **on hold** for now, I will work on it again once I have more time.

> [!CAUTION]
> This is a WIP tool that is **very unstable** and **not fully optimised**, use at your own care! This README will also be reworked.

### 🔍 The magic lens for Overpass queries

OverScry is a tool designed to simplify the process of generating Overpass queries from a YML structure file. By taking a structured YML input that defines specific geographical data and search parameters, OverScry automatically converts it into a valid Overpass query. Overall it makes it easier for both beginners and experts to leverage Overpass API capabilities without needing to understand how to write complex raw queries.

## Getting Started

> [!TIP]
> There is a **web version** of the tool available, no installation required, just your browser.
> Check it out [here](https://overscry.krypton.ninja/online/)

### Installation

Installing the tool can currently only be done via Go and Docker.

#### Go

You need to have [Go](https://go.dev/dl/) installed. You can then install using:

```bash
go install github.com/kkrypt0nn/overscry@latest
```

#### Docker

You can run the tool from the published [Docker image](https://hub.docker.com/r/kkrypt0nn/overscry) using:

```bash
docker run -v .:/data -it kkrypt0nn/overscry gen --settings /data/_examples/housenumber.yml
```

### Example Usage

After installing you can run the tool with

```bash
overscry --help
```

This will give you the list of commands and their respective flags.

You can then simply generate a query with

```bash
overscry gen --settings _examples/housenumber.yml
```

## Settings YML File

The YML file for the settings has the following example structure

```yml
version: "0.0.7-dev"
author: "Krypton (@kkrypt0nn)"
description: "A query to get apartments with 2 levels"
node:
  building:
    value: apartments
    levels:
      value: 2
```

More details about the supported fields is available [here](https://overscry.krypton.ninja/settings/introduction.html).


> [!NOTE]
> A lot of work is put into so that more arguments and features/tags are supported. Please be patient for upcoming changes.

## Supported Platforms

OverScry works on the following platforms:

- **Linux**
- **macOS**
- **Windows**
- **Browser (via WebAssembly)** - See https://overscry.krypton.ninja/online/
- **Docker**
- **Android**

## Troubleshooting

If you encounter issues while using OverScry, consider the following:

- Ensure you are running the latest version
- Join my Discord server [here](https://discord.gg/xj6y5ZaTMr)
- Use the [GitHub issue tracker](https://github.com/kkrypt0nn/OverScry/issues)

## Contributing

People may contribute by following the [Contributing Guidelines](./CONTRIBUTING.md) and
the [Code of Conduct](./CODE_OF_CONDUCT.md) once the tool is out of pre-releases, so once `v0.1.0` comes out.

## License

This tool was made with 💜 by Krypton and is under the [MIT License](./LICENSE.md).
