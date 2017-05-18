package cmd

import (
	"context"
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/fsnotify/fsnotify"
	goalogrus "github.com/goadesign/goa/logging/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zenoss/zenkit"
)

func Logger(ctx context.Context) *logrus.Entry {
	return goalogrus.Entry(ctx)
}

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the {{Name}} server",
	Run: func(cmd *cobra.Command, args []string) {

		// Create a new service with default middleware
		service := zenkit.NewService("{{Name}}")

		// Set the initial log verbosity
		zenkit.SetVerbosity(service, verbosity)

		// Start watching the config file
		go viper.WatchConfig()
		viper.OnConfigChange(func(in fsnotify.Event) {
			// Update the log verbosity
			zenkit.SetVerbosity(service, verbosity)
		})

		if err := service.ListenAndServe(fmt.Sprintf(":%d", viper.GetInt("port"))); err != nil {
			service.LogError("startup", "err", err)
		}
	},
}

func init() {
	RootCmd.AddCommand(serverCmd)

	serverCmd.PersistentFlags().IntP("port", "p", 8080, "Port to which the server should bind")
	viper.BindPFlag("port", serverCmd.PersistentFlags().Lookup("port"))
	viper.SetDefault("port", "{{Port}}")
}
