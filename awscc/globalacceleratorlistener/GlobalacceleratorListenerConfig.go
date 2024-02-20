package globalacceleratorlistener

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GlobalacceleratorListenerConfig struct {
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
	// The Amazon Resource Name (ARN) of the accelerator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/globalaccelerator_listener#accelerator_arn GlobalacceleratorListener#accelerator_arn}
	AcceleratorArn *string `field:"required" json:"acceleratorArn" yaml:"acceleratorArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/globalaccelerator_listener#port_ranges GlobalacceleratorListener#port_ranges}.
	PortRanges interface{} `field:"required" json:"portRanges" yaml:"portRanges"`
	// Client affinity lets you direct all requests from a user to the same endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/globalaccelerator_listener#client_affinity GlobalacceleratorListener#client_affinity}
	ClientAffinity *string `field:"optional" json:"clientAffinity" yaml:"clientAffinity"`
	// The protocol for the listener.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/globalaccelerator_listener#protocol GlobalacceleratorListener#protocol}
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}

