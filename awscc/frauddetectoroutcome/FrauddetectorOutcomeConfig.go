package frauddetectoroutcome

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FrauddetectorOutcomeConfig struct {
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
	// The name of the outcome.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/frauddetector_outcome#name FrauddetectorOutcome#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The outcome description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/frauddetector_outcome#description FrauddetectorOutcome#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Tags associated with this outcome.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/frauddetector_outcome#tags FrauddetectorOutcome#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

