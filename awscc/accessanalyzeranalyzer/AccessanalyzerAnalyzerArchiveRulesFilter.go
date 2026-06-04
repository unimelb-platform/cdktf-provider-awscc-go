package accessanalyzeranalyzer


type AccessanalyzerAnalyzerArchiveRulesFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/accessanalyzer_analyzer#contains AccessanalyzerAnalyzer#contains}.
	Contains *[]*string `field:"optional" json:"contains" yaml:"contains"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/accessanalyzer_analyzer#eq AccessanalyzerAnalyzer#eq}.
	Eq *[]*string `field:"optional" json:"eq" yaml:"eq"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/accessanalyzer_analyzer#exists AccessanalyzerAnalyzer#exists}.
	Exists interface{} `field:"optional" json:"exists" yaml:"exists"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/accessanalyzer_analyzer#neq AccessanalyzerAnalyzer#neq}.
	Neq *[]*string `field:"optional" json:"neq" yaml:"neq"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/accessanalyzer_analyzer#property AccessanalyzerAnalyzer#property}.
	Property *string `field:"optional" json:"property" yaml:"property"`
}

