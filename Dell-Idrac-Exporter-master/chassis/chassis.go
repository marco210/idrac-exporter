package chassis

import (
	"fmt"
	"github.com/alochym01/idrac-exporter/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stmcginnis/gofish/redfish"
	"strings"
)

type Chassis struct{}

func (chassis Chassis) Describe(ch chan<- *prometheus.Desc) {
	ch<- config.C_fan_status
	ch<- config.C_fan_reading
	ch<- config.C_temperature_reading
	ch<- config.C_temperature_status
	ch<- config.C_networkadapter
	ch<- config.C_power_line_input_voltage
}

func (chass Chassis) Collect(ch chan<- prometheus.Metric) {
	metric:=config.GOFISH.Service
	chassisArr, chassisErr := metric.Chassis()

	if nil != chassisErr {
		panic(chassisErr)
	}

	for _, chassis := range chassisArr {
		chass.CollectNetworkAdapter(ch, chassis)

		thermal, thermalErr := chassis.Thermal()
		if nil != thermalErr {
			panic(thermalErr)
		}

		chass.collectTemperature(ch, thermal)
		chass.collectPowerLineInputVoltage(ch, chassis)
	}
}

func (chassis Chassis) CollectNetworkAdapter(ch chan<- prometheus.Metric, chass *redfish.Chassis) {
	adapters, err := chass.NetworkAdapters()

	if nil != err {
		panic(err)
	}

	if 0 != len(adapters) {
		for _, adapter := range adapters {
			status := config.State_dict[string(adapter.Status.Health)]
			ch<- prometheus.MustNewConstMetric(config.C_networkadapter,
				prometheus.GaugeValue,
				float64(status),
				adapter.Description,
				adapter.Manufacturer,
				adapter.Model,
				adapter.PartNumber,
				adapter.SKU,
				adapter.SerialNumber,
			)
		}
	}
}

func (chass Chassis) collectTemperature(ch chan<- prometheus.Metric, thermal *redfish.Thermal) {
	if nil != thermal {
		for _, val := range thermal.Temperatures {
			status := config.State_dict[strings.ToUpper(fmt.Sprintf("%v", val.Status.Health))]
			chass.collectTemperatureReading(ch, status, &val)
			chass.collectTemperatureStatus(ch, status, &val)
		}
	}
}

func (chassis Chassis) collectTemperatureReading(ch chan<- prometheus.Metric, status float64, val *redfish.Temperature) {
	ch <- prometheus.MustNewConstMetric(
		config.C_temperature_reading,
		prometheus.GaugeValue,
		float64(val.ReadingCelsius),
		fmt.Sprintf("%v", val.AdjustedMaxAllowableOperatingValue),
		fmt.Sprintf("%v", val.AdjustedMinAllowableOperatingValue),
		val.DeltaPhysicalContext,
		fmt.Sprintf("%v", val.DeltaReadingCelsius),
		fmt.Sprintf("%v", val.LowerThresholdCritical),
		fmt.Sprintf("%v", val.LowerThresholdFatal),
		fmt.Sprintf("%v", val.LowerThresholdNonCritical),
		fmt.Sprintf("%v", val.LowerThresholdUser),
		fmt.Sprintf("%v", val.MaxAllowableOperatingValue),
		fmt.Sprintf("%v", val.MaxReadingRangeTemp),
		val.MemberID,
		fmt.Sprintf("%v", val.MinAllowableOperatingValue),
		fmt.Sprintf("%v", val.MinReadingRangeTemp),
		val.PhysicalContext,
		fmt.Sprintf("%v", val.SensorNumber),
		fmt.Sprintf("%v", status),
		fmt.Sprintf("%v", val.UpperThresholdCritical),
		fmt.Sprintf("%v", val.UpperThresholdFatal),
		fmt.Sprintf("%v", val.UpperThresholdNonCritical),
		fmt.Sprintf("%v", val.UpperThresholdUser),
	)
}

func (chassis Chassis) collectTemperatureStatus(ch chan<- prometheus.Metric, status float64, val *redfish.Temperature) {
	ch <- prometheus.MustNewConstMetric(
		config.C_temperature_status,
		prometheus.GaugeValue,
		float64(0),
		fmt.Sprintf("%v", val.AdjustedMaxAllowableOperatingValue),
		fmt.Sprintf("%v", val.AdjustedMinAllowableOperatingValue),
		val.DeltaPhysicalContext,
		fmt.Sprintf("%v", val.DeltaReadingCelsius),
		fmt.Sprintf("%v", val.LowerThresholdCritical),
		fmt.Sprintf("%v", val.LowerThresholdFatal),
		fmt.Sprintf("%v", val.LowerThresholdNonCritical),
		fmt.Sprintf("%v", val.LowerThresholdUser),
		fmt.Sprintf("%v", val.MaxAllowableOperatingValue),
		fmt.Sprintf("%v", val.MaxReadingRangeTemp),
		val.MemberID,
		fmt.Sprintf("%v", val.MinAllowableOperatingValue),
		fmt.Sprintf("%v", val.MinReadingRangeTemp),
		val.PhysicalContext,
		fmt.Sprintf("%v", val.ReadingCelsius),
		fmt.Sprintf("%v", val.SensorNumber),
		fmt.Sprintf("%v", status),
		fmt.Sprintf("%v", val.UpperThresholdCritical),
		fmt.Sprintf("%v", val.UpperThresholdFatal),
		fmt.Sprintf("%v", val.UpperThresholdNonCritical),
		fmt.Sprintf("%v", val.UpperThresholdUser),
	)
}

func (chasiss Chassis) collectPowerLineInputVoltage(ch chan<- prometheus.Metric, chass *redfish.Chassis) {
	powers, _ := chass.Power()

	if nil != powers {
		supplies := powers.PowerSupplies

		for _, supply := range supplies {
			ch <- prometheus.MustNewConstMetric(config.C_power_line_input_voltage,
				prometheus.GaugeValue,
				float64(supply.LineInputVoltage),
				supply.MemberID,
				fmt.Sprintf("%v", supply.LineInputVoltageType),
			)
		}

	}
}