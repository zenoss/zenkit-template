{{ $pkg := print ((print (env "GOPATH") "/src/") | trimPrefix (env "PWD")) "/" Name -}}
package cmd

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tylerb/graceful"
	"{{$pkg}}/resources"
	"{{$pkg}}/resources/app"
	"github.com/zenoss/zenkit"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the {{Name}} server",
	Run: func(cmd *cobra.Command, args []string) {

		// Create a new service with default middleware
		service := zenkit.NewService("{{Name}}", viper.GetBool(zenkit.AuthEnabledConfig))

		// Set the initial log verbosity
		zenkit.SetLogLevel(service, viper.GetString(zenkit.LogLevelConfig))

		// Add security
		filename := viper.GetString(zenkit.AuthKeyFileConfig)
		secMW, err := zenkit.JWTMiddleware(service, filename, zenkit.DefaultJWTValidation, app.NewJWTSecurity())
		if err != nil {
			logrus.WithError(err).Fatal("Unable to initialize security middleware")
		}
		app.UseJWTMiddleware(service, secMW)

		// Add tracing, if enabled
		if viper.GetBool(zenkit.TracingEnabledConfig) {
			if err := zenkit.UseXRayMiddleware(service, viper.GetString(zenkit.TracingDaemonConfig), viper.GetInt(zenkit.TracingSampleRateConfig)); err != nil {
				logrus.WithError(err).Fatal("Unable to initialize tracing middleware")
			}
		}

		// Start watching the config file
		go viper.WatchConfig()
		viper.OnConfigChange(func(in fsnotify.Event) {
			// Update the log verbosity
			zenkit.SetLogLevel(service, viper.GetString(zenkit.LogLevelConfig))
		})

		resources.MountAllControllers(service)

		server := &graceful.Server{
			Timeout: time.Duration(15) * time.Second,
			Server: &http.Server{
				Addr:    fmt.Sprintf(":%d", viper.GetInt(zenkit.HTTPPortConfig)),
				Handler: service.Mux,
			},
		}

		go func() {
			if err := server.ListenAndServe(); err != nil {
				logrus.WithError(err).Fatal("Server shut down")
			}
		}()
		logrus.WithField("address", server.Addr).Info("Server started")

		// Wait for the server to exit
		<-server.StopChan()
		logrus.Info("Goodbye")
	},
}

func init() {
	RootCmd.AddCommand(serverCmd)
	zenkit.AddStandardServerOptions(serverCmd, {{Port}})
}
