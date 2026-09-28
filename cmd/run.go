package cmd

import (
	"log/slog"
	"net/http"
	"strconv"

	"codeberg.org/9g10f/marauders/internal/s"
	"codeberg.org/9g10f/marauders/internal/search"
	"github.com/spf13/cobra"
)

func Run() *cobra.Command {
	var port int
	var logLevel string // Full, none
	var sPath string

	cmd := &cobra.Command{
		Use: "run",
		Short: "Run marauders",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			http.HandleFunc("/", s.Script(logLevel, sPath))
			http.HandleFunc("/search", search.Search(logLevel, sPath))

			slog.Info("marauders is running", "port", port)
			err := http.ListenAndServe(":" + strconv.Itoa(port), nil)
			if err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().IntVarP(&port, "port", "p", 443, "Port to serve")
	cmd.Flags().StringVarP(&logLevel, "log-level", "l", "full", "Log level")
	cmd.Flags().StringVarP(&sPath, "s-path", "s", "./s/", "S path")

	return cmd
}