package accountsmetrics

import (
	"fmt"
	"time"

	"go.opentelemetry.io/collector/config"
)

// Config represents the receiver config settings within the collector's config.yaml
type Config struct {
   config.ReceiverSettings `mapstructure:",squash"`
   Interval       string `mapstructure:"interval"`
}


// Validate checks if the receiver configuration is valid
func (cfg *Config) Validate() error {

    interval, _ := time.ParseDuration(cfg.Interval)
	if (interval.Minutes() < 1){
		return fmt.Errorf("when defined, the interval has to be set to at least 1 minute")
	 }
 
	return nil
 }