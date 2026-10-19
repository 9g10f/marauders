package cmd

import (
	"os"
	"os/exec"
	"strconv"

	"github.com/spf13/cobra"
)

func Start() *cobra.Command {
	var port int
	var logLevel string // Full, none
	var sPath string
	var logFilePath string

	cmd := &cobra.Command{
		Use: "start",
		Short: "Start marauders",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			command := exec.Command(os.Args[0], "run", "--port", strconv.Itoa(port), "--s-path", sPath, "--log-file", logFilePath, "--log-level", logLevel)

			err := command.Start()
			if err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().IntVarP(&port, "port", "p", 443, "Port to serve")
	cmd.Flags().StringVarP(&logLevel, "log-level", "l", "full", "Log level")
	cmd.Flags().StringVarP(&sPath, "s-path", "s", "./s/", "S path")
	cmd.Flags().StringVarP(&logFilePath, "log-file", "f", os.DevNull, "Log file")

	return cmd
}