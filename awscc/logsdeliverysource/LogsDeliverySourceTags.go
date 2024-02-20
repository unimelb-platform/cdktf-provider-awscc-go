package logsdeliverysource


type LogsDeliverySourceTags struct {
	// The key name of the tag. You can specify a value that is 1 to 128 Unicode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/logs_delivery_source#key LogsDeliverySource#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value for the tag. You can specify a value that is 0 to 256 Unicode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/logs_delivery_source#value LogsDeliverySource#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

