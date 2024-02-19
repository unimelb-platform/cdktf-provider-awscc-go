package sagemakerproject


type SagemakerProjectServiceCatalogProvisioningDetailsProvisioningParameters struct {
	// The parameter key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_project#key SagemakerProject#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The parameter value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_project#value SagemakerProject#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

