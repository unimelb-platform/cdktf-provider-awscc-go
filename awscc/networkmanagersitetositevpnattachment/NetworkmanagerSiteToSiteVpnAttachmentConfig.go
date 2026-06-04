package networkmanagersitetositevpnattachment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkmanagerSiteToSiteVpnAttachmentConfig struct {
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
	// The ID of a core network where you're creating a site-to-site VPN attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/networkmanager_site_to_site_vpn_attachment#core_network_id NetworkmanagerSiteToSiteVpnAttachment#core_network_id}
	CoreNetworkId *string `field:"required" json:"coreNetworkId" yaml:"coreNetworkId"`
	// The ARN of the site-to-site VPN attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/networkmanager_site_to_site_vpn_attachment#vpn_connection_arn NetworkmanagerSiteToSiteVpnAttachment#vpn_connection_arn}
	VpnConnectionArn *string `field:"required" json:"vpnConnectionArn" yaml:"vpnConnectionArn"`
	// The name of the network function group attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/networkmanager_site_to_site_vpn_attachment#network_function_group_name NetworkmanagerSiteToSiteVpnAttachment#network_function_group_name}
	NetworkFunctionGroupName *string `field:"optional" json:"networkFunctionGroupName" yaml:"networkFunctionGroupName"`
	// The attachment to move from one network function group to another.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/networkmanager_site_to_site_vpn_attachment#proposed_network_function_group_change NetworkmanagerSiteToSiteVpnAttachment#proposed_network_function_group_change}
	ProposedNetworkFunctionGroupChange *NetworkmanagerSiteToSiteVpnAttachmentProposedNetworkFunctionGroupChange `field:"optional" json:"proposedNetworkFunctionGroupChange" yaml:"proposedNetworkFunctionGroupChange"`
	// The attachment to move from one segment to another.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/networkmanager_site_to_site_vpn_attachment#proposed_segment_change NetworkmanagerSiteToSiteVpnAttachment#proposed_segment_change}
	ProposedSegmentChange *NetworkmanagerSiteToSiteVpnAttachmentProposedSegmentChange `field:"optional" json:"proposedSegmentChange" yaml:"proposedSegmentChange"`
	// Tags for the attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/networkmanager_site_to_site_vpn_attachment#tags NetworkmanagerSiteToSiteVpnAttachment#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

