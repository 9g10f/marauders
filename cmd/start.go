package cmd

import (
	"github.com/spf13/cobra"
)

func Start() *cobra.Command {
	cmd := &cobra.Command{
		Use: "start",
		Short: "Start marauders",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return cmd
}