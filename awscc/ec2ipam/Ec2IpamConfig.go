package ec2ipam

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2IpamConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_ipam#description Ec2Ipam#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The regions IPAM is enabled for. Allows pools to be created in these regions, as well as enabling monitoring.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_ipam#operating_regions Ec2Ipam#operating_regions}
	OperatingRegions interface{} `field:"optional" json:"operatingRegions" yaml:"operatingRegions"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_ipam#tags Ec2Ipam#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The tier of the IPAM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_ipam#tier Ec2Ipam#tier}
	Tier *string `field:"optional" json:"tier" yaml:"tier"`
}

