package iottopicruledestination

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IotTopicRuleDestinationConfig struct {
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
	// HTTP URL destination properties.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_topic_rule_destination#http_url_properties IotTopicRuleDestination#http_url_properties}
	HttpUrlProperties *IotTopicRuleDestinationHttpUrlProperties `field:"optional" json:"httpUrlProperties" yaml:"httpUrlProperties"`
	// The status of the TopicRuleDestination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_topic_rule_destination#status IotTopicRuleDestination#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// VPC destination properties.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_topic_rule_destination#vpc_properties IotTopicRuleDestination#vpc_properties}
	VpcProperties *IotTopicRuleDestinationVpcProperties `field:"optional" json:"vpcProperties" yaml:"vpcProperties"`
}

