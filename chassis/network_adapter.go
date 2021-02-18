package chassis

import (
	"github.com/alochym01/idrac-exporter/config"
	"github.com/prometheus/client_golang/prometheus"
)

type NetworkAdapter struct{}

func (network NetworkAdapter) Describe(ch chan<- *prometheus.Desc) {
	ch<- config.C_networkadapter
}

func (network NetworkAdapter) Collect(ch chan<- prometheus.Metric) {
	metric := config.GOFISH.Service
	chassisArr, chassisErr := metric.Chassis()

	if nil != chassisErr {
		panic(chassisErr)
	}

	for _, chassis := range chassisArr {
		adapters, adapterErr := chassis.NetworkAdapters()

		if nil != adapterErr {
			panic(adapterErr)
		}

		if 0 != len(adapters) {
			for _, adapter := range adapters {
				status := config.State_dict[string(adapter.Status.Health)]
				ch<- prometheus.MustNewConstMetric(config.C_networkadapter,
					prometheus.GaugeValue,
					float64(status),
					adapter.Description,
					adapter.Manufacturer,
					adapter.Model,
					adapter.PartNumber,
					adapter.SKU,
					adapter.SerialNumber,
					)
			}
		}
	}
}
