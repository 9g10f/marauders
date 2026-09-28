package cmd

import (
	"github.com/spf13/cobra"
)

func Enable() *cobra.Command {
	cmd := &cobra.Command{
		Use: "enable",
		Short: "Start marauders on startup",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return cmd
}