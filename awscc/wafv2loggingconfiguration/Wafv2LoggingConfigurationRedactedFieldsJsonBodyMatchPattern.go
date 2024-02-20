package wafv2loggingconfiguration


type Wafv2LoggingConfigurationRedactedFieldsJsonBodyMatchPattern struct {
	// Match all of the elements.
	//
	// See also MatchScope in JsonBody. You must specify either this setting or the IncludedPaths setting, but not both.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/wafv2_logging_configuration#all Wafv2LoggingConfiguration#all}
	All *string `field:"optional" json:"all" yaml:"all"`
	// Match only the specified include paths. See also MatchScope in JsonBody.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/wafv2_logging_configuration#included_paths Wafv2LoggingConfiguration#included_paths}
	IncludedPaths *[]*string `field:"optional" json:"includedPaths" yaml:"includedPaths"`
}

