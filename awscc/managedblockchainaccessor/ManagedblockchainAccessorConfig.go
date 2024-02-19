package managedblockchainaccessor

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ManagedblockchainAccessorConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/managedblockchain_accessor#accessor_type ManagedblockchainAccessor#accessor_type}.
	AccessorType *string `field:"required" json:"accessorType" yaml:"accessorType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/managedblockchain_accessor#network_type ManagedblockchainAccessor#network_type}.
	NetworkType *string `field:"optional" json:"networkType" yaml:"networkType"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/managedblockchain_accessor#tags ManagedblockchainAccessor#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

