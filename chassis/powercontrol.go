package chassis

import (
	"fmt"
	"strings"

	"github.com/alochym01/idrac-exporter/config"
	"github.com/prometheus/client_golang/prometheus"
)

// PowerControl is a Chassis Power Control metric
type PowerControl struct{}

// Describe return a description of metrics
func (s PowerControl) Describe(ch chan<- *prometheus.Desc) {
	ch <- config.C_powercontrol
}

// Collect return a metric with all desc value and metric value
func (s PowerControl) Collect(ch chan<- prometheus.Metric) {
	metric := config.GOFISH.Service

	chass, _ := metric.Chassis()

	for _, v := range chass {
		powers, _ := v.Power()
		if powers != nil {
			for _, p := range powers.PowerControl {
				status := config.State_dict[strings.ToUpper(fmt.Sprintf("%v", p.Status.Health))]
				ch <- prometheus.MustNewConstMetric(config.C_powercontrol, prometheus.GaugeValue, status,
					fmt.Sprintf("%v", p.PhysicalContext),
					fmt.Sprintf("%v", p.PowerAllocatedWatts),
					fmt.Sprintf("%v", p.PowerAvailableWatts),
					fmt.Sprintf("%v", p.PowerCapacityWatts),
					fmt.Sprintf("%v", p.PowerConsumedWatts),
					fmt.Sprintf("%v", p.PowerRequestedWatts),
				)

				ch <- prometheus.MustNewConstMetric(config.C_powerconsumedbyall, prometheus.GaugeValue, float64(p.PowerConsumedWatts),
					fmt.Sprintf("%v", p.PowerCapacityWatts),
					fmt.Sprintf("%v", p.MemberID),
				)
			}

		}
	}
}
