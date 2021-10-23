package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var config *DefaultBackendConfiguration
var logger *zap.Logger
var templates *template.Template

// probeHandler responds to external healthchecks
func probeHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info("probe handler", zap.String("path", r.URL.Path))
	w.WriteHeader(200)
	fmt.Fprint(w, "Ok")
}

// rootHandler templates the HTML with the provided configuration
func rootHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info("rootHandler", zap.String("path", r.URL.Path))
	w.WriteHeader(config.StatusCode)
	renderTemplate(w, "index", config)
}

// renderTemplate renders a (cached) template given as string
func renderTemplate(w http.ResponseWriter, tmpl string, c *DefaultBackendConfiguration) {
	err := templates.ExecuteTemplate(w, tmpl+".html.tmpl", c)
	if err != nil {
		logger.Error("Executing template failed", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	logger, _ = zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any

	// Load up viper for configuration
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AddConfigPath("/")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		logger.Fatal("Failed to load system-defined configuration!", zap.Error(err))
	}
	logger.Info("Loaded system-defined configuration", zap.Any("config", viper.AllSettings()))

	// user-specified configmap from a volume mount to /etc/default-backend
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("/etc/default-backend/")
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logger.Info("No user-defined configuration found in /etc/default-backend, skipping...")
		} else {
			logger.Fatal("Failed to load user-defined configmap")
		}
	}

	BindEnvVariables()

	viper.MergeConfigMap(v.AllSettings())
	logger.Info("Merged system-defined and user-defined configmaps", zap.Any("config", viper.AllSettings()))
	// fill in context for template rendering
	config = MakeTemplateContext()
	logger.Info("Created template context", zap.Any("context", config))

	// Parse and cache template
	cwd, err := os.Getwd()
	if err != nil {
		logger.Fatal("Failed to get current working directory, running off of virtual filesystem?", zap.Error(err))
	}
	templates = template.Must(template.ParseFiles(cwd + "/tmpl/index.html.tmpl"))

	// serve static assets from /public
	staticPath := viper.GetString("staticAssetsPath")
	prefix := viper.GetString("staticAssetsPrefix")
	http.Handle(prefix, http.StripPrefix(prefix, (http.FileServer(http.Dir(staticPath)))))

	// healthcheck endpoints
	http.HandleFunc("/healthz", probeHandler)
	http.HandleFunc("/readyz", probeHandler)
	http.HandleFunc("/livez", probeHandler)

	// root page handler
	http.HandleFunc("/", rootHandler)

	logger.Info("Registered endpoints")
	logger.Info("Listening on port 8080")
	logger.Fatal("Listening on 8080 failed, maybe there's already something bound?", zap.Error(http.ListenAndServe(":8080", nil)))
}
