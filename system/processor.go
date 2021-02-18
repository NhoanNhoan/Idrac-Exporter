package system

import (
	"fmt"

	"github.com/alochym01/idrac-exporter/config"

	"github.com/prometheus/client_golang/prometheus"
)

type Processor struct{}

func (pro Processor) Describe(ch chan<- *prometheus.Desc) {
	ch<- config.S_processor
}

func (pro Processor) Collect(ch chan<- prometheus.Metric) {
	metric := config.GOFISH.Service
	systems, systemErr := metric.Systems()

	if nil != systemErr {
		panic(systemErr)
	}

	for _, system := range systems {
		processors, proErr := system.Processors()

		if nil != proErr {
			panic(proErr)
		}

		for _, processor := range processors {
			status := config.State_dict[string(processor.Status.Health)]
			ch<- prometheus.MustNewConstMetric(config.S_processor,
				prometheus.GaugeValue,
				float64(status),
				processor.Actions,
				processor.Description,
				processor.Manufacturer,
				fmt.Sprintf("%v", processor.MaxSpeedMHz),
				fmt.Sprintf("%v", processor.MaxTDPWatts),
				processor.Model,
				fmt.Sprintf("%v", processor.ProcessorType),
				processor.Socket,
				processor.SubProcessors,
				fmt.Sprintf("%v", processor.TDPWatts),
				fmt.Sprintf("%v", processor.TotalCores),
				fmt.Sprintf("%v", processor.TotalEnabledCores),
				fmt.Sprintf("%v", processor.TotalThreads),
				processor.UUID,
				)
		}
	}
}
