package secretsmanagerrotationschedule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SecretsmanagerRotationScheduleConfig struct {
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
	// The ARN or name of the secret to rotate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/secretsmanager_rotation_schedule#secret_id SecretsmanagerRotationSchedule#secret_id}
	SecretId *string `field:"required" json:"secretId" yaml:"secretId"`
	// Creates a new Lambda rotation function based on one of the Secrets Manager rotation function templates.
	//
	// To use a rotation function that already exists, specify RotationLambdaARN instead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/secretsmanager_rotation_schedule#hosted_rotation_lambda SecretsmanagerRotationSchedule#hosted_rotation_lambda}
	HostedRotationLambda *SecretsmanagerRotationScheduleHostedRotationLambda `field:"optional" json:"hostedRotationLambda" yaml:"hostedRotationLambda"`
	// Specifies whether to rotate the secret immediately or wait until the next scheduled rotation window.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/secretsmanager_rotation_schedule#rotate_immediately_on_update SecretsmanagerRotationSchedule#rotate_immediately_on_update}
	RotateImmediatelyOnUpdate interface{} `field:"optional" json:"rotateImmediatelyOnUpdate" yaml:"rotateImmediatelyOnUpdate"`
	// The ARN of an existing Lambda rotation function.
	//
	// To specify a rotation function that is also defined in this template, use the Ref function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/secretsmanager_rotation_schedule#rotation_lambda_arn SecretsmanagerRotationSchedule#rotation_lambda_arn}
	RotationLambdaArn *string `field:"optional" json:"rotationLambdaArn" yaml:"rotationLambdaArn"`
	// A structure that defines the rotation configuration for this secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/secretsmanager_rotation_schedule#rotation_rules SecretsmanagerRotationSchedule#rotation_rules}
	RotationRules *SecretsmanagerRotationScheduleRotationRules `field:"optional" json:"rotationRules" yaml:"rotationRules"`
}

