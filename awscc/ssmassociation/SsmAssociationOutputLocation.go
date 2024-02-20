package ssmassociation


type SsmAssociationOutputLocation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ssm_association#s3_location SsmAssociation#s3_location}.
	S3Location *SsmAssociationOutputLocationS3Location `field:"optional" json:"s3Location" yaml:"s3Location"`
}

