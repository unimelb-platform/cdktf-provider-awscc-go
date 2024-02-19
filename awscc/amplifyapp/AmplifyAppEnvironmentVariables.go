package amplifyapp


type AmplifyAppEnvironmentVariables struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/amplify_app#name AmplifyApp#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/amplify_app#value AmplifyApp#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

