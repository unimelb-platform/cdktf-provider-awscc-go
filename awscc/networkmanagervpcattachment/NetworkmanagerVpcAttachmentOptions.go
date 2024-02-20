package networkmanagervpcattachment


type NetworkmanagerVpcAttachmentOptions struct {
	// Indicates whether to enable ApplianceModeSupport Support for Vpc Attachment. Valid Values: true | false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/networkmanager_vpc_attachment#appliance_mode_support NetworkmanagerVpcAttachment#appliance_mode_support}
	ApplianceModeSupport interface{} `field:"optional" json:"applianceModeSupport" yaml:"applianceModeSupport"`
	// Indicates whether to enable Ipv6 Support for Vpc Attachment. Valid Values: enable | disable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/networkmanager_vpc_attachment#ipv_6_support NetworkmanagerVpcAttachment#ipv_6_support}
	Ipv6Support interface{} `field:"optional" json:"ipv6Support" yaml:"ipv6Support"`
}

