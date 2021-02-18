package chassis

import (
	"fmt"

	"github.com/alochym01/idrac-exporter/config"

	"github.com/prometheus/client_golang/prometheus"
)

type Fan struct{}

func (fan Fan) Describe(ch chan<- *prometheus.Desc) {
	ch <- config.C_fan_status
	ch <- config.C_fan_reading
}

func (fan Fan) Collect(ch chan<- prometheus.Metric) {
	metric := config.GOFISH.Service
	chassisArr, _ := metric.Chassis()

	for _, chassis := range chassisArr {
		thermal, _ := chassis.Thermal()
		if nil != thermal {
			fans := thermal.Fans

			for _, fan := range fans {
				status := config.State_dict[string(fan.Status.Health)]
				ch <- prometheus.MustNewConstMetric(config.C_fan_status, prometheus.GaugeValue, float64(status),
					fmt.Sprintf("%v", fan.IndicatorLED),
					fmt.Sprintf("%v", fan.LowerThresholdCritical),
					fmt.Sprintf("%v", fan.LowerThresholdFatal),
					fmt.Sprintf("%v", fan.LowerThresholdNonCritical),
					fan.Manufacturer,
					fmt.Sprintf("%v", fan.MaxReadingRange),
					fan.MemberID,
					fmt.Sprintf("%v", fan.MinReadingRange),
					fan.Model,
					fan.PartNumber,
					fan.PhysicalContext,
					fmt.Sprintf("%v", fan.Reading),
					fmt.Sprintf("%v", fan.ReadingUnits),
					fmt.Sprintf("%v", fan.RedundancyCount),
					fmt.Sprintf("%v", fan.SensorNumber),
					fan.SerialNumber,
					fan.SparePartNumber,
					fmt.Sprintf("%v", fan.UpperThresholdCritical),
					fmt.Sprintf("%v", fan.UpperThresholdFatal),
					fmt.Sprintf("%v", fan.UpperThresholdNonCritical),
				)

				ch <- prometheus.MustNewConstMetric(config.C_fan_reading, prometheus.GaugeValue, float64(fan.Reading),
					fmt.Sprintf("%v", fan.LowerThresholdCritical),
					fmt.Sprintf("%v", fan.LowerThresholdFatal),
					fmt.Sprintf("%v", fan.LowerThresholdNonCritical),
					fan.Manufacturer,
					fmt.Sprintf("%v", fan.MaxReadingRange),
					fan.MemberID,
					fmt.Sprintf("%v", fan.MinReadingRange),
					fan.Model,
					fan.PartNumber,
					fan.PhysicalContext,
					fmt.Sprintf("%v", fan.ReadingUnits),
					fmt.Sprintf("%v", fan.SensorNumber),
					fan.SerialNumber,
					fan.SparePartNumber,
					fmt.Sprintf("%v", fan.UpperThresholdCritical),
					fmt.Sprintf("%v", fan.UpperThresholdFatal),
					fmt.Sprintf("%v", fan.UpperThresholdNonCritical),
				)
			}
		}
	}
}
