package cloudtrailchannel

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudtrailChannelConfig struct {
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
	// One or more resources to which events arriving through a channel are logged and stored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudtrail_channel#destinations CloudtrailChannel#destinations}
	Destinations interface{} `field:"optional" json:"destinations" yaml:"destinations"`
	// The name of the channel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudtrail_channel#name CloudtrailChannel#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The ARN of an on-premises storage solution or application, or a partner event source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudtrail_channel#source CloudtrailChannel#source}
	Source *string `field:"optional" json:"source" yaml:"source"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudtrail_channel#tags CloudtrailChannel#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

