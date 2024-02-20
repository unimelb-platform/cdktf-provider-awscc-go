package s3bucket


type S3BucketOwnershipControlsRules struct {
	// Specifies an object ownership rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#object_ownership S3Bucket#object_ownership}
	ObjectOwnership *string `field:"optional" json:"objectOwnership" yaml:"objectOwnership"`
}

