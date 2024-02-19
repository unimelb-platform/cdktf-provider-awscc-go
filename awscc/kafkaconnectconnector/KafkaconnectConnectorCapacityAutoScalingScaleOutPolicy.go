package kafkaconnectconnector


type KafkaconnectConnectorCapacityAutoScalingScaleOutPolicy struct {
	// Specifies the CPU utilization percentage threshold at which connector scale out should trigger.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kafkaconnect_connector#cpu_utilization_percentage KafkaconnectConnector#cpu_utilization_percentage}
	CpuUtilizationPercentage *float64 `field:"required" json:"cpuUtilizationPercentage" yaml:"cpuUtilizationPercentage"`
}

