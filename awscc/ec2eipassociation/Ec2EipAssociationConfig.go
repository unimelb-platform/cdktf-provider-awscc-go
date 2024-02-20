package ec2eipassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2EipAssociationConfig struct {
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
	// The allocation ID. This is required for EC2-VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_eip_association#allocation_id Ec2EipAssociation#allocation_id}
	AllocationId *string `field:"optional" json:"allocationId" yaml:"allocationId"`
	// The Elastic IP address to associate with the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_eip_association#eip Ec2EipAssociation#eip}
	Eip *string `field:"optional" json:"eip" yaml:"eip"`
	// The ID of the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_eip_association#instance_id Ec2EipAssociation#instance_id}
	InstanceId *string `field:"optional" json:"instanceId" yaml:"instanceId"`
	// The ID of the network interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_eip_association#network_interface_id Ec2EipAssociation#network_interface_id}
	NetworkInterfaceId *string `field:"optional" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// The primary or secondary private IP address to associate with the Elastic IP address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_eip_association#private_ip_address Ec2EipAssociation#private_ip_address}
	PrivateIpAddress *string `field:"optional" json:"privateIpAddress" yaml:"privateIpAddress"`
}

