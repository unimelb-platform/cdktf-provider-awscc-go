package globalacceleratoraccelerator

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GlobalacceleratorAcceleratorConfig struct {
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
	// Name of accelerator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/globalaccelerator_accelerator#name GlobalacceleratorAccelerator#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Indicates whether an accelerator is enabled. The value is true or false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/globalaccelerator_accelerator#enabled GlobalacceleratorAccelerator#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The IP addresses from BYOIP Prefix pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/globalaccelerator_accelerator#ip_addresses GlobalacceleratorAccelerator#ip_addresses}
	IpAddresses *[]*string `field:"optional" json:"ipAddresses" yaml:"ipAddresses"`
	// IP Address type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/globalaccelerator_accelerator#ip_address_type GlobalacceleratorAccelerator#ip_address_type}
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/globalaccelerator_accelerator#tags GlobalacceleratorAccelerator#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

