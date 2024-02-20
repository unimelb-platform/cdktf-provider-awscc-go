package s3bucket


type S3BucketIntelligentTieringConfigurationsTierings struct {
	// S3 Intelligent-Tiering access tier.
	//
	// See Storage class for automatically optimizing frequently and infrequently accessed objects for a list of access tiers in the S3 Intelligent-Tiering storage class.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#access_tier S3Bucket#access_tier}
	AccessTier *string `field:"required" json:"accessTier" yaml:"accessTier"`
	// The number of consecutive days of no access after which an object will be eligible to be transitioned to the corresponding tier.
	//
	// The minimum number of days specified for Archive Access tier must be at least 90 days and Deep Archive Access tier must be at least 180 days. The maximum can be up to 2 years (730 days).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#days S3Bucket#days}
	Days *float64 `field:"required" json:"days" yaml:"days"`
}

