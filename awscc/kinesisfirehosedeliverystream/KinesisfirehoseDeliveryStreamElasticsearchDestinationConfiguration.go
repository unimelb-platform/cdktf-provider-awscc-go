package kinesisfirehosedeliverystream


type KinesisfirehoseDeliveryStreamElasticsearchDestinationConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#index_name KinesisfirehoseDeliveryStream#index_name}.
	IndexName *string `field:"required" json:"indexName" yaml:"indexName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#role_arn KinesisfirehoseDeliveryStream#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#s3_configuration KinesisfirehoseDeliveryStream#s3_configuration}.
	S3Configuration *KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationS3Configuration `field:"required" json:"s3Configuration" yaml:"s3Configuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#buffering_hints KinesisfirehoseDeliveryStream#buffering_hints}.
	BufferingHints *KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationBufferingHints `field:"optional" json:"bufferingHints" yaml:"bufferingHints"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#cloudwatch_logging_options KinesisfirehoseDeliveryStream#cloudwatch_logging_options}.
	CloudwatchLoggingOptions *KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationCloudwatchLoggingOptions `field:"optional" json:"cloudwatchLoggingOptions" yaml:"cloudwatchLoggingOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#cluster_endpoint KinesisfirehoseDeliveryStream#cluster_endpoint}.
	ClusterEndpoint *string `field:"optional" json:"clusterEndpoint" yaml:"clusterEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#document_id_options KinesisfirehoseDeliveryStream#document_id_options}.
	DocumentIdOptions *KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationDocumentIdOptions `field:"optional" json:"documentIdOptions" yaml:"documentIdOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#domain_arn KinesisfirehoseDeliveryStream#domain_arn}.
	DomainArn *string `field:"optional" json:"domainArn" yaml:"domainArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#index_rotation_period KinesisfirehoseDeliveryStream#index_rotation_period}.
	IndexRotationPeriod *string `field:"optional" json:"indexRotationPeriod" yaml:"indexRotationPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#processing_configuration KinesisfirehoseDeliveryStream#processing_configuration}.
	ProcessingConfiguration *KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationProcessingConfiguration `field:"optional" json:"processingConfiguration" yaml:"processingConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#retry_options KinesisfirehoseDeliveryStream#retry_options}.
	RetryOptions *KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationRetryOptions `field:"optional" json:"retryOptions" yaml:"retryOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#s3_backup_mode KinesisfirehoseDeliveryStream#s3_backup_mode}.
	S3BackupMode *string `field:"optional" json:"s3BackupMode" yaml:"s3BackupMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#type_name KinesisfirehoseDeliveryStream#type_name}.
	TypeName *string `field:"optional" json:"typeName" yaml:"typeName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#vpc_configuration KinesisfirehoseDeliveryStream#vpc_configuration}.
	VpcConfiguration *KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationVpcConfiguration `field:"optional" json:"vpcConfiguration" yaml:"vpcConfiguration"`
}

