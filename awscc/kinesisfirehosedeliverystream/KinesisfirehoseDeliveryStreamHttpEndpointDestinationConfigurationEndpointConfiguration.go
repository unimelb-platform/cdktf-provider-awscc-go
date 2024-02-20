package kinesisfirehosedeliverystream


type KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationEndpointConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#url KinesisfirehoseDeliveryStream#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#access_key KinesisfirehoseDeliveryStream#access_key}.
	AccessKey *string `field:"optional" json:"accessKey" yaml:"accessKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#name KinesisfirehoseDeliveryStream#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

