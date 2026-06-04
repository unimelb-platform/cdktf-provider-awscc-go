package s3bucket


type S3BucketReplicationConfigurationRulesDeleteMarkerReplication struct {
	// Indicates whether to replicate delete markers. Disabled by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/s3_bucket#status S3Bucket#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
}

