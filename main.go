package main

import (
	"github.com/bitsnap/chargebee-api-client/codegen"
	. "github.com/bitsnap/go-util"
)

func main() {
	dir := TargetCodegenDir()
	syncLogger := SetGlobalLogger()
	defer syncLogger()

	Logger().Infow("starting bitsnap chargebee codegen",
		"dir", dir,
	)

	codegen.GenerateInto(dir)
}
