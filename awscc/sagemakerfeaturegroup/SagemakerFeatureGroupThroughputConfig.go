package sagemakerfeaturegroup


type SagemakerFeatureGroupThroughputConfig struct {
	// Throughput mode configuration of the feature group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_feature_group#throughput_mode SagemakerFeatureGroup#throughput_mode}
	ThroughputMode *string `field:"required" json:"throughputMode" yaml:"throughputMode"`
	// For provisioned feature groups with online store enabled, this indicates the read throughput you are billed for and can consume without throttling.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_feature_group#provisioned_read_capacity_units SagemakerFeatureGroup#provisioned_read_capacity_units}
	ProvisionedReadCapacityUnits *float64 `field:"optional" json:"provisionedReadCapacityUnits" yaml:"provisionedReadCapacityUnits"`
	// For provisioned feature groups, this indicates the write throughput you are billed for and can consume without throttling.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_feature_group#provisioned_write_capacity_units SagemakerFeatureGroup#provisioned_write_capacity_units}
	ProvisionedWriteCapacityUnits *float64 `field:"optional" json:"provisionedWriteCapacityUnits" yaml:"provisionedWriteCapacityUnits"`
}

