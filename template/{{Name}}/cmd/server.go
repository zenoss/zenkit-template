{{ $pkg := print ((print (env "GOPATH") "/src/") | trimPrefix (env "PWD")) "/" Name -}}
package cmd

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tylerb/graceful"
	"{{$pkg}}/resources"
	"{{$pkg}}/resources/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware/security/jwt"
	"github.com/zenoss/zenkit"
	"github.com/zenoss/zenkit/auth"
	"github.com/zenoss/zenkit/logging"

	jwtgo "github.com/dgrijalva/jwt-go"
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

		// Read keys for verifying JWT signature
		filename := viper.GetString(zenkit.AuthKeyFileConfig)
		keys, err := auth.GetKeysFromFS(service, []string{filename})
		if err != nil {
			logrus.WithField("authfile", filename).WithError(err).Fatal("Unable to get keys for security middleware")
		}

		// Create and use dev jwt middleware
		if viper.GetBool(zenkit.AuthDisabledConfig) {
			logrus.Info("Auth Disabled, using dev jwt middleware")
			// developer claims for this app
			service.Use(auth.NewDevJWTMiddleware(jwtgo.MapClaims{
				"scopes": "api:access",
				}, jwtgo.SigningMethodHS256, keys[0]))
		}

		// create and use jwt security middleware for this app
		securityMiddleware := jwt.New(jwt.NewSimpleResolver(keys), func(h goa.Handler) goa.Handler {
			return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
				token := jwt.ContextJWT(ctx)
				if err := token.Claims.Valid(); err != nil {
					return err
				}
				return h(ctx, rw, req)
			}
		}, app.NewJWTSecurity())
		app.UseJWTMiddleware(service, securityMiddleware)

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
