package s3bucket


type S3BucketReplicationConfigurationRules struct {
	// Specifies which Amazon S3 bucket to store replicated objects in and their storage class.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#destination S3Bucket#destination}
	Destination *S3BucketReplicationConfigurationRulesDestination `field:"required" json:"destination" yaml:"destination"`
	// Specifies whether the rule is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#status S3Bucket#status}
	Status *string `field:"required" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#delete_marker_replication S3Bucket#delete_marker_replication}.
	DeleteMarkerReplication *S3BucketReplicationConfigurationRulesDeleteMarkerReplication `field:"optional" json:"deleteMarkerReplication" yaml:"deleteMarkerReplication"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#filter S3Bucket#filter}.
	Filter *S3BucketReplicationConfigurationRulesFilter `field:"optional" json:"filter" yaml:"filter"`
	// A unique identifier for the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#id S3Bucket#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// An object key name prefix that identifies the object or objects to which the rule applies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#prefix S3Bucket#prefix}
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#priority S3Bucket#priority}.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// A container that describes additional filters for identifying the source objects that you want to replicate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#source_selection_criteria S3Bucket#source_selection_criteria}
	SourceSelectionCriteria *S3BucketReplicationConfigurationRulesSourceSelectionCriteria `field:"optional" json:"sourceSelectionCriteria" yaml:"sourceSelectionCriteria"`
}

