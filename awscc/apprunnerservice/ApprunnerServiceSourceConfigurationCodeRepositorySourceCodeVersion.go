package apprunnerservice


type ApprunnerServiceSourceConfigurationCodeRepositorySourceCodeVersion struct {
	// Source Code Version Type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apprunner_service#type ApprunnerService#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Source Code Version Value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apprunner_service#value ApprunnerService#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

