package ec2subnetnetworkaclassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2SubnetNetworkAclAssociationConfig struct {
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
	// The ID of the network ACL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_subnet_network_acl_association#network_acl_id Ec2SubnetNetworkAclAssociation#network_acl_id}
	NetworkAclId *string `field:"required" json:"networkAclId" yaml:"networkAclId"`
	// The ID of the subnet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_subnet_network_acl_association#subnet_id Ec2SubnetNetworkAclAssociation#subnet_id}
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
}

