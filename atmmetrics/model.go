package atmmetrics

import (
	"math/rand"
	"time"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

type Atm struct {
	ID           int64
	Version      string
	Name         string
	StateID      string
	SerialNumber string
	ISPNetwork   string
	CountryID    string
}

func generateAtm() Atm {
	i := getRandomNumber(1, 6)
	var newAtm Atm

	switch i {
	case 1:
		newAtm = Atm{
			ID:           111,
			Name:         "ATM-111-IL",
			SerialNumber: "atmxph-2022-111",
			Version:      "v1.0",
			ISPNetwork:   "comcast-chicago",
			StateID:      "IL",
			CountryID:    "USA",
		}

	case 2:
		newAtm = Atm{
			ID:           222,
			Name:         "ATM-222-CA",
			SerialNumber: "atmxph-2022-222",
			Version:      "v1.0",
			ISPNetwork:   "comcast-sanfrancisco",
			StateID:      "CA",
			CountryID:    "USA",
		}

	case 3:
		newAtm = Atm{
			ID:           333,
			Name:         "ATM-333-IL",
			SerialNumber: "atmxph-2022-333",
			Version:      "v1.0",
			ISPNetwork:   "comcast-chicago",
			StateID:      "IL",
			CountryID:    "USA",
		}

	case 4:
		newAtm = Atm{
			ID:           444,
			Name:         "ATM-444-CA",
			SerialNumber: "atmxph-2022-444",
			Version:      "v1.0",
			ISPNetwork:   "att-sandiego",
			StateID:      "CA",
			CountryID:    "USA",
		}

	case 5:
		newAtm = Atm{
			ID:           555,
			Name:         "ATM-555-SP",
			SerialNumber: "atmxph-2022-555",
			Version:      "v1.0",
			ISPNetwork:   "claro-saopaulo",
			StateID:      "SP",
			CountryID:    "BRAZIL",
		}

	case 6:
		newAtm = Atm{
			ID:           666,
			Name:         "ATM-666-RJ",
			SerialNumber: "atmxph-2022-666",
			Version:      "v1.0",
			ISPNetwork:   "oi-rio",
			StateID:      "RJ",
			CountryID:    "BRAZIL",
		}

	}

	return newAtm
}

func fillMetricWithData(m *pmetric.Metric) {
	m.SetName("atm.fastcash.total")
	m.SetDescription("Total value withdrawn using fastcash by ATM.")
	m.SetUnit("$")
	m.SetEmptySum()
	m.Sum().SetIsMonotonic(true)
	m.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func generateMetrics() pmetric.Metrics {
	metrics := pmetric.NewMetrics()

	resourceMetrics := metrics.ResourceMetrics().AppendEmpty()
	atm := generateAtm()
	fillResourceWithAtm(resourceMetrics.Resource(), atm)

	scopeMetrics := resourceMetrics.ScopeMetrics().AppendEmpty()
	scopeMetrics.Scope().SetName("atm-system")
	scopeMetrics.Scope().SetVersion("1.0")

	mFastCashValue := scopeMetrics.Metrics().AppendEmpty()
	fillMetricWithData(&mFastCashValue)
	addDataPointToMetric(&mFastCashValue, atm.Name, atm.CountryID)

	return metrics
}

func addDataPointToMetric(metric *pmetric.Metric, atmNameAttributeValue string, atmCountryAttributeValue string) {
	dp := metric.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(startTime)
	dp.SetTimestamp(pcommon.NewTimestampFromTime(time.Now()))
	metricValue := 1000 * getRandomNumber(1, 5)
	dp.SetDoubleValue(float64(metricValue))
	dp.Attributes().PutStr("atm-name", atmNameAttributeValue)
	dp.Attributes().PutStr("atm-country", atmCountryAttributeValue)
}

func fillResourceWithAtm(resource pcommon.Resource, atm Atm) {
	atmAttrs := resource.Attributes()
	atmAttrs.PutInt("atm.id", atm.ID)
	atmAttrs.PutStr("atm.stateid", atm.StateID)
	atmAttrs.PutStr("atm.ispnetwork", atm.ISPNetwork)
	atmAttrs.PutStr("atm.serialnumber", atm.SerialNumber)
	atmAttrs.PutStr("atm.country", atm.CountryID)
}

func getRandomNumber(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	i := (rand.Intn(max-min+1) + min)
	return i
}
