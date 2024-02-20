package accessanalyzeranalyzer


type AccessanalyzerAnalyzerAnalyzerConfiguration struct {
	// The Configuration for Unused Access Analyzer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/accessanalyzer_analyzer#unused_access_configuration AccessanalyzerAnalyzer#unused_access_configuration}
	UnusedAccessConfiguration *AccessanalyzerAnalyzerAnalyzerConfigurationUnusedAccessConfiguration `field:"optional" json:"unusedAccessConfiguration" yaml:"unusedAccessConfiguration"`
}

