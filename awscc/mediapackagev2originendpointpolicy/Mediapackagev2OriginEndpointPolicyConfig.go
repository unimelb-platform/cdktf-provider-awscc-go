package mediapackagev2originendpointpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Mediapackagev2OriginEndpointPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackagev2_origin_endpoint_policy#policy Mediapackagev2OriginEndpointPolicy#policy}.
	Policy *string `field:"required" json:"policy" yaml:"policy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackagev2_origin_endpoint_policy#channel_group_name Mediapackagev2OriginEndpointPolicy#channel_group_name}.
	ChannelGroupName *string `field:"optional" json:"channelGroupName" yaml:"channelGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackagev2_origin_endpoint_policy#channel_name Mediapackagev2OriginEndpointPolicy#channel_name}.
	ChannelName *string `field:"optional" json:"channelName" yaml:"channelName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackagev2_origin_endpoint_policy#origin_endpoint_name Mediapackagev2OriginEndpointPolicy#origin_endpoint_name}.
	OriginEndpointName *string `field:"optional" json:"originEndpointName" yaml:"originEndpointName"`
}

