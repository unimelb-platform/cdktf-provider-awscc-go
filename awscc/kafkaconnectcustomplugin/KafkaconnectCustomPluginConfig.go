package kafkaconnectcustomplugin

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KafkaconnectCustomPluginConfig struct {
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
	// The type of the plugin file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/kafkaconnect_custom_plugin#content_type KafkaconnectCustomPlugin#content_type}
	ContentType *string `field:"required" json:"contentType" yaml:"contentType"`
	// Information about the location of a custom plugin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/kafkaconnect_custom_plugin#location KafkaconnectCustomPlugin#location}
	Location *KafkaconnectCustomPluginLocation `field:"required" json:"location" yaml:"location"`
	// The name of the custom plugin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/kafkaconnect_custom_plugin#name KafkaconnectCustomPlugin#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A summary description of the custom plugin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/kafkaconnect_custom_plugin#description KafkaconnectCustomPlugin#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/kafkaconnect_custom_plugin#tags KafkaconnectCustomPlugin#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

