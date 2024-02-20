package accessanalyzeranalyzer


type AccessanalyzerAnalyzerAnalyzerConfigurationUnusedAccessConfiguration struct {
	// The specified access age in days for which to generate findings for unused access.
	//
	// For example, if you specify 90 days, the analyzer will generate findings for IAM entities within the accounts of the selected organization for any access that hasn't been used in 90 or more days since the analyzer's last scan. You can choose a value between 1 and 180 days.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/accessanalyzer_analyzer#unused_access_age AccessanalyzerAnalyzer#unused_access_age}
	UnusedAccessAge *float64 `field:"optional" json:"unusedAccessAge" yaml:"unusedAccessAge"`
}

