package cmd

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/goadesign/goa"
	goalogrus "github.com/goadesign/goa/logging/logrus"
	"github.com/goadesign/goa/middleware"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the {{Name}} server",
	Run: func(cmd *cobra.Command, args []string) {
		service := goa.New("{{Name}}")

		logger := logrus.New()
		service.WithLogger(goalogrus.New(logger))

		service.Use(middleware.RequestID())
		service.Use(middleware.LogRequest(true))
		service.Use(middleware.ErrorHandler(service, true))
		service.Use(middleware.Recover())

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
