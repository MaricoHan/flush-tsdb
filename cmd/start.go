package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"flush-tsdb/internal/app"
	"flush-tsdb/internal/pkg/configs"
)

var (
	cfgFile string

	startCmd = &cobra.Command{
		Use:   "start",
		Short: "Start admin data server",
		Run: func(cmd *cobra.Command, args []string) {
			if err := configs.NewConfigs().Read(cmd, cfgFile); err != nil {
				log.WithField("start", "server").Errorf(err.Error())
				return
			}
			app.Server(log.WithField("start", "server"))
		},
	}
)

// flag
func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "configs file (default is app/configs/configs.yaml)")
}
