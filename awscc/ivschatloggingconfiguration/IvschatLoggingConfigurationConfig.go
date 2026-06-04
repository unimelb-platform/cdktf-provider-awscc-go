package ivschatloggingconfiguration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IvschatLoggingConfigurationConfig struct {
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
	// Destination configuration for IVS Chat logging.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ivschat_logging_configuration#destination_configuration IvschatLoggingConfiguration#destination_configuration}
	DestinationConfiguration *IvschatLoggingConfigurationDestinationConfiguration `field:"required" json:"destinationConfiguration" yaml:"destinationConfiguration"`
	// The name of the logging configuration. The value does not need to be unique.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ivschat_logging_configuration#name IvschatLoggingConfiguration#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ivschat_logging_configuration#tags IvschatLoggingConfiguration#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

