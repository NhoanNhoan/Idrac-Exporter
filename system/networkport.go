package system

import (
	"fmt"
	"github.com/alochym01/idrac-exporter/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stmcginnis/gofish/redfish"
)

type NetworkPort struct{}

func (port NetworkPort) Describe(ch chan<- *prometheus.Desc) {
	ch<- config.S_networkport
}

func (port NetworkPort) Collect(ch chan<- prometheus.Metric) {
	metric := config.GOFISH.Service
	systems, sysErr := metric.Systems()

	if nil != sysErr {
		panic(sysErr)
	}

	if 0 != len(systems) {
		port.makeNetworkPortMetricFromSystems(ch, systems)
	}
}

func (port NetworkPort) makeNetworkPortMetricFromSystems(ch chan<- prometheus.Metric,
	systems []*redfish.ComputerSystem) {
	for _, system := range systems {
		interfaces, err := system.NetworkInterfaces()

		if nil != err {
			panic(err)
		}

		if 0 != len(interfaces) {
			port.makeNetworkPortMetricFromNetworkInterfaces(ch, interfaces)
		}
	}
}

func (port NetworkPort) makeNetworkPortMetricFromNetworkInterfaces(ch chan<- prometheus.Metric,
	interfaces []*redfish.NetworkInterface) {
	for _, netInterface := range interfaces {
		adapter, err := netInterface.NetworkAdapter()

		if nil != err {
			panic(err)
		}

		if nil != adapter {
			port.makeNetworkPortMetricFromNetworkAdapter(ch, adapter)
		}
	}
}

func (port NetworkPort) makeNetworkPortMetricFromNetworkAdapter(ch chan<- prometheus.Metric,
	adapter *redfish.NetworkAdapter) {
	networkPorts, err := adapter.NetworkPorts()

	if nil != err {
		panic(err)
	}

	for _, networkPort := range networkPorts {
		status := config.State_dict[string(networkPort.Status.Health)]
		ch<- prometheus.MustNewConstMetric(config.S_networkport,
			prometheus.GaugeValue,
			float64(status),
			adapter.Manufacturer,
			fmt.Sprintf("%v", networkPort.LinkStatus),
			fmt.Sprintf("%v", networkPort.CurrentLinkSpeedMbps),
			networkPort.Description,
			fmt.Sprintf("%v", networkPort.MaxFrameSize),
			fmt.Sprintf("%v", networkPort.NumberDiscoveredRemotePorts),
			networkPort.PhysicalPortNumber,
			fmt.Sprintf("%v", networkPort.PortMaximumMTU),
			)
	}
}
