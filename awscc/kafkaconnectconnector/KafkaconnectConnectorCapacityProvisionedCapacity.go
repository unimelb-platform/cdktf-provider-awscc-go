package kafkaconnectconnector


type KafkaconnectConnectorCapacityProvisionedCapacity struct {
	// Number of workers for a connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kafkaconnect_connector#worker_count KafkaconnectConnector#worker_count}
	WorkerCount *float64 `field:"required" json:"workerCount" yaml:"workerCount"`
	// Specifies how many MSK Connect Units (MCU) are allocated to the connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kafkaconnect_connector#mcu_count KafkaconnectConnector#mcu_count}
	McuCount *float64 `field:"optional" json:"mcuCount" yaml:"mcuCount"`
}

