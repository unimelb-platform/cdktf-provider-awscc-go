package eventseventbus

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EventsEventBusConfig struct {
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
	// The name of the event bus.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/events_event_bus#name EventsEventBus#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// If you are creating a partner event bus, this specifies the partner event source that the new event bus will be matched with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/events_event_bus#event_source_name EventsEventBus#event_source_name}
	EventSourceName *string `field:"optional" json:"eventSourceName" yaml:"eventSourceName"`
	// A JSON string that describes the permission policy statement for the event bus.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/events_event_bus#policy EventsEventBus#policy}
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
	// Any tags assigned to the event bus.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/events_event_bus#tags EventsEventBus#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

