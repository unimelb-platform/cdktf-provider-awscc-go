package kinesisfirehosedeliverystream


type KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationProcessorsParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#parameter_name KinesisfirehoseDeliveryStream#parameter_name}.
	ParameterName *string `field:"required" json:"parameterName" yaml:"parameterName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#parameter_value KinesisfirehoseDeliveryStream#parameter_value}.
	ParameterValue *string `field:"required" json:"parameterValue" yaml:"parameterValue"`
}

