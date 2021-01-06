package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/errata-ai/vale/v2/internal/core"
)

var commandInfo = map[string]string{
	"ls-config": "Print the current configuration to stdout and exit.",
}

var actions = map[string]func(args []string, cfg *core.Config) error{
	"ls-config": printConfig,
	"dc":        printConfig,
	"help":      printUsage,
}

func printConfig(args []string, cfg *core.Config) error {
	cfg, err := core.NewConfig(&flags)
	if err != nil {
		ShowError(err, flags.Output, os.Stderr)
	}

	err = core.From("ini", cfg)
	if err != nil {
		ShowError(err, flags.Output, os.Stderr)
	}

	fmt.Println(cfg.String())
	return err
}

func printUsage(args []string, cfg *core.Config) error {
	flag.Usage()
	return nil
}
