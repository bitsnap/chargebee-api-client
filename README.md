### [Bitsnap](https://bitsnap.io) Chargebee API Client

[Bitsnap](https://bitsnap.io) Operational Intelligence Platform [Chargebee](https://www.chargebee.com/) golang API client package.

### Development

Install [asdf](https://asdf-vm.com/guide/getting-started.html) with all the dependencies

```bash
git clone https://github.com/asdf-vm/asdf.git ~/.asdf --branch v0.14.0
# Add the following to ~/.bashrc:
echo '. "$HOME/.asdf/asdf.sh"' >> ~/.bashrc
# Completions must be configured by adding the following to your .bashrc:
echo '. "$HOME/.asdf/completions/asdf.bash"' >> ~/.bashrc
```

```bash
# install asdf

asdf plugin add golang
asdf plugin add golangci-lint 
asdf plugin add ko
asdf plugin add goreleaser

go install golang.org/x/tools/cmd/goimports@latest
go install mvdan.cc/gofumpt@latest
go install github.com/onsi/ginkgo/v2/ginkgo@latest
```

### Code generation

```bash
go generate chargebee/...
```

### Testing

```bash

```

## License

[Bitsnap](https://bitsnap.io) is a proprietary product developed by [Yuriy Yarosh](mailto:yuriy@yarosh.dev).

[Chargebee](https://www.chargebee.com/) golang API Client is licensed under the terms of [MIT License](LICENSE).