package system

import (
	"fmt"
	"math"
	"github.com/alochym01/idrac-exporter/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stmcginnis/gofish/redfish"
)

type Drive struct{}

func (d Drive) Describe(ch chan<- *prometheus.Desc) {
	ch<- config.S_storage_drive
	ch<- config.S_storage_drive_predicted_media_life_left_percent
}

func (d Drive) Collect(ch chan<- prometheus.Metric) {
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
				drives, driveErr := storage.Drives()

				if nil != driveErr {
					panic(driveErr)
				}

				d.helper(ch, drives)
			}
		}
	}
}

func (d Drive) helper(ch chan<- prometheus.Metric, drives []*redfish.Drive) {
	for _, drive := range drives {
		status := config.State_dict[string(drive.Status.Health)]
		ch<- prometheus.MustNewConstMetric(config.S_storage_drive,
			prometheus.GaugeValue,
			float64(status),
			fmt.Sprintf("%v", drive.BlockSizeBytes),
			fmt.Sprintf("%v", drive.CapableSpeedGbs),
			d.convertCapacity(float64(drive.CapacityBytes)),
			drive.Description,
			fmt.Sprintf("%v", drive.IndicatorLED),
			drive.Manufacturer,
			fmt.Sprintf("%v", drive.MediaType),
			drive.Model,
			drive.PartNumber,
			fmt.Sprintf("%v", drive.Protocol),
			drive.Revision,
			drive.SerialNumber,
		)
		
		if "SSD" == fmt.Sprintf("%v", drive.MediaType) {
			ch<- prometheus.MustNewConstMetric(config.S_storage_drive_predicted_media_life_left_percent,
				prometheus.GaugeValue,
				float64(drive.PredictedMediaLifeLeftPercent),
				fmt.Sprintf("%v", drive.BlockSizeBytes),
				fmt.Sprintf("%v", drive.CapableSpeedGbs),
				d.convertCapacity(float64(drive.CapacityBytes)),
				drive.Description,
				drive.Manufacturer,
				fmt.Sprintf("%v", drive.MediaType),
				drive.Model,
				drive.PartNumber,
				fmt.Sprintf("%v", drive.Protocol),
				drive.Revision,
				drive.SerialNumber,
			)
		}
	}
}

func (d Drive) convertCapacity(num float64) string {
	units := []string{"TB", "GB", "MB", "KB", "B"}
	idx := len(units) - 1
	
	for idx > -1 && num >= 1000 {
		idx -= 1
		num = num / 1000
	}
	
	return fmt.Sprintf("%v", math.RoundToEven(num)) + units[idx]
}
