package wafv2loggingconfiguration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Wafv2LoggingConfigurationConfig struct {
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
	// The Amazon Resource Names (ARNs) of the logging destinations that you want to associate with the web ACL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/wafv2_logging_configuration#log_destination_configs Wafv2LoggingConfiguration#log_destination_configs}
	LogDestinationConfigs *[]*string `field:"required" json:"logDestinationConfigs" yaml:"logDestinationConfigs"`
	// The Amazon Resource Name (ARN) of the web ACL that you want to associate with LogDestinationConfigs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/wafv2_logging_configuration#resource_arn Wafv2LoggingConfiguration#resource_arn}
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
	// Filtering that specifies which web requests are kept in the logs and which are dropped.
	//
	// You can filter on the rule action and on the web request labels that were applied by matching rules during web ACL evaluation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/wafv2_logging_configuration#logging_filter Wafv2LoggingConfiguration#logging_filter}
	LoggingFilter *Wafv2LoggingConfigurationLoggingFilter `field:"optional" json:"loggingFilter" yaml:"loggingFilter"`
	// The parts of the request that you want to keep out of the logs.
	//
	// For example, if you redact the HEADER field, the HEADER field in the firehose will be xxx.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/wafv2_logging_configuration#redacted_fields Wafv2LoggingConfiguration#redacted_fields}
	RedactedFields interface{} `field:"optional" json:"redactedFields" yaml:"redactedFields"`
}

