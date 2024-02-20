package lambdaurl

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LambdaUrlConfig struct {
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
	// Can be either AWS_IAM if the requests are authorized via IAM, or NONE if no authorization is configured on the Function URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_url#auth_type LambdaUrl#auth_type}
	AuthType *string `field:"required" json:"authType" yaml:"authType"`
	// The Amazon Resource Name (ARN) of the function associated with the Function URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_url#target_function_arn LambdaUrl#target_function_arn}
	TargetFunctionArn *string `field:"required" json:"targetFunctionArn" yaml:"targetFunctionArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_url#cors LambdaUrl#cors}.
	Cors *LambdaUrlCors `field:"optional" json:"cors" yaml:"cors"`
	// The invocation mode for the function?s URL. Set to BUFFERED if you want to buffer responses before returning them to the client. Set to RESPONSE_STREAM if you want to stream responses, allowing faster time to first byte and larger response payload sizes. If not set, defaults to BUFFERED.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_url#invoke_mode LambdaUrl#invoke_mode}
	InvokeMode *string `field:"optional" json:"invokeMode" yaml:"invokeMode"`
	// The alias qualifier for the target function. If TargetFunctionArn is unqualified then Qualifier must be passed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_url#qualifier LambdaUrl#qualifier}
	Qualifier *string `field:"optional" json:"qualifier" yaml:"qualifier"`
}

