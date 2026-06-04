package medialivesdisource

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MedialiveSdiSourceConfig struct {
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
	// The name of the SdiSource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/medialive_sdi_source#name MedialiveSdiSource#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The interface mode of the SdiSource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/medialive_sdi_source#type MedialiveSdiSource#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The current state of the SdiSource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/medialive_sdi_source#mode MedialiveSdiSource#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// A collection of key-value pairs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/medialive_sdi_source#tags MedialiveSdiSource#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

