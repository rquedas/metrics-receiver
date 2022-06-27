package atmmetrics

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver/scraperhelper"
)

const (
	typeStr = "atmmetrics"
)

func createDefaultConfig() config.Receiver {
	scs := scraperhelper.NewDefaultScraperControllerSettings(typeStr)
	scs.CollectionInterval = 10 * time.Second
	defaultMetrics := DefaultMetricsSettings()


	return &Config{
		ScraperControllerSettings: scs,
		Metrics: defaultMetrics,
	}
}

func DefaultMetricsSettings() TransactionMetricsSettings {
	return TransactionMetricsSettings{
		FastCashTotal: MetricSettings{
			Enabled: true,
		},
		FastCashAverage: MetricSettings{
			Enabled: true,
		},
	}
}

func createMetricsScraperReceiver(
	ctx context.Context,
	params component.ReceiverCreateSettings,
	config config.Receiver,
	consumer consumer.Metrics,
) (component.MetricsReceiver, error) {
	atmConfig := config.(*Config)

	dsr, err := NewReceiver(ctx, params, atmConfig, consumer)
	if err != nil {
		return nil, err
	}

	return dsr, nil
}

// NewFactory creates a factory for tailtracer receiver.
func NewFactory() component.ReceiverFactory {
	return component.NewReceiverFactory(
		typeStr,
		createDefaultConfig,
		component.WithMetricsReceiver(createMetricsScraperReceiver))
}