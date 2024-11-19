### [Bitsnap](https://bitsnap.io) Chargebee API Client

[Bitsnap](https://bitsnap.io) Operational Intelligence Platform [Chargebee](https://www.chargebee.com/) golang API client package.

Existing accounting platforms deemed to be unreliable - potential ledger data loss with inconsistent API specs made long-term use and interoperability quite a gamble.
Yuriy is publishing an up-to-date Accounting services API Clients with long term support in mind, to uncover and point out API specs and documentation inconsistencies. 
There are no o

### Usage

```go

```

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
# will scrape the existing API documentation pages and emit the respective API client methods
# enums and models validation, where applicable

go generate ./...
```

### Testing

```bash
# trial Chargebee sites are suffixed with *-test 
# make sure to double check - it was not documented properly

CHARGEBEE_API_TOKEN=test_xxx CHARGEBEE_SITE=xxx-test ginkgo run ./...
```

### TODO
 - [ ] refactor scraping
 - [ ] maybe introduce some attr filtering
 - [ ] implement webhook API endpoints for CDC streaming
 - [ ] provide [Apache Flink](https://flink.apache.org/) [JNI source](https://nightlies.apache.org/flink/flink-docs-master/docs/dev/datastream/sources/) binding  
 - [ ] provide [Argo Events](https://argoproj.github.io/argo-events/) source binding

## License

[Bitsnap](https://bitsnap.io) is a proprietary product developed by [Yuriy Yarosh](mailto:yuriy@yarosh.dev).

[Chargebee](https://www.chargebee.com/) golang API Client is licensed under the terms of [MIT License](LICENSE).