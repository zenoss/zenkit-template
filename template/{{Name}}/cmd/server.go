{{ $pkg := print ((print (env "GOPATH") "/src/") | trimPrefix (env "PWD")) "/" Name -}}
package cmd

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tylerb/graceful"
	"{{$pkg}}/resources"
	"github.com/zenoss/zenkit"
	"github.com/zenoss/zenkit/logging"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the {{Name}} server",
	Run: func(cmd *cobra.Command, args []string) {

		// Create a new service with default middleware
		service := zenkit.NewService("{{Name}}")

		// Set the initial log verbosity
		logging.SetLogLevel(service, viper.GetString(zenkit.LogLevelConfig))

		// Add your security middleware here
		// if viper.GetBool(zenkit.AuthDisabledConfig) {
		// 	logrus.Info("Auth Disabled, using MyDevJWTMiddleware")
		// 	service.Use(MyDevJWTMiddleware)
		// }
		// app.UseJWTMiddleware(service, MySecurityMiddleware)

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
			v := viper.New()
			v.SetConfigFile(in.Name)
			if err := v.ReadInConfig(); err != nil {
				logrus.WithField("configfile", in.Name).WithError(err).Error("Could not read config file, log level not updated")
				return
			}
			logging.SetLogLevel(service, v.GetString(zenkit.LogLevelConfig))
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

		// Register health checks
		registerPing(viper.GetInt(zenkit.AdminPortConfig), time.Second, 15 * time.Second)

		// Start the admin service
		adminService := zenkit.NewAdminService(service)
		adminServer := &graceful.Server{
			Timeout: time.Duration(15) * time.Second,
			Server: &http.Server{
				Addr: fmt.Sprintf(":%d", viper.GetInt(zenkit.AdminPortConfig)),
				Handler: adminService.Mux,
			},
		}

		go func() {
			if err := adminServer.ListenAndServe(); err != nil {
				logrus.WithError(err).Fatal("Admin server shut down")
			}
		}()

		// Wait for the server to exit
		<-server.StopChan()
		<-adminServer.StopChan()
		logrus.Info("Goodbye")
	},
}

func init() {
	RootCmd.AddCommand(serverCmd)
	zenkit.AddStandardServerOptions(serverCmd, {{Port}}, {{AdminPort}})
}
