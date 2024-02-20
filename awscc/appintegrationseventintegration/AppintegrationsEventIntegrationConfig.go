package appintegrationseventintegration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppintegrationsEventIntegrationConfig struct {
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
	// The Amazon Eventbridge bus for the event integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appintegrations_event_integration#event_bridge_bus AppintegrationsEventIntegration#event_bridge_bus}
	EventBridgeBus *string `field:"required" json:"eventBridgeBus" yaml:"eventBridgeBus"`
	// The EventFilter (source) associated with the event integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appintegrations_event_integration#event_filter AppintegrationsEventIntegration#event_filter}
	EventFilter *AppintegrationsEventIntegrationEventFilter `field:"required" json:"eventFilter" yaml:"eventFilter"`
	// The name of the event integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appintegrations_event_integration#name AppintegrationsEventIntegration#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The event integration description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appintegrations_event_integration#description AppintegrationsEventIntegration#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The tags (keys and values) associated with the event integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appintegrations_event_integration#tags AppintegrationsEventIntegration#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

