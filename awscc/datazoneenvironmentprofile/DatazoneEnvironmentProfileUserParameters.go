package datazoneenvironmentprofile


type DatazoneEnvironmentProfileUserParameters struct {
	// The name of an environment profile parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datazone_environment_profile#name DatazoneEnvironmentProfile#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The value of an environment profile parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datazone_environment_profile#value DatazoneEnvironmentProfile#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

