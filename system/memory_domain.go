package system

import (
	"fmt"
	"github.com/alochym01/idrac-exporter/config"
	"github.com/prometheus/client_golang/prometheus"
)

type MemoryDomain struct{}

func (domain MemoryDomain) Describe(ch chan<- *prometheus.Desc) {
	ch<- config.S_memorydomain
}

func (domain MemoryDomain) Collect(ch chan<- prometheus.Metric) {
	metric := config.GOFISH.Service
	systems, sysErr := metric.Systems()

	if nil != sysErr {
		panic(sysErr)
	}

	for _, system := range systems {
		domains, domainErr := system.MemoryDomains()

		if nil != domainErr {
			panic(domainErr)
		}

		if 0 != len(domains) {
			for _, memDomain := range domains {
				ch<- prometheus.MustNewConstMetric(config.S_memorydomain,
					prometheus.GaugeValue,
					float64(0),
					fmt.Sprintf("%v", memDomain.AllowsBlockProvisioning),
					fmt.Sprintf("%v", memDomain.AllowsMemoryChunkCreation),
					fmt.Sprintf("%v", memDomain.AllowsMirroring),
					fmt.Sprintf("%v", memDomain.AllowsSparing),
					memDomain.Description)
			}
		}
	}
}
