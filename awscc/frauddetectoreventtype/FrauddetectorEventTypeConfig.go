package frauddetectoreventtype

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FrauddetectorEventTypeConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/frauddetector_event_type#entity_types FrauddetectorEventType#entity_types}.
	EntityTypes interface{} `field:"required" json:"entityTypes" yaml:"entityTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/frauddetector_event_type#event_variables FrauddetectorEventType#event_variables}.
	EventVariables interface{} `field:"required" json:"eventVariables" yaml:"eventVariables"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/frauddetector_event_type#labels FrauddetectorEventType#labels}.
	Labels interface{} `field:"required" json:"labels" yaml:"labels"`
	// The name for the event type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/frauddetector_event_type#name FrauddetectorEventType#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description of the event type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/frauddetector_event_type#description FrauddetectorEventType#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Tags associated with this event type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/frauddetector_event_type#tags FrauddetectorEventType#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

