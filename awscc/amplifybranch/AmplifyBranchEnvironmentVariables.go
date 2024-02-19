package amplifybranch


type AmplifyBranchEnvironmentVariables struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/amplify_branch#name AmplifyBranch#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/amplify_branch#value AmplifyBranch#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

