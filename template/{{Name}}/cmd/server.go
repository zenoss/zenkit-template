{{ $pkg := print ((print (env "GOPATH") "/src/") | trimPrefix (env "PWD")) "/" Name -}}
package cmd

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/cenkalti/backoff"
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

		// Get the secret key
		var key []byte
		filename := viper.GetString("auth.key_file")
		readKey := func() error {
			data, err := afero.ReadFile(fs, filename)
			if err != nil {
					logrus.WithError(err).WithField("keyfile", filename).Info("Unable to load auth key. Retrying.")
					return err
			}
			key = data
			return nil
		}
		// Docker sometimes doesn't mount the secret right away, so we'll do a short retry
		boff := backoff.NewExponentialBackOff()
		boff.MaxElapsedTime = 10 * time.Second
		if err := backoff.Retry(readKey, boff); err != nil {
			logrus.WithError(err).Fatal("Unable to load auth verification key")
		}

		// Create a new service with default middleware
		service := zenkit.NewService("{{Name}}")

		// Set the initial log verbosity
		zenkit.SetLogLevel(service, viper.GetString("log.level"))

		// Add security
		secMW := zenkit.JWTMiddleware(key, zenkit.DefaultJWTValidation, app.NewJWTSecurity())
		app.UseJWTMiddleware(service, secMW)

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
}
