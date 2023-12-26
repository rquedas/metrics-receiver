package atmmetrics

import (
	"errors"
	"time"

	"go.opentelemetry.io/collector/receiver/scraperhelper"
)

// Config represents the receiver config settings within the collector's config.yaml
type Config struct {
	scraperhelper.ScraperControllerSettings `mapstructure:",squash"`
	Metrics                                 TransactionMetricsSettings `mapstructure:"metrics"`
}

// MetricsSettings provides common settings that will be used across all metrics.
type MetricSettings struct {
	Enabled bool `mapstructure:"enabled"`
}

// TransactionMetricsSettings  provides settings for every individual metric
type TransactionMetricsSettings struct {
	FastCashTotal   MetricSettings `mapstructure:"atm.fastcash.total"`
	FastCashAverage MetricSettings `mapstructure:"atm.fastcash.average"`
	FastCashCounter MetricSettings `mapstructure:"atm.fastcash.counter"`
}

// Validate checks if the receiver configuration is valid
func (cfg *Config) Validate() error {

	if cfg.CollectionInterval < time.Second*30 {
		return errors.New("when defined, collection_interval cannot be smaller than 30 seconds")
	}

	return nil
}
