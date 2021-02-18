package system

import (
	"fmt"
	"time"

	"github.com/alochym01/idrac-exporter/config"
	"github.com/prometheus/client_golang/prometheus"
)

// Health is a system health metric
type Health struct{}

// Describe return a description of metrics
func (s Health) Describe(ch chan<- *prometheus.Desc) {
	ch <- config.S_health
}

// Collect return a metric with all desc value and metric value
func (s Health) Collect(ch chan<- prometheus.Metric) {
	start := time.Now()
	metric := config.GOFISH.Service

	service, _ := metric.Systems()
	for _, v := range service {
		status := config.State_dict[string(v.Status.Health)]
		ch <- prometheus.MustNewConstMetric(config.S_health, prometheus.GaugeValue, float64(status),
			v.BIOSVersion,
			v.Description,
			v.HostName,
			v.HostedServices,
			v.Manufacturer,
			v.Model,
			v.Name,
			v.PartNumber,
			fmt.Sprintf("%v", v.PowerRestorePolicy),
			fmt.Sprintf("%v", v.PowerState),
			v.SKU,
			v.SerialNumber,
			v.SubModel,
			fmt.Sprintf("%v", v.SystemType),
			v.UUID,
		)
	}

	end := time.Now()

	fmt.Println(start.Second() - end.Second())
}
