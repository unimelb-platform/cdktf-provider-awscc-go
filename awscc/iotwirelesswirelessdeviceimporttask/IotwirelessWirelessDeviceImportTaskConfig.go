package iotwirelesswirelessdeviceimporttask

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IotwirelessWirelessDeviceImportTaskConfig struct {
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
	// Destination Name for import task.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotwireless_wireless_device_import_task#destination_name IotwirelessWirelessDeviceImportTask#destination_name}
	DestinationName *string `field:"required" json:"destinationName" yaml:"destinationName"`
	// sidewalk contain file for created device and role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotwireless_wireless_device_import_task#sidewalk IotwirelessWirelessDeviceImportTask#sidewalk}
	Sidewalk *IotwirelessWirelessDeviceImportTaskSidewalk `field:"required" json:"sidewalk" yaml:"sidewalk"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotwireless_wireless_device_import_task#tags IotwirelessWirelessDeviceImportTask#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

