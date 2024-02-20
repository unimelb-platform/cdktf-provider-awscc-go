package ssmparameter

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SsmParameterConfig struct {
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
	// The type of parameter.
	//
	// Although ``SecureString`` is included in the list of valid values, CFNlong does *not* currently support creating a ``SecureString`` parameter type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ssm_parameter#type SsmParameter#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The parameter value.
	//
	// If type is ``StringList``, the system returns a comma-separated string with no spaces between commas in the ``Value`` field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ssm_parameter#value SsmParameter#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// A regular expression used to validate the parameter value.
	//
	// For example, for String types with values restricted to numbers, you can specify the following: ``AllowedPattern=^\d+$``
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ssm_parameter#allowed_pattern SsmParameter#allowed_pattern}
	AllowedPattern *string `field:"optional" json:"allowedPattern" yaml:"allowedPattern"`
	// The data type of the parameter, such as ``text`` or ``aws:ec2:image``. The default is ``text``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ssm_parameter#data_type SsmParameter#data_type}
	DataType *string `field:"optional" json:"dataType" yaml:"dataType"`
	// Information about the parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ssm_parameter#description SsmParameter#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the parameter.
	//
	// The maximum length constraint listed below includes capacity for additional system attributes that aren't part of the name. The maximum length for a parameter name, including the full length of the parameter ARN, is 1011 characters. For example, the length of the following parameter name is 65 characters, not 20 characters: ``arn:aws:ssm:us-east-2:111222333444:parameter/ExampleParameterName``
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ssm_parameter#name SsmParameter#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Information about the policies assigned to a parameter.   [Assigning parameter policies](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-policies.html) in the *User Guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ssm_parameter#policies SsmParameter#policies}
	Policies *string `field:"optional" json:"policies" yaml:"policies"`
	// Optional metadata that you assign to a resource in the form of an arbitrary set of tags (key-value pairs).
	//
	// Tags enable you to categorize a resource in different ways, such as by purpose, owner, or environment. For example, you might want to tag a SYS parameter to identify the type of resource to which it applies, the environment, or the type of configuration data referenced by the parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ssm_parameter#tags SsmParameter#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The parameter tier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ssm_parameter#tier SsmParameter#tier}
	Tier *string `field:"optional" json:"tier" yaml:"tier"`
}

