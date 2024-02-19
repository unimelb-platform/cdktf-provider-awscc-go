package kinesisfirehosedeliverystream


type KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationRequestConfigurationCommonAttributes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#attribute_name KinesisfirehoseDeliveryStream#attribute_name}.
	AttributeName *string `field:"required" json:"attributeName" yaml:"attributeName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#attribute_value KinesisfirehoseDeliveryStream#attribute_value}.
	AttributeValue *string `field:"required" json:"attributeValue" yaml:"attributeValue"`
}

