package comprehendflywheel

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComprehendFlywheelConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/comprehend_flywheel#data_access_role_arn ComprehendFlywheel#data_access_role_arn}.
	DataAccessRoleArn *string `field:"required" json:"dataAccessRoleArn" yaml:"dataAccessRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/comprehend_flywheel#data_lake_s3_uri ComprehendFlywheel#data_lake_s3_uri}.
	DataLakeS3Uri *string `field:"required" json:"dataLakeS3Uri" yaml:"dataLakeS3Uri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/comprehend_flywheel#flywheel_name ComprehendFlywheel#flywheel_name}.
	FlywheelName *string `field:"required" json:"flywheelName" yaml:"flywheelName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/comprehend_flywheel#active_model_arn ComprehendFlywheel#active_model_arn}.
	ActiveModelArn *string `field:"optional" json:"activeModelArn" yaml:"activeModelArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/comprehend_flywheel#data_security_config ComprehendFlywheel#data_security_config}.
	DataSecurityConfig *ComprehendFlywheelDataSecurityConfig `field:"optional" json:"dataSecurityConfig" yaml:"dataSecurityConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/comprehend_flywheel#model_type ComprehendFlywheel#model_type}.
	ModelType *string `field:"optional" json:"modelType" yaml:"modelType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/comprehend_flywheel#tags ComprehendFlywheel#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/comprehend_flywheel#task_config ComprehendFlywheel#task_config}.
	TaskConfig *ComprehendFlywheelTaskConfig `field:"optional" json:"taskConfig" yaml:"taskConfig"`
}

