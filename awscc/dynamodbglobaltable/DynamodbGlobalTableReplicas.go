package dynamodbglobaltable


type DynamodbGlobalTableReplicas struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_global_table#region DynamodbGlobalTable#region}.
	Region *string `field:"required" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_global_table#contributor_insights_specification DynamodbGlobalTable#contributor_insights_specification}.
	ContributorInsightsSpecification *DynamodbGlobalTableReplicasContributorInsightsSpecification `field:"optional" json:"contributorInsightsSpecification" yaml:"contributorInsightsSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_global_table#deletion_protection_enabled DynamodbGlobalTable#deletion_protection_enabled}.
	DeletionProtectionEnabled interface{} `field:"optional" json:"deletionProtectionEnabled" yaml:"deletionProtectionEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_global_table#global_secondary_indexes DynamodbGlobalTable#global_secondary_indexes}.
	GlobalSecondaryIndexes interface{} `field:"optional" json:"globalSecondaryIndexes" yaml:"globalSecondaryIndexes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_global_table#kinesis_stream_specification DynamodbGlobalTable#kinesis_stream_specification}.
	KinesisStreamSpecification *DynamodbGlobalTableReplicasKinesisStreamSpecification `field:"optional" json:"kinesisStreamSpecification" yaml:"kinesisStreamSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_global_table#point_in_time_recovery_specification DynamodbGlobalTable#point_in_time_recovery_specification}.
	PointInTimeRecoverySpecification *DynamodbGlobalTableReplicasPointInTimeRecoverySpecification `field:"optional" json:"pointInTimeRecoverySpecification" yaml:"pointInTimeRecoverySpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_global_table#read_provisioned_throughput_settings DynamodbGlobalTable#read_provisioned_throughput_settings}.
	ReadProvisionedThroughputSettings *DynamodbGlobalTableReplicasReadProvisionedThroughputSettings `field:"optional" json:"readProvisionedThroughputSettings" yaml:"readProvisionedThroughputSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_global_table#sse_specification DynamodbGlobalTable#sse_specification}.
	SseSpecification *DynamodbGlobalTableReplicasSseSpecification `field:"optional" json:"sseSpecification" yaml:"sseSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_global_table#table_class DynamodbGlobalTable#table_class}.
	TableClass *string `field:"optional" json:"tableClass" yaml:"tableClass"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_global_table#tags DynamodbGlobalTable#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

