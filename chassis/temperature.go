package chassis

import (
	"fmt"
	
	"strings"

	"github.com/alochym01/idrac-exporter/config"

	"github.com/prometheus/client_golang/prometheus"
)

// Temperature ...
type Temperature struct{}

// Describe ...
func (tem Temperature) Describe(ch chan<- *prometheus.Desc) {
	ch <- config.C_temperature_reading
	ch <- config.C_temperature_status
}

// Collect ...
// Service -> Systems, Chassis
// Chassis -> Thermal()(*Thermal, err)
// Thermal() -> []Temperatures
func (temp Temperature) Collect(ch chan<- prometheus.Metric) {
	metric := config.GOFISH.Service
	chassisArr, _ := metric.Chassis()

	for _, chassis := range chassisArr {
		thermal, _ := chassis.Thermal() // type *Thermal

		if nil != thermal {
			for _, val := range thermal.Temperatures {
				status := config.State_dict[strings.ToUpper(fmt.Sprintf("%v", val.Status.Health))]

				ch <- prometheus.MustNewConstMetric(
					config.C_temperature_reading,
					prometheus.GaugeValue,
					float64(val.ReadingCelsius),
					fmt.Sprintf("%v", val.AdjustedMaxAllowableOperatingValue),
					fmt.Sprintf("%v", val.AdjustedMinAllowableOperatingValue),
					val.DeltaPhysicalContext,
					fmt.Sprintf("%v", val.DeltaReadingCelsius),
					fmt.Sprintf("%v", val.LowerThresholdCritical),
					fmt.Sprintf("%v", val.LowerThresholdFatal),
					fmt.Sprintf("%v", val.LowerThresholdNonCritical),
					fmt.Sprintf("%v", val.LowerThresholdUser),
					fmt.Sprintf("%v", val.MaxAllowableOperatingValue),
					fmt.Sprintf("%v", val.MaxReadingRangeTemp),
					val.MemberID,
					fmt.Sprintf("%v", val.MinAllowableOperatingValue),
					fmt.Sprintf("%v", val.MinReadingRangeTemp),
					val.PhysicalContext,
					fmt.Sprintf("%v", val.SensorNumber),
					fmt.Sprintf("%v", status),
					fmt.Sprintf("%v", val.UpperThresholdCritical),
					fmt.Sprintf("%v", val.UpperThresholdFatal),
					fmt.Sprintf("%v", val.UpperThresholdNonCritical),
					fmt.Sprintf("%v", val.UpperThresholdUser),
				)

				ch <- prometheus.MustNewConstMetric(
					config.C_temperature_status,
					prometheus.GaugeValue,
					float64(0),
					fmt.Sprintf("%v", val.AdjustedMaxAllowableOperatingValue),
					fmt.Sprintf("%v", val.AdjustedMinAllowableOperatingValue),
					val.DeltaPhysicalContext,
					fmt.Sprintf("%v", val.DeltaReadingCelsius),
					fmt.Sprintf("%v", val.LowerThresholdCritical),
					fmt.Sprintf("%v", val.LowerThresholdFatal),
					fmt.Sprintf("%v", val.LowerThresholdNonCritical),
					fmt.Sprintf("%v", val.LowerThresholdUser),
					fmt.Sprintf("%v", val.MaxAllowableOperatingValue),
					fmt.Sprintf("%v", val.MaxReadingRangeTemp),
					val.MemberID,
					fmt.Sprintf("%v", val.MinAllowableOperatingValue),
					fmt.Sprintf("%v", val.MinReadingRangeTemp),
					val.PhysicalContext,
					fmt.Sprintf("%v", val.ReadingCelsius),
					fmt.Sprintf("%v", val.SensorNumber),
					fmt.Sprintf("%v", status),
					fmt.Sprintf("%v", val.UpperThresholdCritical),
					fmt.Sprintf("%v", val.UpperThresholdFatal),
					fmt.Sprintf("%v", val.UpperThresholdNonCritical),
					fmt.Sprintf("%v", val.UpperThresholdUser),
				)
			}
		}
	}
}
