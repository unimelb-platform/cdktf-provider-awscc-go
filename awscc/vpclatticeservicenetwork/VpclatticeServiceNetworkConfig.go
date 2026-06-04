package vpclatticeservicenetwork

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VpclatticeServiceNetworkConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/vpclattice_service_network#auth_type VpclatticeServiceNetwork#auth_type}.
	AuthType *string `field:"optional" json:"authType" yaml:"authType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/vpclattice_service_network#name VpclatticeServiceNetwork#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/vpclattice_service_network#sharing_config VpclatticeServiceNetwork#sharing_config}.
	SharingConfig *VpclatticeServiceNetworkSharingConfig `field:"optional" json:"sharingConfig" yaml:"sharingConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/vpclattice_service_network#tags VpclatticeServiceNetwork#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

