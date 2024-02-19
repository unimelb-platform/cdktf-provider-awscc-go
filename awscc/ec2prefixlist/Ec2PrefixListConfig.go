package ec2prefixlist

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2PrefixListConfig struct {
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
	// Ip Version of Prefix List.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_prefix_list#address_family Ec2PrefixList#address_family}
	AddressFamily *string `field:"required" json:"addressFamily" yaml:"addressFamily"`
	// Max Entries of Prefix List.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_prefix_list#max_entries Ec2PrefixList#max_entries}
	MaxEntries *float64 `field:"required" json:"maxEntries" yaml:"maxEntries"`
	// Name of Prefix List.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_prefix_list#prefix_list_name Ec2PrefixList#prefix_list_name}
	PrefixListName *string `field:"required" json:"prefixListName" yaml:"prefixListName"`
	// Entries of Prefix List.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_prefix_list#entries Ec2PrefixList#entries}
	Entries interface{} `field:"optional" json:"entries" yaml:"entries"`
	// Tags for Prefix List.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_prefix_list#tags Ec2PrefixList#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

