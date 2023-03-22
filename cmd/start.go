package cmd

import (
	"flush-tsdb/internal/pkg/tsdb"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"flush-tsdb/internal/app"
)

var (
	startCmd = &cobra.Command{
		Use:   "start",
		Short: "",
		Run: func(cmd *cobra.Command, args []string) {
			app.Server(log.WithField("start", "server"))
		},
	}
)

// flag
func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.PersistentFlags().StringVarP(&tsdb.DSN, "tsdbdsn", "d", "", "tsdb dsn, format: '<user>:<password>@tcp(<host>:<port>)/'")
}
