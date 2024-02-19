package s3bucket


type S3BucketReplicationConfigurationRulesDestinationMetricsEventThreshold struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#minutes S3Bucket#minutes}.
	Minutes *float64 `field:"required" json:"minutes" yaml:"minutes"`
}

