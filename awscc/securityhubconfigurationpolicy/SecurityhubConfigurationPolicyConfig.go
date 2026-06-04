package securityhubconfigurationpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SecurityhubConfigurationPolicyConfig struct {
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
	// An object that defines how Security Hub is configured.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/securityhub_configuration_policy#configuration_policy SecurityhubConfigurationPolicy#configuration_policy}
	ConfigurationPolicy *SecurityhubConfigurationPolicyConfigurationPolicy `field:"required" json:"configurationPolicy" yaml:"configurationPolicy"`
	// The name of the configuration policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/securityhub_configuration_policy#name SecurityhubConfigurationPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description of the configuration policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/securityhub_configuration_policy#description SecurityhubConfigurationPolicy#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A key-value pair to associate with a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/securityhub_configuration_policy#tags SecurityhubConfigurationPolicy#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

