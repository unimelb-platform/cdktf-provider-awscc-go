package mediaconnectflow

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MediaconnectFlowConfig struct {
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
	// The name of the flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_flow#name MediaconnectFlow#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The source of the flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_flow#source MediaconnectFlow#source}
	Source *MediaconnectFlowSource `field:"required" json:"source" yaml:"source"`
	// The Availability Zone that you want to create the flow in.
	//
	// These options are limited to the Availability Zones within the current AWS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_flow#availability_zone MediaconnectFlow#availability_zone}
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The source failover config of the flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_flow#source_failover_config MediaconnectFlow#source_failover_config}
	SourceFailoverConfig *MediaconnectFlowSourceFailoverConfig `field:"optional" json:"sourceFailoverConfig" yaml:"sourceFailoverConfig"`
}

