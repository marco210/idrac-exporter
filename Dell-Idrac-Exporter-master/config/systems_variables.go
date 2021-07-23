package config

import "github.com/prometheus/client_golang/prometheus"

var (
	// S_health => System Health Metric
	S_health = prometheus.NewDesc(
		"idrac_system_health_status",
		"idrac_system_health {0: OK, 1: Warning, 2: Critical}",
		[]string{
			"bios_version",
			"description",
			"hostname",
			"hosted_services",
			"manufacturer",
			"model",
			"name",
			"part_number",
			"power_restore_policy",
			"power_state",
			"sku",
			"serial_number",
			"submodel",
			"system_type",
			"uuid",
		},
		nil,
	)

	// S_memory => system's memory
	S_memory = prometheus.NewDesc(
		"idrac_system_memory_status",
		"System Memory {0: OK, 1: Warning, 2: Critical}",
		[]string{
			"allocation_alignment_MiB",
			"allocation_increment_MiB",
			"base_module_type",
			"bus_width_bits",
			"cache_size_MiB",
			"capacity_MiB",
			"configuration_locked",
			"data_width_bits",
			"description",
			"device_locator",
			"error_correction",
			"firmware_api_version",
			"firmware_revision",
			"is_rank_square_enabled",
			"is_square_device_enabled",
			"logical_size_MiB",
			"manufacturer",
			"memory_device_type",
			"memory_type",
			"operating_speed_Mhz",
			"part_number",
			"rank_count",
			"serial_number",
		},
		nil,
	)

	// S_ethernetinterface => system's ethernet interface
	S_ethernetinterface = prometheus.NewDesc(
		"idrac_system_ethernet_interface_status",
		"System Ethernet Interface{0: OK, 1: Warning, 2: Critical}",
		[]string{
			"auto_negotiation",
			"description",
			"ethernet_interface_type",
			"fqdn",
			"full_duplex",
			"host_name",
			"mac_address",
			"mtu_size",
			"speed_Mbps",
		},
		nil,
	)

	// S_processor => system's processor
	S_processor = prometheus.NewDesc(
		"idrac_system_processor_status",
		"System processor {0: OK, 1:Warning, 2: Critical}",
		[]string{
			"actions",
			"description",
			"manufacturer",
			"max_speed_MHz",
			"max_td_watts",
			"model",
			"processor_type",
			"socket",
			"sub_processors",
			"td_watts",
			"total_cores",
			"total_enabled_cores",
			"total_threads",
			"uuid",
		},
		nil,
	)

	// S_storage => systems' storage
	S_storage = prometheus.NewDesc(
		"idrac_system_storage",
		"System storage {0: OK, 1: Warning, 2: Critical}",
		[]string{
			"description",
			"drives_count",
			"redundancy_count",
			"EnclosuresCount",
		},
		nil,
		)

	// S_storage_drive => computer system -> storage -> drive
	S_storage_drive = prometheus.NewDesc(
		"idrac_storage_drive_status",
		"System storage drive {0: OK, 1: Warning, 2: Critical}",
		[]string{
			"block_size_bytes",
			"capable_speed_gbs",
			"capacity",
			"description",
			"indicator_led",
			"manufacturer",
			"media_type",
			"model",
			"part_number",
			"protocol",
			"revision",
			"serial_number",
		},
		nil,
		)

	S_storage_drive_predicted_media_life_left_percent = prometheus.NewDesc(
		"idrac_ssd_drive_predicted_media_life_left_percent",
		"System storage ssd drive predicted media life left percent",
		[]string{
			"block_size_bytes",
			"capable_speed_gbs",
			"capacity",
			"description",
			"manufacturer",
			"mediatype",
			"model",
			"part_number",
			"protocol",
			"revision",
			"serial_number",
		},
		nil,
	)

	S_storage_volume = prometheus.NewDesc(
		"idrac_storage_volume_status",
		"Storage volume status",
		[]string{
			"description",
			"capacity",
			"volume_type",
			"encrypted",
			"block_size_bytes",
			"drives_count",
			"associated_drives_id",
		},
		nil,
		)

	// S_bios => system's bios
	S_bios = prometheus.NewDesc(
		"idrac_system_bios",
		"System bios",
		[]string{
			"attribute_registry",
			"description",
		},
		nil,
	)

	// S_networkport => system's network port
	S_networkport = prometheus.NewDesc(
		"system_network_port_status",
		"System Network Port",
		[]string{
			"adapter_manufacturer",
			"link_status",
			"current_link_speed_mbps",
			"description",
			"max_frame_size",
			"number_discovered_remote_ports",
			"physical_port_number",
			"port_maximum_mtu",
		},
		nil,
		)



	// S_network_adapter_status => system network adapter status
	S_network_adapter_status = prometheus.NewDesc(
		"idrac_network_adapter_status",
		"System Controller {0: OK, 1: Warning, 2: Critical}",
		[]string{
			"network_adapter_manufacture",
			"firmware_package_version",
			"network_device_functions_count",
			"network_ports_count",
		},
		nil,
		)
)
