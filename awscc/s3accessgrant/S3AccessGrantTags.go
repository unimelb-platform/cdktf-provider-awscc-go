package s3accessgrant


type S3AccessGrantTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/s3_access_grant#key S3AccessGrant#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/s3_access_grant#value S3AccessGrant#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

