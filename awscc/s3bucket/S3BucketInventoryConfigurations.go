package s3bucket


type S3BucketInventoryConfigurations struct {
	// Specifies information about where to publish analysis or configuration results for an Amazon S3 bucket and S3 Replication Time Control (S3 RTC).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#destination S3Bucket#destination}
	Destination *S3BucketInventoryConfigurationsDestination `field:"required" json:"destination" yaml:"destination"`
	// Specifies whether the inventory is enabled or disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#enabled S3Bucket#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The ID used to identify the inventory configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#id S3Bucket#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Object versions to include in the inventory list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#included_object_versions S3Bucket#included_object_versions}
	IncludedObjectVersions *string `field:"required" json:"includedObjectVersions" yaml:"includedObjectVersions"`
	// Specifies the schedule for generating inventory results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#schedule_frequency S3Bucket#schedule_frequency}
	ScheduleFrequency *string `field:"required" json:"scheduleFrequency" yaml:"scheduleFrequency"`
	// Contains the optional fields that are included in the inventory results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#optional_fields S3Bucket#optional_fields}
	OptionalFields *[]*string `field:"optional" json:"optionalFields" yaml:"optionalFields"`
	// The prefix that is prepended to all inventory results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#prefix S3Bucket#prefix}
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

