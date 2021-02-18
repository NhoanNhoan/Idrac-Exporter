package system

import (
	"fmt"
	"math"
	"strings"
	"github.com/alochym01/idrac-exporter/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stmcginnis/gofish/redfish"
)

type Volume struct{}

func (vol Volume) Describe(ch chan<- *prometheus.Desc) {
	ch<- config.S_storage_volume
}

func (vol Volume) Collect(ch chan<- prometheus.Metric) {
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

		for _, storage := range storages{
			volumes, volErr := storage.Volumes()
			if nil != volErr {
				panic(volErr)
			}

			vol.helper(ch, volumes)
		}
	}
}

func (vol Volume) helper(ch chan<- prometheus.Metric, volumes []*redfish.Volume) {
	for _, volume := range volumes {
		status := config.State_dict[string(volume.Status.Health)]
		ch <- prometheus.MustNewConstMetric(config.S_storage_volume,
			prometheus.GaugeValue,
			float64(status),
			volume.Description,
			vol.convertCapacityBytes(float64(volume.CapacityBytes)),
			fmt.Sprintf("%v", volume.VolumeType),
			fmt.Sprintf("%v", volume.Encrypted),
			fmt.Sprintf("%v", volume.BlockSizeBytes),
			fmt.Sprintf("%v", volume.DrivesCount),
			fmt.Sprintf("%v", vol.associatedDriveIds(volume)),
		)
	}
}

func (vol Volume) convertCapacityBytes(capacity float64) string {
	units := []string{"TB", "GB", "MB", "KB", "B"}
	idx := len(units) - 1

	for idx > -1 && capacity >= 1000 {
		capacity = capacity / 1000
		idx = idx - 1
	}

	return fmt.Sprintf("%v", math.RoundToEven(capacity)) + units[idx]
}

func (vol Volume) associatedDriveIds(volume *redfish.Volume) []string {
	drives, _ := volume.Drives()
	driveId := make([]string, 0)	

	if 0 != len(drives) {
		for _, drive := range drives {
			words := strings.Split(drive.Description, " ")
			driveId = append(driveId, words[len(words) - 1])
		}
	}

	return driveId
}
