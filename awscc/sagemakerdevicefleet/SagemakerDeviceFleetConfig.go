package sagemakerdevicefleet

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SagemakerDeviceFleetConfig struct {
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
	// The name of the edge device fleet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_device_fleet#device_fleet_name SagemakerDeviceFleet#device_fleet_name}
	DeviceFleetName *string `field:"required" json:"deviceFleetName" yaml:"deviceFleetName"`
	// S3 bucket and an ecryption key id (if available) to store outputs for the fleet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_device_fleet#output_config SagemakerDeviceFleet#output_config}
	OutputConfig *SagemakerDeviceFleetOutputConfig `field:"required" json:"outputConfig" yaml:"outputConfig"`
	// Role associated with the device fleet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_device_fleet#role_arn SagemakerDeviceFleet#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Description for the edge device fleet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_device_fleet#description SagemakerDeviceFleet#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Associate tags with the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_device_fleet#tags SagemakerDeviceFleet#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

