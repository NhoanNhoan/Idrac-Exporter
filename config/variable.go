package config

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stmcginnis/gofish"
)

var (
	// GOFISH is global variable
	GOFISH *gofish.APIClient // gofish/client.go -> APIClient

	// Status map
	Status = map[string]float64{"OK": 0.0}

	// Idracuser info
	Idracuser = "root"
	// Idracpassword info
	Idracpassword = "calvin"

	// Map status -> number
	State_dict = map[string]float64{"OK": 0.0, 
				"WARNING": 1.0, 
				"CRITICAL": 2.0,
	}

	// M_manager => Manager
	M_manager = prometheus.NewDesc(
		"idrac_manager",
		"Manager",
		[]string{
			"auto_dst_enabled",
			"datetime",
			"datetime_local_offset",
			"description",
			"firmware_version",
			"manager_type",
			"manufacturer",
			"model",
			"part_number",
			"power_state",
			"redundancy_count",
			"remote_redfish_service_URI",
			"serial_number",
			"service_entry_point_uuid",
			"status_health",
			"uuid",
			"manager_for_chassis_count",
			"manager_for_server_count",
			"manager_for_switches_count",
		},
		nil,
	)

	// SS_storage => storage service
	SS_storageservice = prometheus.NewDesc(
		"idrac_storage_service",
		"No helpful",
		[]string{
			"description",
			"redundancy_count",
			"status_health",
		},
		nil,
	)
)
