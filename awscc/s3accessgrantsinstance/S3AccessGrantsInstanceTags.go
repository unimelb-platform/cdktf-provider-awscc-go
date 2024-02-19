package s3accessgrantsinstance


type S3AccessGrantsInstanceTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_access_grants_instance#key S3AccessGrantsInstance#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_access_grants_instance#value S3AccessGrantsInstance#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

