package kafkaconnectconnector


type KafkaconnectConnectorKafkaClusterApacheKafkaCluster struct {
	// The bootstrap servers string of the Apache Kafka cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kafkaconnect_connector#bootstrap_servers KafkaconnectConnector#bootstrap_servers}
	BootstrapServers *string `field:"required" json:"bootstrapServers" yaml:"bootstrapServers"`
	// Information about a VPC used with the connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kafkaconnect_connector#vpc KafkaconnectConnector#vpc}
	Vpc *KafkaconnectConnectorKafkaClusterApacheKafkaClusterVpc `field:"required" json:"vpc" yaml:"vpc"`
}

