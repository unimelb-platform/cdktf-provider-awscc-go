package configconfigrule


type ConfigConfigRuleScope struct {
	// ID of the only one resource which we want to trigger the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/config_config_rule#compliance_resource_id ConfigConfigRule#compliance_resource_id}
	ComplianceResourceId *string `field:"optional" json:"complianceResourceId" yaml:"complianceResourceId"`
	// Resource types of resources which we want to trigger the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/config_config_rule#compliance_resource_types ConfigConfigRule#compliance_resource_types}
	ComplianceResourceTypes *[]*string `field:"optional" json:"complianceResourceTypes" yaml:"complianceResourceTypes"`
	// Tag key applied only to resources which we want to trigger the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/config_config_rule#tag_key ConfigConfigRule#tag_key}
	TagKey *string `field:"optional" json:"tagKey" yaml:"tagKey"`
	// Tag value applied only to resources which we want to trigger the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/config_config_rule#tag_value ConfigConfigRule#tag_value}
	TagValue *string `field:"optional" json:"tagValue" yaml:"tagValue"`
}

