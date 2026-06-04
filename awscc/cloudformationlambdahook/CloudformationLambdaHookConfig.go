package cloudformationlambdahook

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudformationLambdaHookConfig struct {
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
	// The typename alias for the hook.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cloudformation_lambda_hook#alias CloudformationLambdaHook#alias}
	Alias *string `field:"required" json:"alias" yaml:"alias"`
	// The execution role ARN assumed by Hooks to invoke Lambda.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cloudformation_lambda_hook#execution_role CloudformationLambdaHook#execution_role}
	ExecutionRole *string `field:"required" json:"executionRole" yaml:"executionRole"`
	// Attribute to specify CloudFormation behavior on hook failure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cloudformation_lambda_hook#failure_mode CloudformationLambdaHook#failure_mode}
	FailureMode *string `field:"required" json:"failureMode" yaml:"failureMode"`
	// Amazon Resource Name (ARN), Partial ARN, name, version, or alias of the Lambda function to invoke with this hook.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cloudformation_lambda_hook#lambda_function CloudformationLambdaHook#lambda_function}
	LambdaFunction *string `field:"required" json:"lambdaFunction" yaml:"lambdaFunction"`
	// Which operations should this Hook run against? Resource changes, stacks or change sets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cloudformation_lambda_hook#target_operations CloudformationLambdaHook#target_operations}
	TargetOperations *[]*string `field:"required" json:"targetOperations" yaml:"targetOperations"`
	// Attribute to specify which stacks this hook applies to or should get invoked for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cloudformation_lambda_hook#hook_status CloudformationLambdaHook#hook_status}
	HookStatus *string `field:"optional" json:"hookStatus" yaml:"hookStatus"`
	// Filters to allow hooks to target specific stack attributes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cloudformation_lambda_hook#stack_filters CloudformationLambdaHook#stack_filters}
	StackFilters *CloudformationLambdaHookStackFilters `field:"optional" json:"stackFilters" yaml:"stackFilters"`
	// Attribute to specify which targets should invoke the hook.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cloudformation_lambda_hook#target_filters CloudformationLambdaHook#target_filters}
	TargetFilters *CloudformationLambdaHookTargetFilters `field:"optional" json:"targetFilters" yaml:"targetFilters"`
}

