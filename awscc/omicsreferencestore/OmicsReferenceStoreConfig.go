package omicsreferencestore

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OmicsReferenceStoreConfig struct {
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
	// A name for the store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/omics_reference_store#name OmicsReferenceStore#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A description for the store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/omics_reference_store#description OmicsReferenceStore#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Server-side encryption (SSE) settings for a store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/omics_reference_store#sse_config OmicsReferenceStore#sse_config}
	SseConfig *OmicsReferenceStoreSseConfig `field:"optional" json:"sseConfig" yaml:"sseConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/omics_reference_store#tags OmicsReferenceStore#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

