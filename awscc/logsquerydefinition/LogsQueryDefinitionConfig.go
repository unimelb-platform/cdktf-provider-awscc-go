package logsquerydefinition

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LogsQueryDefinitionConfig struct {
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
	// A name for the saved query definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/logs_query_definition#name LogsQueryDefinition#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The query string to use for this definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/logs_query_definition#query_string LogsQueryDefinition#query_string}
	QueryString *string `field:"required" json:"queryString" yaml:"queryString"`
	// Optionally define specific log groups as part of your query definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/logs_query_definition#log_group_names LogsQueryDefinition#log_group_names}
	LogGroupNames *[]*string `field:"optional" json:"logGroupNames" yaml:"logGroupNames"`
}

