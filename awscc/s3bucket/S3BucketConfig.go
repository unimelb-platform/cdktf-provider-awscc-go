package s3bucket

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type S3BucketConfig struct {
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
	// Configuration for the transfer acceleration state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#accelerate_configuration S3Bucket#accelerate_configuration}
	AccelerateConfiguration *S3BucketAccelerateConfiguration `field:"optional" json:"accelerateConfiguration" yaml:"accelerateConfiguration"`
	// A canned access control list (ACL) that grants predefined permissions to the bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#access_control S3Bucket#access_control}
	AccessControl *string `field:"optional" json:"accessControl" yaml:"accessControl"`
	// The configuration and any analyses for the analytics filter of an Amazon S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#analytics_configurations S3Bucket#analytics_configurations}
	AnalyticsConfigurations interface{} `field:"optional" json:"analyticsConfigurations" yaml:"analyticsConfigurations"`
	// Specifies default encryption for a bucket using server-side encryption with either Amazon S3-managed keys (SSE-S3) or AWS KMS-managed keys (SSE-KMS).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#bucket_encryption S3Bucket#bucket_encryption}
	BucketEncryption *S3BucketBucketEncryption `field:"optional" json:"bucketEncryption" yaml:"bucketEncryption"`
	// A name for the bucket.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the bucket name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#bucket_name S3Bucket#bucket_name}
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// Rules that define cross-origin resource sharing of objects in this bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#cors_configuration S3Bucket#cors_configuration}
	CorsConfiguration *S3BucketCorsConfiguration `field:"optional" json:"corsConfiguration" yaml:"corsConfiguration"`
	// Specifies the S3 Intelligent-Tiering configuration for an Amazon S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#intelligent_tiering_configurations S3Bucket#intelligent_tiering_configurations}
	IntelligentTieringConfigurations interface{} `field:"optional" json:"intelligentTieringConfigurations" yaml:"intelligentTieringConfigurations"`
	// The inventory configuration for an Amazon S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#inventory_configurations S3Bucket#inventory_configurations}
	InventoryConfigurations interface{} `field:"optional" json:"inventoryConfigurations" yaml:"inventoryConfigurations"`
	// Rules that define how Amazon S3 manages objects during their lifetime.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#lifecycle_configuration S3Bucket#lifecycle_configuration}
	LifecycleConfiguration *S3BucketLifecycleConfiguration `field:"optional" json:"lifecycleConfiguration" yaml:"lifecycleConfiguration"`
	// Settings that define where logs are stored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#logging_configuration S3Bucket#logging_configuration}
	LoggingConfiguration *S3BucketLoggingConfiguration `field:"optional" json:"loggingConfiguration" yaml:"loggingConfiguration"`
	// Settings that define a metrics configuration for the CloudWatch request metrics from the bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#metrics_configurations S3Bucket#metrics_configurations}
	MetricsConfigurations interface{} `field:"optional" json:"metricsConfigurations" yaml:"metricsConfigurations"`
	// Configuration that defines how Amazon S3 handles bucket notifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#notification_configuration S3Bucket#notification_configuration}
	NotificationConfiguration *S3BucketNotificationConfiguration `field:"optional" json:"notificationConfiguration" yaml:"notificationConfiguration"`
	// Places an Object Lock configuration on the specified bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#object_lock_configuration S3Bucket#object_lock_configuration}
	ObjectLockConfiguration *S3BucketObjectLockConfiguration `field:"optional" json:"objectLockConfiguration" yaml:"objectLockConfiguration"`
	// Indicates whether this bucket has an Object Lock configuration enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#object_lock_enabled S3Bucket#object_lock_enabled}
	ObjectLockEnabled interface{} `field:"optional" json:"objectLockEnabled" yaml:"objectLockEnabled"`
	// Specifies the container element for object ownership rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#ownership_controls S3Bucket#ownership_controls}
	OwnershipControls *S3BucketOwnershipControls `field:"optional" json:"ownershipControls" yaml:"ownershipControls"`
	// Configuration that defines how Amazon S3 handles public access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#public_access_block_configuration S3Bucket#public_access_block_configuration}
	PublicAccessBlockConfiguration *S3BucketPublicAccessBlockConfiguration `field:"optional" json:"publicAccessBlockConfiguration" yaml:"publicAccessBlockConfiguration"`
	// Configuration for replicating objects in an S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#replication_configuration S3Bucket#replication_configuration}
	ReplicationConfiguration *S3BucketReplicationConfiguration `field:"optional" json:"replicationConfiguration" yaml:"replicationConfiguration"`
	// An arbitrary set of tags (key-value pairs) for this S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#tags S3Bucket#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Describes the versioning state of an Amazon S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#versioning_configuration S3Bucket#versioning_configuration}
	VersioningConfiguration *S3BucketVersioningConfiguration `field:"optional" json:"versioningConfiguration" yaml:"versioningConfiguration"`
	// Specifies website configuration parameters for an Amazon S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#website_configuration S3Bucket#website_configuration}
	WebsiteConfiguration *S3BucketWebsiteConfiguration `field:"optional" json:"websiteConfiguration" yaml:"websiteConfiguration"`
}

