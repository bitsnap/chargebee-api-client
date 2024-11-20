### [Bitsnap](https://bitsnap.io) Chargebee API Client

[Bitsnap](https://bitsnap.io) Operational Intelligence Platform [Chargebee](https://www.chargebee.com/) client package.

Existing accounting platforms often suffer from reliability issues, such as potential ledger data loss and inconsistent API specifications, 
making long-term use and interoperability risky. To address these challenges, Yuriy is publishing up-to-date Accounting Services API Clients with a focus on long-term support. 
These clients aim to uncover and highlight inconsistencies in API specifications and documentation.

### Usage

```go
package main 

import (
  "time"
  
  "github.com/bitsnap/chargebee-api-client"
)

// set CHARGEBEE_API_TOKEN and CHARGEBEE_SITE variables
//
func main() {
  // or specify the respective values explicitly
  chargebee.UseToken("test_very_secure_token")
  chargebee.UseSite("test-ier-site")
  
  // Lists all customers, without pagination
  // Be mindful of API requests limits
  customers, err := chargebee.ListAllCustomers()
  
  // Lists with pagination offset
  customers, offset, err := chargebee.ListCustomers()

  // Gets next page with the previous offset
  // All entries are sorted by updated_at field, by default
  moreCustomers, newOffset, err := chargebee.ListCustomersPage(offset)
  
  // Use created_at updated_at filters to stream changes
  halfHourAgo := time.Now().Add(-30 * time.Minute)
  updatedCustomers, err := chargebee.ListCustomersUpdatedSince(&halfHourAgo)
  createdCustomers, err := chargebee.ListCustomersCreatedSince(&halfHourAgo)
}

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