package system

import (
	"github.com/alochym01/idrac-exporter/config"
	"github.com/prometheus/client_golang/prometheus"
)

type Bios struct{}

func (bios Bios) Describe(ch chan<- *prometheus.Desc) {
	ch<- config.S_bios
}

func (bios Bios) Collect(ch chan<- prometheus.Metric) {
	metric := config.GOFISH.Service
	systems, sysErr := metric.Systems()

	if nil != sysErr {
		panic(sysErr)
	}

	for _, system := range systems {
		val, biosErr := system.Bios()

		if nil != biosErr {
			panic(biosErr)
		}

		ch<- prometheus.MustNewConstMetric(config.S_bios,
			prometheus.GaugeValue,
			float64(0),
			val.AttributeRegistry,
			val.Description,
			)
	}
}
