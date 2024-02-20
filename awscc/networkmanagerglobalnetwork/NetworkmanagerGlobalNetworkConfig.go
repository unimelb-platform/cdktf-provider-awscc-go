package networkmanagerglobalnetwork

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkmanagerGlobalNetworkConfig struct {
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
	// The date and time that the global network was created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/networkmanager_global_network#created_at NetworkmanagerGlobalNetwork#created_at}
	CreatedAt *string `field:"optional" json:"createdAt" yaml:"createdAt"`
	// The description of the global network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/networkmanager_global_network#description NetworkmanagerGlobalNetwork#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The state of the global network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/networkmanager_global_network#state NetworkmanagerGlobalNetwork#state}
	State *string `field:"optional" json:"state" yaml:"state"`
	// The tags for the global network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/networkmanager_global_network#tags NetworkmanagerGlobalNetwork#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

