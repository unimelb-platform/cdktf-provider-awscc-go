package lambdaeventsourcemapping


type LambdaEventSourceMappingAmazonManagedKafkaEventSourceConfigSchemaRegistryConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/lambda_event_source_mapping#access_configs LambdaEventSourceMapping#access_configs}.
	AccessConfigs interface{} `field:"optional" json:"accessConfigs" yaml:"accessConfigs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/lambda_event_source_mapping#event_record_format LambdaEventSourceMapping#event_record_format}.
	EventRecordFormat *string `field:"optional" json:"eventRecordFormat" yaml:"eventRecordFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/lambda_event_source_mapping#schema_registry_uri LambdaEventSourceMapping#schema_registry_uri}.
	SchemaRegistryUri *string `field:"optional" json:"schemaRegistryUri" yaml:"schemaRegistryUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/lambda_event_source_mapping#schema_validation_configs LambdaEventSourceMapping#schema_validation_configs}.
	SchemaValidationConfigs interface{} `field:"optional" json:"schemaValidationConfigs" yaml:"schemaValidationConfigs"`
}

