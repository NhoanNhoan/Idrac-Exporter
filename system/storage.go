package system

import (
	"fmt"
	"github.com/alochym01/idrac-exporter/config"
	"github.com/prometheus/client_golang/prometheus"
)

type Storage struct{}

func (storage Storage) Describe(ch chan<- *prometheus.Desc) {
	ch <- config.S_storage
}

func (storage Storage) Collect(ch chan<- prometheus.Metric) {
	metric := config.GOFISH.Service
	systems, sysErr := metric.Systems()

	if nil != sysErr {
		panic(sysErr)
	}

	for _, system := range systems {
		storages, storageErr := system.Storage()

		if nil != storageErr {
			panic(storageErr)
		}

		if 0 != len(storages) {
			for _, storage := range storages {
				status := config.State_dict[string(storage.Status.Health)]
				ch<- prometheus.MustNewConstMetric(config.S_storage,
					prometheus.GaugeValue,
					float64(status),
					storage.Description,
					fmt.Sprintf("%v", storage.DrivesCount),
					fmt.Sprintf("%v", storage.RedundancyCount),
					fmt.Sprintf("%v", storage.EnclosuresCount))
			}
		}
	}
}
