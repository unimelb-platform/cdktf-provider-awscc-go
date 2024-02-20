package lambdafunction


type LambdaFunctionFileSystemConfigs struct {
	// The Amazon Resource Name (ARN) of the Amazon EFS access point that provides access to the file system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_function#arn LambdaFunction#arn}
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// The path where the function can access the file system, starting with /mnt/.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_function#local_mount_path LambdaFunction#local_mount_path}
	LocalMountPath *string `field:"required" json:"localMountPath" yaml:"localMountPath"`
}

