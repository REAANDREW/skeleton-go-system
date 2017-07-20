package main

// Configuration for the Appication.
type Configuration struct {
	LogLevel string `json:"log_level"`
}

func DefaultConfiguration() Configuration {
	return Configuration{
		LogLevel: "INFO",
	}
}
