package accountsmetrics

import (
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"math/rand"
	"time"
)

type Atm struct{
    ID           int64
	Version      string
	Name         string
	StateID      string
	SerialNumber string
	ISPNetwork   string
}

func generateAtm() Atm{
	i := getRandomNumber(1, 2)
    var newAtm Atm

	switch i {
		case 1:
			newAtm = Atm{
				ID: 111,
				Name: "ATM-111-IL",
				SerialNumber: "atmxph-2022-111",
				Version: "v1.0",
				ISPNetwork: "comcast-chicago",
				StateID: "IL",
		
			}
		
		case 2:
			newAtm = Atm{
				ID: 222,
				Name: "ATM-222-CA",
				SerialNumber: "atmxph-2022-222",
				Version: "v1.0",
				ISPNetwork: "comcast-sanfrancisco",
				StateID: "CA",
			}
	}

	return newAtm
}


func fillMetricWithData(m *pmetric.Metric) {
	m.SetName("accounts.fastcash.value")
	m.SetDescription("Total value withdrawn using fastcash by ATM.")
	m.SetUnit("$")
	m.SetDataType(pmetric.MetricDataTypeSum)
	m.Sum().SetIsMonotonic(true)
	m.Sum().SetAggregationTemporality(pmetric.MetricAggregationTemporalityCumulative)
}

func generateMetrics() pmetric.Metrics{
    metrics := pmetric.NewMetrics()

	resourceMetrics := metrics.ResourceMetrics().AppendEmpty()
	atm := generateAtm()
	fillResourceWithAtm(resourceMetrics.Resource(), atm)

	scopeMetrics := resourceMetrics.ScopeMetrics().AppendEmpty()
	scopeMetrics.Scope().SetName("atm-system")
	scopeMetrics.Scope().SetVersion("1.0")

	mFastCashValue := scopeMetrics.Metrics().AppendEmpty()
    addDataPointToMetric(&mFastCashValue, atm.Name)

	return metrics
}

func addDataPointToMetric(metric *pmetric.Metric, atmNameAttributeValue string) {
	dp := metric.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(startTime)
	dp.SetTimestamp(pcommon.NewTimestampFromTime(time.Now()))
	metricValue := 1000 * getRandomNumber(1, 5)
	dp.SetDoubleVal(float64(metricValue))
	dp.Attributes().Insert("atm-name", pcommon.NewValueString(atmNameAttributeValue))
}

func fillResourceWithAtm(resource pcommon.Resource, atm Atm){
	atmAttrs := resource.Attributes()
	atmAttrs.InsertInt("atm.id", atm.ID)
	atmAttrs.InsertString("atm.stateid", atm.StateID)
	atmAttrs.InsertString("atm.ispnetwork", atm.ISPNetwork)
	atmAttrs.InsertString("atm.serialnumber", atm.SerialNumber) 
 }


 func getRandomNumber(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	i := (rand.Intn(max - min + 1) + min)
    return i	
} 