package dynamodbtable

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DynamodbTableConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_table#key_schema DynamodbTable#key_schema}.
	KeySchema *string `field:"required" json:"keySchema" yaml:"keySchema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_table#attribute_definitions DynamodbTable#attribute_definitions}.
	AttributeDefinitions interface{} `field:"optional" json:"attributeDefinitions" yaml:"attributeDefinitions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_table#billing_mode DynamodbTable#billing_mode}.
	BillingMode *string `field:"optional" json:"billingMode" yaml:"billingMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_table#contributor_insights_specification DynamodbTable#contributor_insights_specification}.
	ContributorInsightsSpecification *DynamodbTableContributorInsightsSpecification `field:"optional" json:"contributorInsightsSpecification" yaml:"contributorInsightsSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_table#deletion_protection_enabled DynamodbTable#deletion_protection_enabled}.
	DeletionProtectionEnabled interface{} `field:"optional" json:"deletionProtectionEnabled" yaml:"deletionProtectionEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_table#global_secondary_indexes DynamodbTable#global_secondary_indexes}.
	GlobalSecondaryIndexes interface{} `field:"optional" json:"globalSecondaryIndexes" yaml:"globalSecondaryIndexes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_table#import_source_specification DynamodbTable#import_source_specification}.
	ImportSourceSpecification *DynamodbTableImportSourceSpecification `field:"optional" json:"importSourceSpecification" yaml:"importSourceSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_table#kinesis_stream_specification DynamodbTable#kinesis_stream_specification}.
	KinesisStreamSpecification *DynamodbTableKinesisStreamSpecification `field:"optional" json:"kinesisStreamSpecification" yaml:"kinesisStreamSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_table#local_secondary_indexes DynamodbTable#local_secondary_indexes}.
	LocalSecondaryIndexes interface{} `field:"optional" json:"localSecondaryIndexes" yaml:"localSecondaryIndexes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_table#point_in_time_recovery_specification DynamodbTable#point_in_time_recovery_specification}.
	PointInTimeRecoverySpecification *DynamodbTablePointInTimeRecoverySpecification `field:"optional" json:"pointInTimeRecoverySpecification" yaml:"pointInTimeRecoverySpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_table#provisioned_throughput DynamodbTable#provisioned_throughput}.
	ProvisionedThroughput *DynamodbTableProvisionedThroughput `field:"optional" json:"provisionedThroughput" yaml:"provisionedThroughput"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_table#sse_specification DynamodbTable#sse_specification}.
	SseSpecification *DynamodbTableSseSpecification `field:"optional" json:"sseSpecification" yaml:"sseSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_table#stream_specification DynamodbTable#stream_specification}.
	StreamSpecification *DynamodbTableStreamSpecification `field:"optional" json:"streamSpecification" yaml:"streamSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_table#table_class DynamodbTable#table_class}.
	TableClass *string `field:"optional" json:"tableClass" yaml:"tableClass"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_table#table_name DynamodbTable#table_name}.
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_table#tags DynamodbTable#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_table#time_to_live_specification DynamodbTable#time_to_live_specification}.
	TimeToLiveSpecification *DynamodbTableTimeToLiveSpecification `field:"optional" json:"timeToLiveSpecification" yaml:"timeToLiveSpecification"`
}

