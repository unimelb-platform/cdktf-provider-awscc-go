package kinesisfirehosedeliverystream

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KinesisfirehoseDeliveryStreamConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#amazon_open_search_serverless_destination_configuration KinesisfirehoseDeliveryStream#amazon_open_search_serverless_destination_configuration}.
	AmazonOpenSearchServerlessDestinationConfiguration *KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfiguration `field:"optional" json:"amazonOpenSearchServerlessDestinationConfiguration" yaml:"amazonOpenSearchServerlessDestinationConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#amazonopensearchservice_destination_configuration KinesisfirehoseDeliveryStream#amazonopensearchservice_destination_configuration}.
	AmazonopensearchserviceDestinationConfiguration *KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfiguration `field:"optional" json:"amazonopensearchserviceDestinationConfiguration" yaml:"amazonopensearchserviceDestinationConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#delivery_stream_encryption_configuration_input KinesisfirehoseDeliveryStream#delivery_stream_encryption_configuration_input}.
	DeliveryStreamEncryptionConfigurationInput *KinesisfirehoseDeliveryStreamDeliveryStreamEncryptionConfigurationInput `field:"optional" json:"deliveryStreamEncryptionConfigurationInput" yaml:"deliveryStreamEncryptionConfigurationInput"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#delivery_stream_name KinesisfirehoseDeliveryStream#delivery_stream_name}.
	DeliveryStreamName *string `field:"optional" json:"deliveryStreamName" yaml:"deliveryStreamName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#delivery_stream_type KinesisfirehoseDeliveryStream#delivery_stream_type}.
	DeliveryStreamType *string `field:"optional" json:"deliveryStreamType" yaml:"deliveryStreamType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#elasticsearch_destination_configuration KinesisfirehoseDeliveryStream#elasticsearch_destination_configuration}.
	ElasticsearchDestinationConfiguration *KinesisfirehoseDeliveryStreamElasticsearchDestinationConfiguration `field:"optional" json:"elasticsearchDestinationConfiguration" yaml:"elasticsearchDestinationConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#extended_s3_destination_configuration KinesisfirehoseDeliveryStream#extended_s3_destination_configuration}.
	ExtendedS3DestinationConfiguration *KinesisfirehoseDeliveryStreamExtendedS3DestinationConfiguration `field:"optional" json:"extendedS3DestinationConfiguration" yaml:"extendedS3DestinationConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#http_endpoint_destination_configuration KinesisfirehoseDeliveryStream#http_endpoint_destination_configuration}.
	HttpEndpointDestinationConfiguration *KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfiguration `field:"optional" json:"httpEndpointDestinationConfiguration" yaml:"httpEndpointDestinationConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#kinesis_stream_source_configuration KinesisfirehoseDeliveryStream#kinesis_stream_source_configuration}.
	KinesisStreamSourceConfiguration *KinesisfirehoseDeliveryStreamKinesisStreamSourceConfiguration `field:"optional" json:"kinesisStreamSourceConfiguration" yaml:"kinesisStreamSourceConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#msk_source_configuration KinesisfirehoseDeliveryStream#msk_source_configuration}.
	MskSourceConfiguration *KinesisfirehoseDeliveryStreamMskSourceConfiguration `field:"optional" json:"mskSourceConfiguration" yaml:"mskSourceConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#redshift_destination_configuration KinesisfirehoseDeliveryStream#redshift_destination_configuration}.
	RedshiftDestinationConfiguration *KinesisfirehoseDeliveryStreamRedshiftDestinationConfiguration `field:"optional" json:"redshiftDestinationConfiguration" yaml:"redshiftDestinationConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#s3_destination_configuration KinesisfirehoseDeliveryStream#s3_destination_configuration}.
	S3DestinationConfiguration *KinesisfirehoseDeliveryStreamS3DestinationConfiguration `field:"optional" json:"s3DestinationConfiguration" yaml:"s3DestinationConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#snowflake_destination_configuration KinesisfirehoseDeliveryStream#snowflake_destination_configuration}.
	SnowflakeDestinationConfiguration *KinesisfirehoseDeliveryStreamSnowflakeDestinationConfiguration `field:"optional" json:"snowflakeDestinationConfiguration" yaml:"snowflakeDestinationConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#splunk_destination_configuration KinesisfirehoseDeliveryStream#splunk_destination_configuration}.
	SplunkDestinationConfiguration *KinesisfirehoseDeliveryStreamSplunkDestinationConfiguration `field:"optional" json:"splunkDestinationConfiguration" yaml:"splunkDestinationConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#tags KinesisfirehoseDeliveryStream#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

