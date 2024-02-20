package ec2networkinterfaceattachment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2NetworkInterfaceAttachmentConfig struct {
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
	// The network interface's position in the attachment order.
	//
	// For example, the first attached network interface has a DeviceIndex of 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_network_interface_attachment#device_index Ec2NetworkInterfaceAttachment#device_index}
	DeviceIndex *string `field:"required" json:"deviceIndex" yaml:"deviceIndex"`
	// The ID of the instance to which you will attach the ENI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_network_interface_attachment#instance_id Ec2NetworkInterfaceAttachment#instance_id}
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// The ID of the ENI that you want to attach.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_network_interface_attachment#network_interface_id Ec2NetworkInterfaceAttachment#network_interface_id}
	NetworkInterfaceId *string `field:"required" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// Whether to delete the network interface when the instance terminates. By default, this value is set to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_network_interface_attachment#delete_on_termination Ec2NetworkInterfaceAttachment#delete_on_termination}
	DeleteOnTermination interface{} `field:"optional" json:"deleteOnTermination" yaml:"deleteOnTermination"`
}

