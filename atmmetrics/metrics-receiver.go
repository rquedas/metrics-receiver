package atmmetrics

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver/scraperhelper"
	"go.uber.org/zap"
)

type atmmetricsreceiver struct {
    host component.Host
	logger       *zap.Logger
	settings component.ReceiverCreateSettings
	config       *Config
}

var startTime pcommon.Timestamp


func (acntsMtrcRcvr *atmmetricsreceiver) start(ctx context.Context, host component.Host) error {
	acntsMtrcRcvr.host = host
	startTime = pcommon.NewTimestampFromTime(time.Now())
	return nil
}

func (acntsMtrcRcvr *atmmetricsreceiver) scrape(ctx context.Context) (pmetric.Metrics, error) {
	acntsMtrcRcvr.logger.Info("I should start processing metrics now!")
	md := generateMetrics()
	return md, nil
}

func NewReceiver(
	_ context.Context,
	set component.ReceiverCreateSettings,
	config *Config,
	nextConsumer consumer.Metrics,
) (component.MetricsReceiver, error) {
	atmRecv := atmmetricsreceiver{
		config:   config,
		settings: set,
	}

	atmMetricsScrp, err := scraperhelper.NewScraper(typeStr, atmRecv.scrape, scraperhelper.WithStart(atmRecv.start))
	if err != nil {
		return nil, err
	}
	
	return scraperhelper.NewScraperControllerReceiver(&atmRecv.config.ScraperControllerSettings, set, nextConsumer, scraperhelper.AddScraper(atmMetricsScrp))
}