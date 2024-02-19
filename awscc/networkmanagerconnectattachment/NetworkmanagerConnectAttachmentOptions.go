package networkmanagerconnectattachment


type NetworkmanagerConnectAttachmentOptions struct {
	// Tunnel protocol for connect attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/networkmanager_connect_attachment#protocol NetworkmanagerConnectAttachment#protocol}
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}

