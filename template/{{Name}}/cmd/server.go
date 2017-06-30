{{ $pkg := print ((print (env "GOPATH") "/src/") | trimPrefix (env "PWD")) "/" Name -}}
package cmd

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/fsnotify/fsnotify"
	goalogrus "github.com/goadesign/goa/logging/logrus"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tylerb/graceful"
	"{{$pkg}}/resources"
	"{{$pkg}}/resources/app"
	"github.com/zenoss/zenkit"
)

func Logger(ctx context.Context) *logrus.Entry {
	return goalogrus.Entry(ctx)
}

var fs = afero.NewReadOnlyFs(afero.NewOsFs())

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the {{Name}} server",
	Run: func(cmd *cobra.Command, args []string) {

		// Create a new service with default middleware
		service := zenkit.NewService("{{Name}}", viper.GetBool("devmode"))

		// Set the initial log verbosity
		zenkit.SetLogLevel(service, viper.GetString("log.level"))

		// Add security
		filename := viper.GetString("auth.key_file")
		secMW, err := zenkit.JWTMiddleware(service, filename, zenkit.DefaultJWTValidation, app.NewJWTSecurity())
		if err != nil {
			logrus.WithError(err).Fatal("Unable to initialize security middleware")
		}
		app.UseJWTMiddleware(service, secMW)

		// Add tracing, if enabled
		if viper.GetBool("tracing.enabled") {
			if err := zenkit.UseXRayMiddleware(service, viper.GetString("tracing.daemon"), viper.GetInt("tracing.sample_rate")); err != nil {
				logrus.WithError(err).Fatal("Unable to initialize tracing middleware")
			}
		}

		// Start watching the config file
		go viper.WatchConfig()
		viper.OnConfigChange(func(in fsnotify.Event) {
			// Update the log verbosity
			zenkit.SetLogLevel(service, viper.GetString("log.level"))
		})

		resources.MountAllControllers(service)

		server := &graceful.Server{
			Timeout: time.Duration(15) * time.Second,
			Server: &http.Server{
				Addr:    fmt.Sprintf(":%d", viper.GetInt("http.port")),
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

	serverCmd.PersistentFlags().IntP("http-port", "p", {{Port}}, "Port to which the server should bind")
	viper.BindPFlag("http.port", serverCmd.PersistentFlags().Lookup("http-port"))
	viper.SetDefault("http.port", "{{Port}}")

	serverCmd.PersistentFlags().String("key-file", "", "File containing authentication verification key")
	viper.BindPFlag("auth.key_file", serverCmd.PersistentFlags().Lookup("key-file"))
	viper.SetDefault("auth.key_file", "")

	serverCmd.PersistentFlags().Bool("trace-enabled", false, "Whether to send trace info to AWS X-Ray")
	viper.BindPFlag("trace.enabled", serverCmd.PersistentFlags().Lookup("trace-enabled"))
	viper.SetDefault("trace.enabled", false)

	serverCmd.PersistentFlags().String("trace-daemon", "", "Address of the AWS X-Ray daemon")
	viper.BindPFlag("trace.daemon", serverCmd.PersistentFlags().Lookup("trace-daemon"))
	viper.SetDefault("trace.daemon", "")

	serverCmd.PersistentFlags().Int("trace-sample-rate", 100, "Rate at which tracing should sample requests")
	viper.BindPFlag("trace.sample_rate", serverCmd.PersistentFlags().Lookup("trace-sample-rate"))
	viper.SetDefault("trace.sample_rate", 100)

	serverCmd.PersistentFlags().Bool("dev-mode", false, "Run the daemon in dev mode, which gives anonymous requests an admin scope")
	viper.BindPFlag("devmode", serverCmd.PersistentFlags().Lookup("dev-mode"))
	viper.SetDefault("devmode", false)
}
