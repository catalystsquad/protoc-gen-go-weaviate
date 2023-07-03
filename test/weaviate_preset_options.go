package test

type Option func(preset *WeaviatePreset)

func WithVersion(version string) Option {
	return func(o *WeaviatePreset) {
		o.Version = version
	}
}

func WithEnvVars(envVars map[string]string) Option {
	return func(o *WeaviatePreset) {
		o.EnvVars = envVars
	}
}
