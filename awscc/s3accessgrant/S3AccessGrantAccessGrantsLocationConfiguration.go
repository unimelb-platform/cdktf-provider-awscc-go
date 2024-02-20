package s3accessgrant


type S3AccessGrantAccessGrantsLocationConfiguration struct {
	// The S3 sub prefix of a registered location in your S3 Access Grants instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_access_grant#s3_sub_prefix S3AccessGrant#s3_sub_prefix}
	S3SubPrefix *string `field:"required" json:"s3SubPrefix" yaml:"s3SubPrefix"`
}

