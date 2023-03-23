package cmd

import (
	"flush-tsdb/internal/pkg/tsdb"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"flush-tsdb/internal/app"
)

var flushTagCmd = &cobra.Command{
	Use:   "flush-tag",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		app.FlushTag(logrus.WithField("service", "flush-tag"))
	},
}

func init() {
	rootCmd.AddCommand(flushTagCmd)

	flushTagCmd.PersistentFlags().StringVarP(&tsdb.DSN, "tsdbdsn", "d", "", "tsdb dsn, format: '<user>:<password>@tcp(<host>:<port>)/'")
}
