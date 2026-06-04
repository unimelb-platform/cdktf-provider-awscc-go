package ec2eip

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2EipConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_eip#address Ec2Eip#address}.
	Address *string `field:"optional" json:"address" yaml:"address"`
	// The network (``vpc``).
	//
	// If you define an Elastic IP address and associate it with a VPC that is defined in the same template, you must declare a dependency on the VPC-gateway attachment by using the [DependsOn Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html) on this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_eip#domain Ec2Eip#domain}
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// The ID of the instance.
	//
	// Updates to the ``InstanceId`` property may require *some interruptions*. Updates on an EIP reassociates the address on its associated resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_eip#instance_id Ec2Eip#instance_id}
	InstanceId *string `field:"optional" json:"instanceId" yaml:"instanceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_eip#ipam_pool_id Ec2Eip#ipam_pool_id}.
	IpamPoolId *string `field:"optional" json:"ipamPoolId" yaml:"ipamPoolId"`
	// A unique set of Availability Zones, Local Zones, or Wavelength Zones from which AWS advertises IP addresses.
	//
	// Use this parameter to limit the IP address to this location. IP addresses cannot move between network border groups.
	//  Use [DescribeAvailabilityZones](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeAvailabilityZones.html) to view the network border groups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_eip#network_border_group Ec2Eip#network_border_group}
	NetworkBorderGroup *string `field:"optional" json:"networkBorderGroup" yaml:"networkBorderGroup"`
	// The ID of an address pool that you own.
	//
	// Use this parameter to let Amazon EC2 select an address from the address pool.
	//   Updates to the ``PublicIpv4Pool`` property may require *some interruptions*. Updates on an EIP reassociates the address on its associated resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_eip#public_ipv_4_pool Ec2Eip#public_ipv_4_pool}
	PublicIpv4Pool *string `field:"optional" json:"publicIpv4Pool" yaml:"publicIpv4Pool"`
	// Any tags assigned to the Elastic IP address.
	//
	// Updates to the ``Tags`` property may require *some interruptions*. Updates on an EIP reassociates the address on its associated resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_eip#tags Ec2Eip#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The Elastic IP address you are accepting for transfer.
	//
	// You can only accept one transferred address. For more information on Elastic IP address transfers, see [Transfer Elastic IP addresses](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-eips.html#transfer-EIPs-intro) in the *Amazon Virtual Private Cloud User Guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_eip#transfer_address Ec2Eip#transfer_address}
	TransferAddress *string `field:"optional" json:"transferAddress" yaml:"transferAddress"`
}

