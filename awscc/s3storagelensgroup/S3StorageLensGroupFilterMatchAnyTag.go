package s3storagelensgroup


type S3StorageLensGroupFilterMatchAnyTag struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_storage_lens_group#key S3StorageLensGroup#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_storage_lens_group#value S3StorageLensGroup#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

