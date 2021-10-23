package main

import (
	"html/template"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type DarkTheme struct {
	Enabled bool `yaml:"enabled"`
}

type Theme struct {
	DarkTheme *DarkTheme `yaml:"darkTheme"`
	TextColor string     `yaml:"textColor"`
	Color     string     `yaml:"color"`
}

type Contact struct {
	Enabled bool          `yaml:"enabled"`
	Body    template.HTML `yaml:"body"`
}

type Image struct {
	Enabled bool         `yaml:"enabled"`
	URL     template.URL `yaml:"url"`
}

type Author struct {
	Enabled bool         `yaml:"enabled"`
	Name    string       `yaml:"name"`
	URL     template.URL `yaml:"url"`
}

type CustomCode struct {
	Enabled bool          `yaml:"enabled"`
	CSS     template.CSS  `yaml:"css"`
	HTML    template.HTML `yaml:"html"`
}

type DefaultBackendConfiguration struct {
	PageTitle  string        `yaml:"pageTitle"`
	Title      string        `yaml:"title"`
	Body       template.HTML `yaml:"body"`
	Author     *Author       `yaml:"author"`
	Contact    *Contact      `yaml:"contact"`
	Image      *Image        `yaml:"image"`
	Theme      *Theme        `yaml:"theme"`
	CustomCode *CustomCode   `yaml:"customCode"`
}

// MakeTemplateContext unmarshals the viper configmap into the struct above
func MakeTemplateContext() *DefaultBackendConfiguration {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	config := &DefaultBackendConfiguration{}
	err := viper.Unmarshal(config)
	if err != nil {
		logger.Fatal("Failed to unmarshal configuration", zap.Error(err))
	}

	return config
}

// BindEnvVariables maps each top-level ENV variable to a field in the config struct
func BindEnvVariables() {
	viper.BindEnv("staticAssetsPath", "STATIC_ASSETS_PATH")
	viper.BindEnv("staticAssetsPrefix", "STATIC_ASSETS_PREFIX")
	viper.BindEnv("title", "PAGE_TITLE")
	viper.BindEnv("title", "TITLE")
	viper.BindEnv("image.enabled", "IMAGE_ENABLED")
	viper.BindEnv("image.url", "IMAGE_URL")
	viper.BindEnv("body", "BODY")
	viper.BindEnv("contact.enabled", "CONTACT_ENABLED")
	viper.BindEnv("contact.body", "CONTACT_BODY")
	viper.BindEnv("author.enabled", "AUTHOR_ENABLED")
	viper.BindEnv("author.name", "AUTHOR_NAME")
	viper.BindEnv("author.url", "AUTHOR_URL")
	viper.BindEnv("theme.textColor", "TEXT_COLOR")
	viper.BindEnv("theme.color", "BACKGROUND_COLOR")
	viper.BindEnv("theme.darkTheme.enabled", "DARK_THEME_ENABLED")
	viper.BindEnv("customCode.enabled", "CUSTOM_CODE")
	viper.BindEnv("customCode.css", "CUSTOM_CSS")
	viper.BindEnv("customCode.html", "CUSTOM_HTML")
}
