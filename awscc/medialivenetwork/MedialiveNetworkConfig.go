package medialivenetwork

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MedialiveNetworkConfig struct {
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
	// The list of IP address cidr pools for the network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/medialive_network#ip_pools MedialiveNetwork#ip_pools}
	IpPools interface{} `field:"required" json:"ipPools" yaml:"ipPools"`
	// The user-specified name of the Network to be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/medialive_network#name MedialiveNetwork#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The routes for the network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/medialive_network#routes MedialiveNetwork#routes}
	Routes interface{} `field:"optional" json:"routes" yaml:"routes"`
	// A collection of key-value pairs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/medialive_network#tags MedialiveNetwork#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

