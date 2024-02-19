package frauddetectorlist

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FrauddetectorListConfig struct {
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
	// The name of the list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/frauddetector_list#name FrauddetectorList#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description of the list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/frauddetector_list#description FrauddetectorList#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The elements in this list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/frauddetector_list#elements FrauddetectorList#elements}
	Elements *[]*string `field:"optional" json:"elements" yaml:"elements"`
	// Tags associated with this list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/frauddetector_list#tags FrauddetectorList#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The variable type of the list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/frauddetector_list#variable_type FrauddetectorList#variable_type}
	VariableType *string `field:"optional" json:"variableType" yaml:"variableType"`
}

