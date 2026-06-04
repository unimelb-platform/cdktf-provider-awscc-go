package networkmanagerlink

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkmanagerLinkConfig struct {
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
	// The Bandwidth for the link.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/networkmanager_link#bandwidth NetworkmanagerLink#bandwidth}
	Bandwidth *NetworkmanagerLinkBandwidth `field:"required" json:"bandwidth" yaml:"bandwidth"`
	// The ID of the global network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/networkmanager_link#global_network_id NetworkmanagerLink#global_network_id}
	GlobalNetworkId *string `field:"required" json:"globalNetworkId" yaml:"globalNetworkId"`
	// The ID of the site.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/networkmanager_link#site_id NetworkmanagerLink#site_id}
	SiteId *string `field:"required" json:"siteId" yaml:"siteId"`
	// The description of the link.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/networkmanager_link#description NetworkmanagerLink#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The provider of the link.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/networkmanager_link#provider_name NetworkmanagerLink#provider_name}
	ProviderName *string `field:"optional" json:"providerName" yaml:"providerName"`
	// The tags for the link.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/networkmanager_link#tags NetworkmanagerLink#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The type of the link.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/networkmanager_link#type NetworkmanagerLink#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

