package workspacesthinclientenvironment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WorkspacesthinclientEnvironmentConfig struct {
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
	// The Amazon Resource Name (ARN) of the desktop to stream from Amazon WorkSpaces, WorkSpaces Web, or AppStream 2.0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/workspacesthinclient_environment#desktop_arn WorkspacesthinclientEnvironment#desktop_arn}
	DesktopArn *string `field:"required" json:"desktopArn" yaml:"desktopArn"`
	// The ID of the software set to apply.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/workspacesthinclient_environment#desired_software_set_id WorkspacesthinclientEnvironment#desired_software_set_id}
	DesiredSoftwareSetId *string `field:"optional" json:"desiredSoftwareSetId" yaml:"desiredSoftwareSetId"`
	// The URL for the identity provider login (only for environments that use AppStream 2.0).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/workspacesthinclient_environment#desktop_endpoint WorkspacesthinclientEnvironment#desktop_endpoint}
	DesktopEndpoint *string `field:"optional" json:"desktopEndpoint" yaml:"desktopEndpoint"`
	// The Amazon Resource Name (ARN) of the AWS Key Management Service key used to encrypt the environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/workspacesthinclient_environment#kms_key_arn WorkspacesthinclientEnvironment#kms_key_arn}
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// A specification for a time window to apply software updates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/workspacesthinclient_environment#maintenance_window WorkspacesthinclientEnvironment#maintenance_window}
	MaintenanceWindow *WorkspacesthinclientEnvironmentMaintenanceWindow `field:"optional" json:"maintenanceWindow" yaml:"maintenanceWindow"`
	// The name of the environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/workspacesthinclient_environment#name WorkspacesthinclientEnvironment#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An option to define which software updates to apply.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/workspacesthinclient_environment#software_set_update_mode WorkspacesthinclientEnvironment#software_set_update_mode}
	SoftwareSetUpdateMode *string `field:"optional" json:"softwareSetUpdateMode" yaml:"softwareSetUpdateMode"`
	// An option to define if software updates should be applied within a maintenance window.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/workspacesthinclient_environment#software_set_update_schedule WorkspacesthinclientEnvironment#software_set_update_schedule}
	SoftwareSetUpdateSchedule *string `field:"optional" json:"softwareSetUpdateSchedule" yaml:"softwareSetUpdateSchedule"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/workspacesthinclient_environment#tags WorkspacesthinclientEnvironment#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

