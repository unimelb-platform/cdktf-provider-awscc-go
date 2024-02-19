package lambdafunction

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LambdaFunctionConfig struct {
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
	// The code for the function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_function#code LambdaFunction#code}
	Code *LambdaFunctionCode `field:"required" json:"code" yaml:"code"`
	// The Amazon Resource Name (ARN) of the function's execution role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_function#role LambdaFunction#role}
	Role *string `field:"required" json:"role" yaml:"role"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_function#architectures LambdaFunction#architectures}.
	Architectures *[]*string `field:"optional" json:"architectures" yaml:"architectures"`
	// A unique Arn for CodeSigningConfig resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_function#code_signing_config_arn LambdaFunction#code_signing_config_arn}
	CodeSigningConfigArn *string `field:"optional" json:"codeSigningConfigArn" yaml:"codeSigningConfigArn"`
	// A dead letter queue configuration that specifies the queue or topic where Lambda sends asynchronous events when they fail processing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_function#dead_letter_config LambdaFunction#dead_letter_config}
	DeadLetterConfig *LambdaFunctionDeadLetterConfig `field:"optional" json:"deadLetterConfig" yaml:"deadLetterConfig"`
	// A description of the function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_function#description LambdaFunction#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Environment variables that are accessible from function code during execution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_function#environment LambdaFunction#environment}
	Environment *LambdaFunctionEnvironment `field:"optional" json:"environment" yaml:"environment"`
	// A function's ephemeral storage settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_function#ephemeral_storage LambdaFunction#ephemeral_storage}
	EphemeralStorage *LambdaFunctionEphemeralStorage `field:"optional" json:"ephemeralStorage" yaml:"ephemeralStorage"`
	// Connection settings for an Amazon EFS file system.
	//
	// To connect a function to a file system, a mount target must be available in every Availability Zone that your function connects to. If your template contains an AWS::EFS::MountTarget resource, you must also specify a DependsOn attribute to ensure that the mount target is created or updated before the function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_function#file_system_configs LambdaFunction#file_system_configs}
	FileSystemConfigs interface{} `field:"optional" json:"fileSystemConfigs" yaml:"fileSystemConfigs"`
	// The name of the Lambda function, up to 64 characters in length.
	//
	// If you don't specify a name, AWS CloudFormation generates one.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_function#function_name LambdaFunction#function_name}
	FunctionName *string `field:"optional" json:"functionName" yaml:"functionName"`
	// The name of the method within your code that Lambda calls to execute your function.
	//
	// The format includes the file name. It can also include namespaces and other qualifiers, depending on the runtime
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_function#handler LambdaFunction#handler}
	Handler *string `field:"optional" json:"handler" yaml:"handler"`
	// ImageConfig.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_function#image_config LambdaFunction#image_config}
	ImageConfig *LambdaFunctionImageConfig `field:"optional" json:"imageConfig" yaml:"imageConfig"`
	// The ARN of the AWS Key Management Service (AWS KMS) key that's used to encrypt your function's environment variables.
	//
	// If it's not provided, AWS Lambda uses a default service key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_function#kms_key_arn LambdaFunction#kms_key_arn}
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// A list of function layers to add to the function's execution environment.
	//
	// Specify each layer by its ARN, including the version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_function#layers LambdaFunction#layers}
	Layers *[]*string `field:"optional" json:"layers" yaml:"layers"`
	// The logging configuration of your function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_function#logging_config LambdaFunction#logging_config}
	LoggingConfig *LambdaFunctionLoggingConfig `field:"optional" json:"loggingConfig" yaml:"loggingConfig"`
	// The amount of memory that your function has access to.
	//
	// Increasing the function's memory also increases its CPU allocation. The default value is 128 MB. The value must be a multiple of 64 MB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_function#memory_size LambdaFunction#memory_size}
	MemorySize *float64 `field:"optional" json:"memorySize" yaml:"memorySize"`
	// PackageType.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_function#package_type LambdaFunction#package_type}
	PackageType *string `field:"optional" json:"packageType" yaml:"packageType"`
	// The number of simultaneous executions to reserve for the function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_function#reserved_concurrent_executions LambdaFunction#reserved_concurrent_executions}
	ReservedConcurrentExecutions *float64 `field:"optional" json:"reservedConcurrentExecutions" yaml:"reservedConcurrentExecutions"`
	// The identifier of the function's runtime.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_function#runtime LambdaFunction#runtime}
	Runtime *string `field:"optional" json:"runtime" yaml:"runtime"`
	// RuntimeManagementConfig.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_function#runtime_management_config LambdaFunction#runtime_management_config}
	RuntimeManagementConfig *LambdaFunctionRuntimeManagementConfig `field:"optional" json:"runtimeManagementConfig" yaml:"runtimeManagementConfig"`
	// The SnapStart setting of your function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_function#snap_start LambdaFunction#snap_start}
	SnapStart *LambdaFunctionSnapStart `field:"optional" json:"snapStart" yaml:"snapStart"`
	// A list of tags to apply to the function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_function#tags LambdaFunction#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The amount of time that Lambda allows a function to run before stopping it.
	//
	// The default is 3 seconds. The maximum allowed value is 900 seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_function#timeout LambdaFunction#timeout}
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
	// Set Mode to Active to sample and trace a subset of incoming requests with AWS X-Ray.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_function#tracing_config LambdaFunction#tracing_config}
	TracingConfig *LambdaFunctionTracingConfig `field:"optional" json:"tracingConfig" yaml:"tracingConfig"`
	// For network connectivity to AWS resources in a VPC, specify a list of security groups and subnets in the VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_function#vpc_config LambdaFunction#vpc_config}
	VpcConfig *LambdaFunctionVpcConfig `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

