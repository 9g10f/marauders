package cmd

import (
	"github.com/spf13/cobra"
)

var Root = &cobra.Command{
	Use: "marauders",
	Version: Version,
	Short: "The CLI server for Marauder Launcher",
	CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
}