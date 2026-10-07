package main

import (
	"github.com/heliumcode-labs/helium/cmd"
	"github.com/heliumcode-labs/helium/internal/logging"
)

func main() {
	defer logging.RecoverPanic("main", func() {
		logging.ErrorPersist("Application terminated due to unhandled panic")
	})

	cmd.Execute()
}
