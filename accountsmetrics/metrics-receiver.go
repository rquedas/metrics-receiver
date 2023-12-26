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

func (acntsMtrcRcvr *accountsmetricsreceiver) Start(ctx context.Context, host component.Host) error {
    acntsMtrcRcvr.host = host
    ctx = context.Background()
	ctx, acntsMtrcRcvr.cancel = context.WithCancel(ctx)
 
	interval, _ := time.ParseDuration(acntsMtrcRcvr.config.Interval)
	startTime = pcommon.NewTimestampFromTime(time.Now())
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				acntsMtrcRcvr.logger.Info("I should start processing metrics now!")
				acntsMtrcRcvr.nextConsumer.ConsumeMetrics(ctx, generateMetrics())
			case <-ctx.Done():
				return
			}
		}
	}()

	return nil
}

func (acntsMtrcRcvr *accountsmetricsreceiver) Shutdown(ctx context.Context) error {
	acntsMtrcRcvr.cancel()
	acntsMtrcRcvr.logger.Info("I am done and ready to shutdown!")
	return nil
}