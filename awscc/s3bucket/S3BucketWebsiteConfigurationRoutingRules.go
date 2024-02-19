package s3bucket


type S3BucketWebsiteConfigurationRoutingRules struct {
	// Container for redirect information.
	//
	// You can redirect requests to another host, to another page, or with another protocol. In the event of an error, you can specify a different error code to return.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#redirect_rule S3Bucket#redirect_rule}
	RedirectRule *S3BucketWebsiteConfigurationRoutingRulesRedirectRule `field:"required" json:"redirectRule" yaml:"redirectRule"`
	// A container for describing a condition that must be met for the specified redirect to apply.You must specify at least one of HttpErrorCodeReturnedEquals and KeyPrefixEquals.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#routing_rule_condition S3Bucket#routing_rule_condition}
	RoutingRuleCondition *S3BucketWebsiteConfigurationRoutingRulesRoutingRuleCondition `field:"optional" json:"routingRuleCondition" yaml:"routingRuleCondition"`
}

