# CircleOps - CircleCI in the Terminal

**_This project is still in early development_**

CircleOps is a project that provides a command line tool - `circlectl` - for interacting with the CircleCI API. It is currently in early development.

## Installation 

`go get -u github.com/jaymickey/circleops/cmd/circlectl`

Running `circlectl setup` will prompt for a server URL and API token. Configuration is stored by default in `$HOME/.circlectl/config.yaml`.

## Development

This project uses Go Modules, therefore the repo can be cloned into any location of the filesystem, and doesn't require `$GOPATH`. However, it does require **at least Go v1.11**.