package cmd

import (
	"github.com/spf13/cobra"
)

func Disable() *cobra.Command {
	cmd := &cobra.Command{
		Use: "disable",
		Short: "Stop starting marauders on startup",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return cmd
}