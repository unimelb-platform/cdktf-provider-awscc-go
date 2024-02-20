package applicationinsightsapplication


type ApplicationinsightsApplicationLogPatternSetsLogPatterns struct {
	// The log pattern.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationinsights_application#pattern ApplicationinsightsApplication#pattern}
	Pattern *string `field:"required" json:"pattern" yaml:"pattern"`
	// The name of the log pattern.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationinsights_application#pattern_name ApplicationinsightsApplication#pattern_name}
	PatternName *string `field:"required" json:"patternName" yaml:"patternName"`
	// Rank of the log pattern.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationinsights_application#rank ApplicationinsightsApplication#rank}
	Rank *float64 `field:"required" json:"rank" yaml:"rank"`
}

