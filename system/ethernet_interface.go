package system

import (
	"fmt"

	"github.com/alochym01/idrac-exporter/config"

	"github.com/prometheus/client_golang/prometheus"
)

type EthernetInterface struct{}

func (e EthernetInterface) Describe(ch chan<- *prometheus.Desc) {
	ch<- config.S_ethernetinterface
}

func (e EthernetInterface) Collect(ch chan<- prometheus.Metric) {
	metric := config.GOFISH.Service
	systems, err := metric.Systems()

	if nil != err {
		panic(err)
	}

	for _, system := range systems {
		ethernetInterfaces, ethernetErr := system.EthernetInterfaces()
		if nil != ethernetErr {
			panic(ethernetErr)
		}

		for _, ethernetInterface := range ethernetInterfaces {
			status := config.State_dict[string(ethernetInterface.Status.Health)]
			ch <- prometheus.MustNewConstMetric(config.S_ethernetinterface,
				prometheus.GaugeValue,
				float64(status),
				fmt.Sprintf("%v", ethernetInterface.AutoNeg),
				ethernetInterface.Description,
				fmt.Sprintf("%v", ethernetInterface.EthernetInterfaceType),
				ethernetInterface.FQDN,
				fmt.Sprintf("%v", ethernetInterface.FullDuplex),
				ethernetInterface.HostName,
				ethernetInterface.MACAddress,
				fmt.Sprintf("%v", ethernetInterface.MTUSize),
				fmt.Sprintf("%v", ethernetInterface.SpeedMbps),
			)
		}
	}
}
