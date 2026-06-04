package medialivesignalmap

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MedialiveSignalMapConfig struct {
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
	// A top-level supported AWS resource ARN to discovery a signal map from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/medialive_signal_map#discovery_entry_point_arn MedialiveSignalMap#discovery_entry_point_arn}
	DiscoveryEntryPointArn *string `field:"required" json:"discoveryEntryPointArn" yaml:"discoveryEntryPointArn"`
	// A resource's name. Names must be unique within the scope of a resource type in a specific region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/medialive_signal_map#name MedialiveSignalMap#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/medialive_signal_map#cloudwatch_alarm_template_group_identifiers MedialiveSignalMap#cloudwatch_alarm_template_group_identifiers}.
	CloudwatchAlarmTemplateGroupIdentifiers *[]*string `field:"optional" json:"cloudwatchAlarmTemplateGroupIdentifiers" yaml:"cloudwatchAlarmTemplateGroupIdentifiers"`
	// A resource's optional description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/medialive_signal_map#description MedialiveSignalMap#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/medialive_signal_map#event_bridge_rule_template_group_identifiers MedialiveSignalMap#event_bridge_rule_template_group_identifiers}.
	EventBridgeRuleTemplateGroupIdentifiers *[]*string `field:"optional" json:"eventBridgeRuleTemplateGroupIdentifiers" yaml:"eventBridgeRuleTemplateGroupIdentifiers"`
	// If true, will force a rediscovery of a signal map if an unchanged discoveryEntryPointArn is provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/medialive_signal_map#force_rediscovery MedialiveSignalMap#force_rediscovery}
	ForceRediscovery interface{} `field:"optional" json:"forceRediscovery" yaml:"forceRediscovery"`
	// Represents the tags associated with a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/medialive_signal_map#tags MedialiveSignalMap#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

