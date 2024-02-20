package s3bucket


type S3BucketCorsConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#cors_rules S3Bucket#cors_rules}.
	CorsRules interface{} `field:"required" json:"corsRules" yaml:"corsRules"`
}

