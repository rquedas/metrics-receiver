package accountsmetrics

import (
	"fmt"
	"time"
)

// Config represents the receiver config settings within the collector's config.yaml
type Config struct {
	Interval string `mapstructure:"interval"`
}

// Validate checks if the receiver configuration is valid
func (cfg *Config) Validate() error {

	interval, _ := time.ParseDuration(cfg.Interval)
	if interval.Seconds() < 10 {
		return fmt.Errorf("when defined, the interval has to be set to at least 10 seconds")
	}

	return nil
}
