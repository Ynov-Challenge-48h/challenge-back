package config

import "api_test/data"

// ConfigApi initializes and returns an ApiConfig instance with host and port settings.
func ConfigApi() *data.ApiConfig {
	return &data.ApiConfig{
		Host: "localhost", // Host on which the application will run.
		Port: "8080",      // Port on which the application will run.
	}
}
