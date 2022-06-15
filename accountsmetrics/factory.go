package accountsmetrics

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/consumer"
)

const (
	typeStr = "accountsmetrics"
	defaultInterval = "1m"
)

func createDefaultConfig() config.Receiver {
	return &Config{
		ReceiverSettings:   config.NewReceiverSettings(config.NewComponentID(typeStr)),
		Interval: defaultInterval,
	}
}

func createMetricsReceiver(_ context.Context, params component.ReceiverCreateSettings, baseCfg config.Receiver, consumer consumer.Metrics) (component.MetricsReceiver, error) {
	if consumer == nil {
		return nil, component.ErrNilNextConsumer
	}
	
	logger := params.Logger
	tailtracerCfg := baseCfg.(*Config)

	traceRcvr := &accountsmetricsreceiver{
		logger:       logger,
		nextConsumer: consumer,
		config:       tailtracerCfg,
	}
	
	return traceRcvr, nil

}

// NewFactory creates a factory for tailtracer receiver.
func NewFactory() component.ReceiverFactory {
	return component.NewReceiverFactory(
		typeStr,
		createDefaultConfig,
		component.WithMetricsReceiver(createMetricsReceiver))
}