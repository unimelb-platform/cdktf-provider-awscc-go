package m2application


type M2ApplicationDefinition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/m2_application#content M2Application#content}.
	Content *string `field:"optional" json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/m2_application#s3_location M2Application#s3_location}.
	S3Location *string `field:"optional" json:"s3Location" yaml:"s3Location"`
}

