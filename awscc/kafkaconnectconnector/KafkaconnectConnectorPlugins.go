package kafkaconnectconnector


type KafkaconnectConnectorPlugins struct {
	// Details about a custom plugin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/kafkaconnect_connector#custom_plugin KafkaconnectConnector#custom_plugin}
	CustomPlugin *KafkaconnectConnectorPluginsCustomPlugin `field:"required" json:"customPlugin" yaml:"customPlugin"`
}

