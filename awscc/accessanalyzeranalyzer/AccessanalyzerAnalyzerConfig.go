package accessanalyzeranalyzer

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AccessanalyzerAnalyzerConfig struct {
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
	// The type of the analyzer, must be one of ACCOUNT, ORGANIZATION, ACCOUNT_UNUSED_ACCESS or ORGANIZATION_UNUSED_ACCESS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/accessanalyzer_analyzer#type AccessanalyzerAnalyzer#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The configuration for the analyzer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/accessanalyzer_analyzer#analyzer_configuration AccessanalyzerAnalyzer#analyzer_configuration}
	AnalyzerConfiguration *AccessanalyzerAnalyzerAnalyzerConfiguration `field:"optional" json:"analyzerConfiguration" yaml:"analyzerConfiguration"`
	// Analyzer name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/accessanalyzer_analyzer#analyzer_name AccessanalyzerAnalyzer#analyzer_name}
	AnalyzerName *string `field:"optional" json:"analyzerName" yaml:"analyzerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/accessanalyzer_analyzer#archive_rules AccessanalyzerAnalyzer#archive_rules}.
	ArchiveRules interface{} `field:"optional" json:"archiveRules" yaml:"archiveRules"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/accessanalyzer_analyzer#tags AccessanalyzerAnalyzer#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

