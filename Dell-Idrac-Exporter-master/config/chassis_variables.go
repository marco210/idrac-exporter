package config

import "github.com/prometheus/client_golang/prometheus"

var (

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
