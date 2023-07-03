package test

import (
	"context"
	"fmt"
	"github.com/orlangure/gnomock"
)

const (
	defaultPort    = 8080
	defaultVersion = "latest"
)

var defaultVars = map[string]string{
	"AUTHENTICATION_ANONYMOUS_ACCESS_ENABLED": "true",
	"DEFAULT_VECTORIZER_MODULE":               "none",
}

type WeaviatePreset struct {
	Version string            `json:"version"`
	EnvVars map[string]string `json:"envVars"`
}

func (w *WeaviatePreset) Image() string {
	return fmt.Sprintf("semitechnologies/weaviate:%s", w.Version)
}

func (w *WeaviatePreset) Ports() gnomock.NamedPorts {
	return gnomock.DefaultTCP(defaultPort)
}

func (w *WeaviatePreset) Options() []gnomock.Option {
	w.setDefaults()
	opts := []gnomock.Option{
		gnomock.WithHealthCheck(w.healthcheck),
	}
	for name, value := range w.EnvVars {
		opts = append(opts, gnomock.WithEnv(fmt.Sprintf("%s=%s", name, value)))
	}
	return opts
}

func (w *WeaviatePreset) healthcheck(_ context.Context, c *gnomock.Container) error {
	return nil
}

func (w *WeaviatePreset) setDefaults() {
	if w.Version == "" {
		w.Version = defaultVersion
	}
}
func Preset(opts ...Option) gnomock.Preset {
	preset := &WeaviatePreset{}
	for _, opt := range opts {
		opt(preset)
	}

	return preset
}
