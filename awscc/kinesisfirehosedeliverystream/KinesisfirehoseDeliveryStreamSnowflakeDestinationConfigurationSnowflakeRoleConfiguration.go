package kinesisfirehosedeliverystream


type KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationSnowflakeRoleConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#enabled KinesisfirehoseDeliveryStream#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#snowflake_role KinesisfirehoseDeliveryStream#snowflake_role}.
	SnowflakeRole *string `field:"optional" json:"snowflakeRole" yaml:"snowflakeRole"`
}

