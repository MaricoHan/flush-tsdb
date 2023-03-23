package cmd

import (
	"flush-tsdb/internal/pkg/tsdb"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"flush-tsdb/internal/app"
)

var GenerateTestDataAmount int

var genTestDataCmd = &cobra.Command{
	Use:   "gen-test-data",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		app.GenerateTestData(GenerateTestDataAmount, logrus.WithField("service", "gen-test-data"))
	},
}

func init() {
	rootCmd.AddCommand(genTestDataCmd)

	genTestDataCmd.PersistentFlags().StringVarP(&tsdb.DSN, "tsdbdsn", "d", "", "tsdb dsn, format: '<user>:<password>@tcp(<host>:<port>)/'")
	genTestDataCmd.PersistentFlags().IntVarP(&GenerateTestDataAmount, "n", "n", 10000, "data generated amount")
}
