package amplifyapp

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AmplifyAppConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/amplify_app#name AmplifyApp#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/amplify_app#access_token AmplifyApp#access_token}.
	AccessToken *string `field:"optional" json:"accessToken" yaml:"accessToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/amplify_app#auto_branch_creation_config AmplifyApp#auto_branch_creation_config}.
	AutoBranchCreationConfig *AmplifyAppAutoBranchCreationConfig `field:"optional" json:"autoBranchCreationConfig" yaml:"autoBranchCreationConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/amplify_app#basic_auth_config AmplifyApp#basic_auth_config}.
	BasicAuthConfig *AmplifyAppBasicAuthConfig `field:"optional" json:"basicAuthConfig" yaml:"basicAuthConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/amplify_app#build_spec AmplifyApp#build_spec}.
	BuildSpec *string `field:"optional" json:"buildSpec" yaml:"buildSpec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/amplify_app#custom_headers AmplifyApp#custom_headers}.
	CustomHeaders *string `field:"optional" json:"customHeaders" yaml:"customHeaders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/amplify_app#custom_rules AmplifyApp#custom_rules}.
	CustomRules interface{} `field:"optional" json:"customRules" yaml:"customRules"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/amplify_app#description AmplifyApp#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/amplify_app#enable_branch_auto_deletion AmplifyApp#enable_branch_auto_deletion}.
	EnableBranchAutoDeletion interface{} `field:"optional" json:"enableBranchAutoDeletion" yaml:"enableBranchAutoDeletion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/amplify_app#environment_variables AmplifyApp#environment_variables}.
	EnvironmentVariables interface{} `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/amplify_app#iam_service_role AmplifyApp#iam_service_role}.
	IamServiceRole *string `field:"optional" json:"iamServiceRole" yaml:"iamServiceRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/amplify_app#oauth_token AmplifyApp#oauth_token}.
	OauthToken *string `field:"optional" json:"oauthToken" yaml:"oauthToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/amplify_app#platform AmplifyApp#platform}.
	Platform *string `field:"optional" json:"platform" yaml:"platform"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/amplify_app#repository AmplifyApp#repository}.
	Repository *string `field:"optional" json:"repository" yaml:"repository"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/amplify_app#tags AmplifyApp#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

