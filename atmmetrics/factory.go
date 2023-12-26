package atmmetrics

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	receiver "go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/scraperhelper"
)

const (
	typeStr = "atmmetrics"
)

func createDefaultConfig() component.Config {
	scs := scraperhelper.NewDefaultScraperControllerSettings(typeStr)
	scs.CollectionInterval = 30 * time.Second
	defaultMetrics := DefaultMetricsSettings()

	return &Config{
		ScraperControllerSettings: scs,
		Metrics:                   defaultMetrics,
	}
}

func DefaultMetricsSettings() TransactionMetricsSettings {
	return TransactionMetricsSettings{
		FastCashCounter: MetricSettings{
			Enabled: true,
		},
	}
}

func createMetricsScraperReceiver(
	ctx context.Context,
	params receiver.CreateSettings,
	config component.Config,
	consumer consumer.Metrics,
) (receiver.Metrics, error) {
	// atmConfig := config.(*Config)

	// dsr, err := NewReceiver(ctx, params, atmConfig, consumer)
	// if err != nil {
	// 	return nil, err
	// }

	// return dsr, nil
	return nil, nil
}

// NewFactory creates a factory for tailtracer receiver.
func NewFactory() receiver.Factory {
	return receiver.NewFactory(
		typeStr,
		createDefaultConfig,
		receiver.WithMetrics(createMetricsScraperReceiver, component.StabilityLevelStable))
}
