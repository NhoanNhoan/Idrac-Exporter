package system

import (
	"fmt"

	"github.com/alochym01/idrac-exporter/config"

	"github.com/prometheus/client_golang/prometheus"
)

type Memory struct{}

func (mem Memory) Describe(ch chan<- *prometheus.Desc) {
	ch <- config.S_memory
}

func (mem Memory) Collect(ch chan<- prometheus.Metric){
	metric := config.GOFISH.Service
	systems, _ := metric.Systems()

	for _, system := range systems {
		memories, err := system.Memory()
		if nil == err {
			for _, memory := range memories {
				status := config.State_dict[string(memory.Status.Health)]
				ch <- prometheus.MustNewConstMetric(config.S_memory, prometheus.GaugeValue, float64(status),
					fmt.Sprintf("%v", memory.AllocationAlignmentMiB),
					fmt.Sprintf("%v", memory.AllocationIncrementMiB),
					fmt.Sprintf("%v", memory.BaseModuleType),
					fmt.Sprintf("%v", memory.BusWidthBits),
					fmt.Sprintf("%v", memory.CacheSizeMiB),
					fmt.Sprintf("%v", memory.CapacityMiB),
					fmt.Sprintf("%v", memory.ConfigurationLocked),
					fmt.Sprintf("%v", memory.DataWidthBits),
					memory.Description,
					memory.DeviceLocator,
					fmt.Sprintf("%v", memory.ErrorCorrection),
					memory.FirmwareAPIVersion,
					memory.FirmwareRevision,
					fmt.Sprintf("%v", memory.IsRankSpareEnabled),
					fmt.Sprintf("%v", memory.IsSpareDeviceEnabled),
					fmt.Sprintf("%v", memory.LogicalSizeMiB),
					memory.Manufacturer,
					fmt.Sprintf("%v", memory.MemoryDeviceType),
					fmt.Sprintf("%v", memory.MemoryType),
					fmt.Sprintf("%v", memory.OperatingSpeedMhz),
					memory.PartNumber,
					fmt.Sprintf("%v", memory.RankCount),
					memory.SerialNumber,
					)
			}
		}
	}
}
