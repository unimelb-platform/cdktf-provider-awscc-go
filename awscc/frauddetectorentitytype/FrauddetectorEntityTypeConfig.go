package frauddetectorentitytype

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FrauddetectorEntityTypeConfig struct {
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
	// The name of the entity type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/frauddetector_entity_type#name FrauddetectorEntityType#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The entity type description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/frauddetector_entity_type#description FrauddetectorEntityType#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Tags associated with this entity type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/frauddetector_entity_type#tags FrauddetectorEntityType#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

