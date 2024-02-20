package iotfleetwisecampaign


type IotfleetwiseCampaignDataDestinationConfigs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotfleetwise_campaign#s3_config IotfleetwiseCampaign#s3_config}.
	S3Config *IotfleetwiseCampaignDataDestinationConfigsS3Config `field:"optional" json:"s3Config" yaml:"s3Config"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotfleetwise_campaign#timestream_config IotfleetwiseCampaign#timestream_config}.
	TimestreamConfig *IotfleetwiseCampaignDataDestinationConfigsTimestreamConfig `field:"optional" json:"timestreamConfig" yaml:"timestreamConfig"`
}

