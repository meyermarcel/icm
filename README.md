[![build](https://github.com/meyermarcel/icm/actions/workflows/build.yml/badge.svg)](https://github.com/meyermarcel/icm/actions/workflows/build.yml)  [![Go Report Card](https://goreportcard.com/badge/github.com/meyermarcel/icm)](https://goreportcard.com/report/github.com/meyermarcel/icm)

# icm (intermodal container markings)

icm generates or validates single data or whole data sets of intermodal container markings according to [ISO 6346](https://en.wikipedia.org/wiki/ISO_6346).

See examples for [`generate` command](docs/icm_generate.md#examples) and [`validate` command](docs/icm_validate.md#examples).

## Demo

![Demo](docs/gif/demo.gif)

## Installation

### macOS with [Homebrew](https://brew.sh)

```
brew install meyermarcel/tap/icm
```

### Linux with [Homebrew on Linux](https://docs.brew.sh/Homebrew-on-Linux)

```
brew install meyermarcel/tap/icm
```

### Windows with [Scoop](https://scoop.sh)

```
scoop bucket add meyermarcel-bucket https://github.com/meyermarcel/scoop-bucket.git
scoop install icm
```

### Manual

Download your binary in the [Releases](https://github.com/meyermarcel/icm/releases) section.

See the [`completion` command](docs/icm_completion.md) for **bash**, **zsh**, **fish** and **powershell** completions.

See the [`doc` command](docs/icm_doc.md) for **manual pages** and **markdown**.

## Contribution

1. Fork it

1. Download your fork
    ```
    git clone https://github.com/github_username/icm && cd icm
    ```

1. Create your feature branch
    ```
    git checkout -b my-new-feature
    ```

1. Make changes and add them
    ```
    git add .
    ```

1. Commit your changes
    ```
    git commit -m 'Add some feature'
    ```

1. Push to the branch
    ```
    git push origin my-new-feature
    ```

1. Create new pull request

## Development

1. Requirements
    * [Golang 1.21.x](https://golang.org/doc/install)
    * [golangci-lint latest version](https://github.com/golangci/golangci-lint#install)
    * [GNU Make 4.x.x](https://www.gnu.org/software/make/)
    * gufumpt -> `go install mvdan.cc/gofumpt@latest`

1. To build project execute
    ```
    make
    ```

    ---
    See all available `make` targets

    ```
    make <TAB> <TAB>
    ```

## Release

1. Dry run with `goreleaser`
    ```
    goreleaser release --clean --skip=validate --skip=publish
    ```

1. Create version tag according to [SemVer](https://semver.org)
    ```
    git tag 'x.y.z'
    ```

1. Push tag and let GitHub Actions and Goreleaser do the work
    ```
    git push --tags
    ```

## License

icm is released under the MIT license. See [LICENSE](https://github.com/meyermarcel/icm/blob/master/LICENSE)
