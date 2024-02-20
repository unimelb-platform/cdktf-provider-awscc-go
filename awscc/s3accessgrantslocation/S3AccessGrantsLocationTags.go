package s3accessgrantslocation


type S3AccessGrantsLocationTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_access_grants_location#key S3AccessGrantsLocation#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_access_grants_location#value S3AccessGrantsLocation#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

