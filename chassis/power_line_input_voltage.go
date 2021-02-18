package chassis

import (
	"fmt"
	"github.com/alochym01/idrac-exporter/config"
	"github.com/prometheus/client_golang/prometheus"
)

type PowerLineInputVoltage struct{}

func (power PowerLineInputVoltage) Describe(ch chan<- *prometheus.Desc) {
	ch<- config.C_power_line_input_voltage
}

func (power PowerLineInputVoltage) Collect(ch chan<- prometheus.Metric) {
	metric := config.GOFISH.Service

	chass, _ := metric.Chassis()

	for _, v := range chass {
		powers, _ := v.Power()

		if nil != powers {
			supplies := powers.PowerSupplies

			for _, supply := range supplies {
				ch <- prometheus.MustNewConstMetric(config.C_power_line_input_voltage,
					prometheus.GaugeValue,
					float64(supply.LineInputVoltage),
					supply.MemberID,
					fmt.Sprintf("%v", supply.LineInputVoltageType),
				)
			}

		}
	}
}
