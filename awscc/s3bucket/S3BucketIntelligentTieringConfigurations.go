package s3bucket


type S3BucketIntelligentTieringConfigurations struct {
	// The ID used to identify the S3 Intelligent-Tiering configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#id S3Bucket#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Specifies the status of the configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#status S3Bucket#status}
	Status *string `field:"required" json:"status" yaml:"status"`
	// Specifies a list of S3 Intelligent-Tiering storage class tiers in the configuration.
	//
	// At least one tier must be defined in the list. At most, you can specify two tiers in the list, one for each available AccessTier: ARCHIVE_ACCESS and DEEP_ARCHIVE_ACCESS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#tierings S3Bucket#tierings}
	Tierings interface{} `field:"required" json:"tierings" yaml:"tierings"`
	// An object key name prefix that identifies the subset of objects to which the rule applies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#prefix S3Bucket#prefix}
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// A container for a key-value pair.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#tag_filters S3Bucket#tag_filters}
	TagFilters interface{} `field:"optional" json:"tagFilters" yaml:"tagFilters"`
}

