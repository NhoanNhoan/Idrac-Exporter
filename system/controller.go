package system

import (
	"fmt"
	"github.com/alochym01/idrac-exporter/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stmcginnis/gofish/redfish"
)

type Controller struct{}

func (controller Controller) Describe(ch chan<- *prometheus.Desc) {
	ch<- config.S_network_adapter_status
}

func (controller Controller) Collect(ch chan<- prometheus.Metric) {
	metric := config.GOFISH.Service
	systems, err := metric.Systems()

	if nil != err {
		panic(err)
	}

	if 0 != len(systems) {
		controller.makeMetricFromSystems(ch, systems)
	}
}

func (controller Controller) makeMetricFromSystems(ch chan<- prometheus.Metric,
	systems []*redfish.ComputerSystem) {
	for _, system := range systems {
		interfaces, err := system.NetworkInterfaces()

		if nil != err {
			panic(err)
		}

		if 0 != len(interfaces) {
			controller.makeMetricFromNetworkInterfaces(ch, interfaces)
		}
	}
}

func (controller Controller) makeMetricFromNetworkInterfaces(ch chan<- prometheus.Metric,
	interfaces []*redfish.NetworkInterface) {
	for _, netInterface := range interfaces {
		adapter, err := netInterface.NetworkAdapter()

		if nil != err {
			panic(err)
		}

		if nil != adapter {
			controller.makeMetricFromNetworkAdapter(ch, adapter)
		}
	}
}

func (controller Controller) makeMetricFromNetworkAdapter(ch chan<- prometheus.Metric,
	adapter *redfish.NetworkAdapter) {
	controllers := adapter.Controllers

	if 0 != len(controllers) {
		for _, control := range controllers {
			ch <- prometheus.MustNewConstMetric(config.S_network_adapter_status,
				prometheus.GaugeValue,
				float64(0),
				adapter.Manufacturer,
				control.FirmwarePackageVersion,
				fmt.Sprintf("%v", control.NetworkDeviceFunctionsCount),
				fmt.Sprintf("%v", control.NetworkPortsCount),
				)
		}
	}
}
