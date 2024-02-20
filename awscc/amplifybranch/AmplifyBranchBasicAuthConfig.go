package amplifybranch


type AmplifyBranchBasicAuthConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/amplify_branch#password AmplifyBranch#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/amplify_branch#username AmplifyBranch#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/amplify_branch#enable_basic_auth AmplifyBranch#enable_basic_auth}.
	EnableBasicAuth interface{} `field:"optional" json:"enableBasicAuth" yaml:"enableBasicAuth"`
}

