package ssmcontactscontactchannel

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SsmcontactsContactChannelConfig struct {
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
	// The details that SSM Incident Manager uses when trying to engage the contact channel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ssmcontacts_contact_channel#channel_address SsmcontactsContactChannel#channel_address}
	ChannelAddress *string `field:"optional" json:"channelAddress" yaml:"channelAddress"`
	// The device name. String of 6 to 50 alphabetical, numeric, dash, and underscore characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ssmcontacts_contact_channel#channel_name SsmcontactsContactChannel#channel_name}
	ChannelName *string `field:"optional" json:"channelName" yaml:"channelName"`
	// Device type, which specify notification channel. Currently supported values: ?SMS?, ?VOICE?, ?EMAIL?, ?CHATBOT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ssmcontacts_contact_channel#channel_type SsmcontactsContactChannel#channel_type}
	ChannelType *string `field:"optional" json:"channelType" yaml:"channelType"`
	// ARN of the contact resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ssmcontacts_contact_channel#contact_id SsmcontactsContactChannel#contact_id}
	ContactId *string `field:"optional" json:"contactId" yaml:"contactId"`
	// If you want to activate the channel at a later time, you can choose to defer activation.
	//
	// SSM Incident Manager can't engage your contact channel until it has been activated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ssmcontacts_contact_channel#defer_activation SsmcontactsContactChannel#defer_activation}
	DeferActivation interface{} `field:"optional" json:"deferActivation" yaml:"deferActivation"`
}

