package lambdapermission

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LambdaPermissionConfig struct {
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
	// The action that the principal can use on the function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_permission#action LambdaPermission#action}
	Action *string `field:"required" json:"action" yaml:"action"`
	// The name of the Lambda function, version, or alias.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_permission#function_name LambdaPermission#function_name}
	FunctionName *string `field:"required" json:"functionName" yaml:"functionName"`
	// The AWS service or account that invokes the function.
	//
	// If you specify a service, use SourceArn or SourceAccount to limit who can invoke the function through that service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_permission#principal LambdaPermission#principal}
	Principal *string `field:"required" json:"principal" yaml:"principal"`
	// For Alexa Smart Home functions, a token that must be supplied by the invoker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_permission#event_source_token LambdaPermission#event_source_token}
	EventSourceToken *string `field:"optional" json:"eventSourceToken" yaml:"eventSourceToken"`
	// The type of authentication that your function URL uses.
	//
	// Set to AWS_IAM if you want to restrict access to authenticated users only. Set to NONE if you want to bypass IAM authentication to create a public endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_permission#function_url_auth_type LambdaPermission#function_url_auth_type}
	FunctionUrlAuthType *string `field:"optional" json:"functionUrlAuthType" yaml:"functionUrlAuthType"`
	// The identifier for your organization in AWS Organizations.
	//
	// Use this to grant permissions to all the AWS accounts under this organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_permission#principal_org_id LambdaPermission#principal_org_id}
	PrincipalOrgId *string `field:"optional" json:"principalOrgId" yaml:"principalOrgId"`
	// For Amazon S3, the ID of the account that owns the resource.
	//
	// Use this together with SourceArn to ensure that the resource is owned by the specified account. It is possible for an Amazon S3 bucket to be deleted by its owner and recreated by another account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_permission#source_account LambdaPermission#source_account}
	SourceAccount *string `field:"optional" json:"sourceAccount" yaml:"sourceAccount"`
	// For AWS services, the ARN of the AWS resource that invokes the function.
	//
	// For example, an Amazon S3 bucket or Amazon SNS topic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_permission#source_arn LambdaPermission#source_arn}
	SourceArn *string `field:"optional" json:"sourceArn" yaml:"sourceArn"`
}

