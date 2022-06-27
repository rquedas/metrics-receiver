package atmmetrics

import (
	"errors"

	"go.opentelemetry.io/collector/receiver/scraperhelper"
)

// Config represents the receiver config settings within the collector's config.yaml
type Config struct {
   scraperhelper.ScraperControllerSettings `mapstructure:",squash"`
   Metrics TransactionMetricsSettings `mapstructure:"metrics"`
}

type MetricSettings struct {
	Enabled bool `mapstructure:"enabled"`
}

// MetricsSettings provides settings for hostmetricsreceiver/load metrics.
type TransactionMetricsSettings struct {
	FastCashTotal MetricSettings `mapstructure:"atm.fastcash.total"`
	FastCashAverage MetricSettings `mapstructure:"atm.fastcash.average"`
}

// Validate checks if the receiver configuration is valid
func (cfg *Config) Validate() error {

	if cfg.CollectionInterval == 0 {
		return errors.New("collection_interval must be a positive duration")
	}
 
	return nil
 }