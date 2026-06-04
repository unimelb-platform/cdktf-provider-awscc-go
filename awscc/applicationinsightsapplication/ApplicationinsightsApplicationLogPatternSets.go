package applicationinsightsapplication


type ApplicationinsightsApplicationLogPatternSets struct {
	// The log patterns of a set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/applicationinsights_application#log_patterns ApplicationinsightsApplication#log_patterns}
	LogPatterns interface{} `field:"optional" json:"logPatterns" yaml:"logPatterns"`
	// The name of the log pattern set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/applicationinsights_application#pattern_set_name ApplicationinsightsApplication#pattern_set_name}
	PatternSetName *string `field:"optional" json:"patternSetName" yaml:"patternSetName"`
}

