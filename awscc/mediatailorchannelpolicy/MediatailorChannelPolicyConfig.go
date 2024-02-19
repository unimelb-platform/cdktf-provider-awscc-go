package mediatailorchannelpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MediatailorChannelPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediatailor_channel_policy#channel_name MediatailorChannelPolicy#channel_name}.
	ChannelName *string `field:"required" json:"channelName" yaml:"channelName"`
	// <p>The IAM policy for the channel. IAM policies are used to control access to your channel.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediatailor_channel_policy#policy MediatailorChannelPolicy#policy}
	Policy *string `field:"required" json:"policy" yaml:"policy"`
}

