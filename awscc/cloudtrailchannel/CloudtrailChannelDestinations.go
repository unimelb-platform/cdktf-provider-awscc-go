package cloudtrailchannel


type CloudtrailChannelDestinations struct {
	// The ARN of a resource that receives events from a channel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudtrail_channel#location CloudtrailChannel#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The type of destination for events arriving from a channel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudtrail_channel#type CloudtrailChannel#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

