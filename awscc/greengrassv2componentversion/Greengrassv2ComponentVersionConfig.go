package greengrassv2componentversion

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Greengrassv2ComponentVersionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/greengrassv2_component_version#inline_recipe Greengrassv2ComponentVersion#inline_recipe}.
	InlineRecipe *string `field:"optional" json:"inlineRecipe" yaml:"inlineRecipe"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/greengrassv2_component_version#lambda_function Greengrassv2ComponentVersion#lambda_function}.
	LambdaFunction *Greengrassv2ComponentVersionLambdaFunction `field:"optional" json:"lambdaFunction" yaml:"lambdaFunction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/greengrassv2_component_version#tags Greengrassv2ComponentVersion#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

