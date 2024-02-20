package sagemakermodelpackage


type SagemakerModelPackageValidationSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_package#validation_profiles SagemakerModelPackage#validation_profiles}.
	ValidationProfiles interface{} `field:"required" json:"validationProfiles" yaml:"validationProfiles"`
	// The IAM roles to be used for the validation of the model package.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_package#validation_role SagemakerModelPackage#validation_role}
	ValidationRole *string `field:"required" json:"validationRole" yaml:"validationRole"`
}

