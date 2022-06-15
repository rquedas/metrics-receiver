package accountsmetrics

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.uber.org/zap"
)

type accountsmetricsreceiver struct {
    host component.Host
	cancel context.CancelFunc
	logger       *zap.Logger
	nextConsumer consumer.Metrics
	config       *Config
}

var startTime pcommon.Timestamp

func (tailtracerRcvr *accountsmetricsreceiver) Start(ctx context.Context, host component.Host) error {
    tailtracerRcvr.host = host
    ctx = context.Background()
	ctx, tailtracerRcvr.cancel = context.WithCancel(ctx)
 
	interval, _ := time.ParseDuration(tailtracerRcvr.config.Interval)
	startTime = pcommon.NewTimestampFromTime(time.Now())
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				tailtracerRcvr.logger.Info("I should start processing metrics now!")
				tailtracerRcvr.nextConsumer.ConsumeMetrics(ctx, generateMetrics())
			case <-ctx.Done():
				return
			}
		}
	}()

	return nil
}

func (tailtracerRcvr *accountsmetricsreceiver) Shutdown(ctx context.Context) error {
	tailtracerRcvr.cancel()
	tailtracerRcvr.logger.Info("I am done and ready to shutdown!")
	return nil
}