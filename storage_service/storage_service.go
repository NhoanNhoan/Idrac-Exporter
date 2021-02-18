package storage_service

import (
	"fmt"
	"github.com/alochym01/idrac-exporter/config"
	"github.com/prometheus/client_golang/prometheus"
)

type StorageService struct{}

func (storage StorageService) Describe(ch chan<- *prometheus.Desc) {
	ch<- config.SS_storageservice
}

func (storage StorageService) Collect(ch chan<- prometheus.Metric) {
	metric := config.GOFISH.Service
	storageServices, storageErr := metric.StorageServices()

	if nil != storageErr {
		panic(storageErr)
	}

	for _, service := range storageServices {
		ch<- prometheus.MustNewConstMetric(config.SS_storageservice,
			prometheus.GaugeValue,
			float64(0),
			service.Description,
			fmt.Sprintf("%v", service.RedundancyCount),
			fmt.Sprintf("%v", service.Status.Health),
		)
	}
}
