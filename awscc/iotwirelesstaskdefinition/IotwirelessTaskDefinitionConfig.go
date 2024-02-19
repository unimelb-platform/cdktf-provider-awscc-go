package iotwirelesstaskdefinition

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IotwirelessTaskDefinitionConfig struct {
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
	// Whether to automatically create tasks using this task definition for all gateways with the specified current version.
	//
	// If false, the task must me created by calling CreateWirelessGatewayTask.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotwireless_task_definition#auto_create_tasks IotwirelessTaskDefinition#auto_create_tasks}
	AutoCreateTasks interface{} `field:"required" json:"autoCreateTasks" yaml:"autoCreateTasks"`
	// The list of task definitions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotwireless_task_definition#lo_ra_wan_update_gateway_task_entry IotwirelessTaskDefinition#lo_ra_wan_update_gateway_task_entry}
	LoRaWanUpdateGatewayTaskEntry *IotwirelessTaskDefinitionLoRaWanUpdateGatewayTaskEntry `field:"optional" json:"loRaWanUpdateGatewayTaskEntry" yaml:"loRaWanUpdateGatewayTaskEntry"`
	// The name of the new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotwireless_task_definition#name IotwirelessTaskDefinition#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of key-value pairs that contain metadata for the destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotwireless_task_definition#tags IotwirelessTaskDefinition#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// A filter to list only the wireless gateway task definitions that use this task definition type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotwireless_task_definition#task_definition_type IotwirelessTaskDefinition#task_definition_type}
	TaskDefinitionType *string `field:"optional" json:"taskDefinitionType" yaml:"taskDefinitionType"`
	// Information about the gateways to update.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotwireless_task_definition#update IotwirelessTaskDefinition#update}
	Update *IotwirelessTaskDefinitionUpdate `field:"optional" json:"update" yaml:"update"`
}

