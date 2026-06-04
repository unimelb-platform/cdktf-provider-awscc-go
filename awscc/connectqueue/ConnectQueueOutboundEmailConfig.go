package connectqueue


type ConnectQueueOutboundEmailConfig struct {
	// The email address connect resource ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/connect_queue#outbound_email_address_id ConnectQueue#outbound_email_address_id}
	OutboundEmailAddressId *string `field:"optional" json:"outboundEmailAddressId" yaml:"outboundEmailAddressId"`
}

