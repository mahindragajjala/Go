package content

import "fmt"

/*
Used for clean configuration of struct initializations, especially when the struct has many optional fields.
*/
// Define the struct
type Config struct {
	Host     string
	Port     int
	Timeout  int
	LogLevel string
}

// Define the Option type
type Option func(*Config)

// Option functions
func WithHost(host string) Option {
	return func(c *Config) {
		c.Host = host
	}
}

func WithPort(port int) Option {
	return func(c *Config) {
		c.Port = port
	}
}

func WithTimeout(timeout int) Option {
	return func(c *Config) {
		c.Timeout = timeout
	}
}

func WithLogLevel(level string) Option {
	return func(c *Config) {
		c.LogLevel = level
	}
}

// Constructor using functional options
func NewConfig(opts ...Option) *Config {
	// Default values
	cfg := &Config{
		Host:     "localhost",
		Port:     8080,
		Timeout:  30,
		LogLevel: "INFO",
	}
	// Apply user-defined options
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

func Functional_options_patterns() {
	cfg := NewConfig(
		WithHost("0.0.0.0"),
		WithPort(9090),
		WithLogLevel("DEBUG"),
	)

	fmt.Printf("%+v\n", cfg)
}
