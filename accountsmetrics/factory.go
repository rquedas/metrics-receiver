package accountsmetrics

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	receiver "go.opentelemetry.io/collector/receiver"
)

const (
	typeStr         = "accountsmetrics"
	defaultInterval = "1m"
)

func createDefaultConfig() component.Config {
	return &Config{
		Interval: defaultInterval,
	}
}

func createMetricsReceiver(_ context.Context, params receiver.CreateSettings, baseCfg component.Config, consumer consumer.Metrics) (receiver.Metrics, error) {
	if consumer == nil {
		return nil, component.ErrNilNextConsumer
	}

	logger := params.Logger
	accountsReceiverCfg := baseCfg.(*Config)

	atmMetricsRcvr := &accountsmetricsreceiver{
		logger:       logger,
		nextConsumer: consumer,
		config:       accountsReceiverCfg,
	}

	return atmMetricsRcvr, nil

}

// NewFactory creates a factory for tailtracer receiver.
func NewFactory() receiver.Factory {
	return receiver.NewFactory(
		typeStr,
		createDefaultConfig,
		receiver.WithMetrics(createMetricsReceiver, component.StabilityLevelStable))
}
