package config

import "github.com/prometheus/client_golang/prometheus"

var (
	// C_powercontrol => Chassis Power Control Metric
	C_powercontrol = prometheus.NewDesc(
		"idrac_power_control",
		"Power Control",
		[]string{
			"physical_context",
			"power_allocated_watts",
			"power_available_watts",
			"power_capacity_watts",
			"power_consumed_watts",
			"power_request_watts",
		},
		nil,
	)

	// C_powerconsumedbyall => Chassis Power Control Metric
	C_powerconsumedbyall = prometheus.NewDesc(
		"idrac_power_consumed_by_all",
		"Power Consumed By All",
		[]string{
			"capacity",
			"id",
		},
		nil,
	)

	// C_powerconsumedbyeach => Chassis Power Control Metric
	C_powerconsumedbyeach = prometheus.NewDesc(
		"idrac_power_consumed_by_each",
		"Power Consumed By each",
		[]string{
			"capacity",
			"id",
			"model",
			"location",
		},
		nil,
	)

	// C_powersupplystatus => Chassis Power Control Metric
	C_powersupplystatus = prometheus.NewDesc(
		"idrac_power_supply_status",
		"Power Supply status {0: OK, 1: Warning, 2: Critical}",
		[]string{
			"firmware_version",
			"last_power_output_watts",
			"manufacturer",
			"member_id",
			"model",
			"part_number",
			"power_capacity_watts",
			"power_input_watts",
			"power_output_watts",
			"power_supply_type",
			"serial_number",
			"spare_part_number",
		},
		nil,
	)

	C_power_line_input_voltage = prometheus.NewDesc(
		"idrac_power_line_input_voltage",
		"Power Line Input Voltage",
		[]string{
			"member_id",
			"line_input_voltage_type",
		},
		nil)

	// C_fan_status => Chassis fan status
	C_fan_status = prometheus.NewDesc(
		"idrac_fan_status",
		"Chassis fan {0: OK, 1: Warning, 2: Critical}",
		[]string{
			"indicator_LED",
			"lower_threshold_critical",
			"lower_threshold_fatal",
			"lower_threshold_non_critical",
			"manufacturer",
			"max_reading_range",
			"member_id",
			"min_reading_range",
			"model",
			"part_number",
			"physical_context",
			"reading",
			"reading_units",
			"redundancy_count",
			"sensor_number",
			"serial_number",
			"spare_part_number",
			"upper_threshold_critical",
			"upper_threshold_fatal",
			"upper_threshold_non_critical",
		},
		nil,
	)

	// C_fan_reading => Chassis fan reading
	C_fan_reading = prometheus.NewDesc(
		"idrac_fan_reading",
		"Chassis fan reading",
		[]string{
			"lower_threshold_critical",
			"lower_threshold_fatal",
			"lower_threshold_non_critical",
			"manufacturer",
			"max_reading_range",
			"member_id",
			"min_reading_range",
			"model",
			"part_number",
			"physical_context",
			"reading_units",
			"sensor_number",
			"serial_number",
			"spare_part_number",
			"upper_threshold_critical",
			"upper_threshold_fatal",
			"upper_threshold_non_critical",
		},
		nil,
	)

	C_temperature_reading = prometheus.NewDesc(
		"idrac_temperature_reading",
		"Chassis temperature {0: OK, 1: Warning, 2: Critical}",
		[]string{
			"adjusted_max_allowable_operating_value",
			"adjusted_min_allowable_operating_value",
			"delta_physical_context",
			"delta_reading_celsius",
			"lower_threshold_critical",
			"lower_threshold_fatal",
			"lower_threshold_non_critical",
			"lower_threshold_user",
			"max_allowable_operating_value",
			"max_reading_range_temp",
			"member_id",
			"min_allowable_operating_value",
			"min_reading_range_temp",
			"physical_context",
			"sensor_number",
			"status_health",
			"upper_threshold_critical",
			"upper_threshold_fatal",
			"upper_threshold_non_critical",
			"upper_threshold_user",
		},
		nil,
	)

	// C_temperature_status => Chassis temperature status
	C_temperature_status = prometheus.NewDesc(
		"idrac_temperature_status",
		"Chassis temperature {0: OK, 1: Warning, 2: Critical}",
		[]string{
			"adjusted_max_allowable_operating_value",
			"adjusted_min_allowable_operating_value",
			"delta_physical_context",
			"delta_reading_celsius",
			"lower_threshold_critical",
			"lower_threshold_fatal",
			"lower_threshold_non_critical",
			"lower_threshold_user",
			"max_allowable_operating_value",
			"max_reading_range_temp",
			"member_id",
			"min_allowable_operating_value",
			"min_reading_range_temp",
			"physical_context",
			"reading_celsius",
			"sensor_number",
			"status_health",
			"upper_threshold_critical",
			"upper_threshold_fatal",
			"upper_threshold_non_critical",
			"upper_threshold_user",
		},
		nil,
	)

	// C_networkadapter => network adapter of the chassis
	C_networkadapter = prometheus.NewDesc(
		"idrac_chassis_network_adapter",
		"Chassis network adapter {0: OK, 1: Warning, 2: Critical}",
		[]string{
			"description",
			"manufacturer",
			"model",
			"part_number",
			"sku",
			"serial_number",
		},
		nil,
	)
)
