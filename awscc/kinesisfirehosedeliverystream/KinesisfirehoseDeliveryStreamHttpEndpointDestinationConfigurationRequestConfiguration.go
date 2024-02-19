package kinesisfirehosedeliverystream


type KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationRequestConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#common_attributes KinesisfirehoseDeliveryStream#common_attributes}.
	CommonAttributes interface{} `field:"optional" json:"commonAttributes" yaml:"commonAttributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#content_encoding KinesisfirehoseDeliveryStream#content_encoding}.
	ContentEncoding *string `field:"optional" json:"contentEncoding" yaml:"contentEncoding"`
}

