package ec2ipamresourcediscovery

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2IpamResourceDiscoveryConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_ipam_resource_discovery#description Ec2IpamResourceDiscovery#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The regions Resource Discovery is enabled for.
	//
	// Allows resource discoveries to be created in these regions, as well as enabling monitoring
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_ipam_resource_discovery#operating_regions Ec2IpamResourceDiscovery#operating_regions}
	OperatingRegions interface{} `field:"optional" json:"operatingRegions" yaml:"operatingRegions"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_ipam_resource_discovery#tags Ec2IpamResourceDiscovery#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

