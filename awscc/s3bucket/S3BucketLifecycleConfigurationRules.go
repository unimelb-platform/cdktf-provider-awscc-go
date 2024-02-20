package s3bucket


type S3BucketLifecycleConfigurationRules struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#status S3Bucket#status}.
	Status *string `field:"required" json:"status" yaml:"status"`
	// Specifies the days since the initiation of an incomplete multipart upload that Amazon S3 will wait before permanently removing all parts of the upload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#abort_incomplete_multipart_upload S3Bucket#abort_incomplete_multipart_upload}
	AbortIncompleteMultipartUpload *S3BucketLifecycleConfigurationRulesAbortIncompleteMultipartUpload `field:"optional" json:"abortIncompleteMultipartUpload" yaml:"abortIncompleteMultipartUpload"`
	// The date value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ssZ).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#expiration_date S3Bucket#expiration_date}
	ExpirationDate *string `field:"optional" json:"expirationDate" yaml:"expirationDate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#expiration_in_days S3Bucket#expiration_in_days}.
	ExpirationInDays *float64 `field:"optional" json:"expirationInDays" yaml:"expirationInDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#expired_object_delete_marker S3Bucket#expired_object_delete_marker}.
	ExpiredObjectDeleteMarker interface{} `field:"optional" json:"expiredObjectDeleteMarker" yaml:"expiredObjectDeleteMarker"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#id S3Bucket#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Container for the expiration rule that describes when noncurrent objects are expired.
	//
	// If your bucket is versioning-enabled (or versioning is suspended), you can set this action to request that Amazon S3 expire noncurrent object versions at a specific period in the object's lifetime
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#noncurrent_version_expiration S3Bucket#noncurrent_version_expiration}
	NoncurrentVersionExpiration *S3BucketLifecycleConfigurationRulesNoncurrentVersionExpiration `field:"optional" json:"noncurrentVersionExpiration" yaml:"noncurrentVersionExpiration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#noncurrent_version_expiration_in_days S3Bucket#noncurrent_version_expiration_in_days}.
	NoncurrentVersionExpirationInDays *float64 `field:"optional" json:"noncurrentVersionExpirationInDays" yaml:"noncurrentVersionExpirationInDays"`
	// Container for the transition rule that describes when noncurrent objects transition to the STANDARD_IA, ONEZONE_IA, INTELLIGENT_TIERING, GLACIER_IR, GLACIER, or DEEP_ARCHIVE storage class.
	//
	// If your bucket is versioning-enabled (or versioning is suspended), you can set this action to request that Amazon S3 transition noncurrent object versions to the STANDARD_IA, ONEZONE_IA, INTELLIGENT_TIERING, GLACIER_IR, GLACIER, or DEEP_ARCHIVE storage class at a specific period in the object's lifetime.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#noncurrent_version_transition S3Bucket#noncurrent_version_transition}
	NoncurrentVersionTransition *S3BucketLifecycleConfigurationRulesNoncurrentVersionTransition `field:"optional" json:"noncurrentVersionTransition" yaml:"noncurrentVersionTransition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#noncurrent_version_transitions S3Bucket#noncurrent_version_transitions}.
	NoncurrentVersionTransitions interface{} `field:"optional" json:"noncurrentVersionTransitions" yaml:"noncurrentVersionTransitions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#object_size_greater_than S3Bucket#object_size_greater_than}.
	ObjectSizeGreaterThan *string `field:"optional" json:"objectSizeGreaterThan" yaml:"objectSizeGreaterThan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#object_size_less_than S3Bucket#object_size_less_than}.
	ObjectSizeLessThan *string `field:"optional" json:"objectSizeLessThan" yaml:"objectSizeLessThan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#prefix S3Bucket#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#tag_filters S3Bucket#tag_filters}.
	TagFilters interface{} `field:"optional" json:"tagFilters" yaml:"tagFilters"`
	// You must specify at least one of "TransitionDate" and "TransitionInDays".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#transition S3Bucket#transition}
	Transition *S3BucketLifecycleConfigurationRulesTransition `field:"optional" json:"transition" yaml:"transition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#transitions S3Bucket#transitions}.
	Transitions interface{} `field:"optional" json:"transitions" yaml:"transitions"`
}

