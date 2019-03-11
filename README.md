[![Build Status](https://travis-ci.org/meyermarcel/icm.svg "Build Status")](https://travis-ci.org/meyermarcel/icm)

# Validate or generate intermodal container markings

![Demo](docs/gif/demo.gif)

## Documentation

See [docs](docs/icm.md)

## Installation

### macOS

```
brew install meyermarcel/tap/icm
```

### Linux & Windows

Download your binary in the [Releases](https://github.com/meyermarcel/icm/releases) section.

See the [misc command](docs/icm_misc.md) for **bash/zsh completions** and **manual pages**.

Pull requests to publish automatically to Scoop and Snapcraft in [Goreleaser](https://goreleaser.com/customization) are welcomed.

## Contributing

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
    * [Golang 1.12.x](https://golang.org/doc/install)
    * [Go Dep 0.5.x](https://golang.github.io/dep/docs/installation.html)
    * [GolangCI-Lint 1.x.x](https://github.com/golangci/golangci-lint#install)
    * [GNU Make 4.2.x](https://www.gnu.org/software/make/)

1. To build project execute
    ```
    make
    ```

## Release

1. Create version tag according to [SemVer](https://semver.org)
    ```
    git tag "x.y.z"
    ```

1. Push tag and let Travis CI and Goreleaser do the work
+
    ```
    git push --tags
    ```

## License

icm is released under the Apache 2.0 license. See [LICENSE](https://github.com/meyermarcel/icm/blob/master/LICENSE)
