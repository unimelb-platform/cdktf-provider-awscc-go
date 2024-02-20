package networkmanagervpcattachment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkmanagerVpcAttachmentConfig struct {
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
	// The ID of a core network for the VPC attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/networkmanager_vpc_attachment#core_network_id NetworkmanagerVpcAttachment#core_network_id}
	CoreNetworkId *string `field:"required" json:"coreNetworkId" yaml:"coreNetworkId"`
	// Subnet Arn list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/networkmanager_vpc_attachment#subnet_arns NetworkmanagerVpcAttachment#subnet_arns}
	SubnetArns *[]*string `field:"required" json:"subnetArns" yaml:"subnetArns"`
	// The ARN of the VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/networkmanager_vpc_attachment#vpc_arn NetworkmanagerVpcAttachment#vpc_arn}
	VpcArn *string `field:"required" json:"vpcArn" yaml:"vpcArn"`
	// Vpc options of the attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/networkmanager_vpc_attachment#options NetworkmanagerVpcAttachment#options}
	Options *NetworkmanagerVpcAttachmentOptions `field:"optional" json:"options" yaml:"options"`
	// The attachment to move from one segment to another.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/networkmanager_vpc_attachment#proposed_segment_change NetworkmanagerVpcAttachment#proposed_segment_change}
	ProposedSegmentChange *NetworkmanagerVpcAttachmentProposedSegmentChange `field:"optional" json:"proposedSegmentChange" yaml:"proposedSegmentChange"`
	// Tags for the attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/networkmanager_vpc_attachment#tags NetworkmanagerVpcAttachment#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

