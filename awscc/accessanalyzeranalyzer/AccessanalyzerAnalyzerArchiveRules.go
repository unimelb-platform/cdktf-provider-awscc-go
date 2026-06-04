package accessanalyzeranalyzer


type AccessanalyzerAnalyzerArchiveRules struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/accessanalyzer_analyzer#filter AccessanalyzerAnalyzer#filter}.
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
	// The archive rule name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/accessanalyzer_analyzer#rule_name AccessanalyzerAnalyzer#rule_name}
	RuleName *string `field:"optional" json:"ruleName" yaml:"ruleName"`
}

