package timber

import (
	"fmt"
	"os"
	"time"
)

const (
	defaultEnvAPIKey   = "TIMBER_API_KEY"
	defaultEnvSourceID = "TIMBER_SOURCE_ID"

	defaultEndpoint           = "https://logs.timber.io"
	defaultTimeoput           = time.Second * 30
	defaultCheckpointSize     = 64
	defaultCheckpointInterval = time.Second * 2
)

var (
	envAPIKey   string
	envSourceID string
)

func init() {
	envAPIKey = os.Getenv(defaultEnvAPIKey)
	envSourceID = os.Getenv(defaultEnvSourceID)
}

// Config contains parameters for Timber.io.
type Config struct {
	APIKey   string
	SourceID string

	// Log context
	Environment string
	Hostname    string

	MinimumLevel   string
	Sync           bool
	Debug          bool
	NoRetry        bool
	Timeout        time.Duration
	CustomEndpoint string

	// Daemon configurations
	CheckpointSize     int
	CheckpointInterval time.Duration

	// used internally
	url      string
	apikey   string
	hostname string
	pid      int
	timeout  time.Duration
}

// Init intialize a config.
func (c *Config) Init() {
	c.url = c.getURL()
	c.apikey = c.getAPIKey()

	c.hostname = c.getHostname()
	c.pid = os.Getpid()
	c.timeout = c.getTimeout()
}

func (c Config) getAPIKey() string {
	if c.APIKey != "" {
		return c.APIKey
	}
	return envAPIKey
}

func (c Config) getURL() string {
	endpoint := c.CustomEndpoint
	if endpoint == "" {
		endpoint = defaultEndpoint
	}

	sourceID := c.SourceID
	if sourceID == "" {
		sourceID = envSourceID
	}

	switch {
	case sourceID != "":
		return fmt.Sprintf("%s/sources/%s/frames", endpoint, sourceID)
	default:
		return fmt.Sprintf("%s/frames", endpoint)
	}
}

func (c Config) getHostname() string {
	if c.Hostname != "" {
		return c.Hostname
	}

	name, _ := os.Hostname()
	return name
}

func (c Config) getTimeout() time.Duration {
	if c.Timeout > 0 {
		return c.Timeout
	}
	return defaultTimeoput
}

func (c Config) getCheckpointSize() int {
	if c.CheckpointSize > 0 {
		return c.CheckpointSize
	}
	return defaultCheckpointSize
}

func (c Config) getCheckpointInterval() time.Duration {
	if c.CheckpointInterval > 0 {
		return c.CheckpointInterval
	}
	return defaultCheckpointInterval
}
