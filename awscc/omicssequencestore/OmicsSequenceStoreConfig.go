package omicssequencestore

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OmicsSequenceStoreConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/omics_sequence_store#name OmicsSequenceStore#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A description for the store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/omics_sequence_store#description OmicsSequenceStore#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An S3 URI representing the bucket and folder to store failed read set uploads.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/omics_sequence_store#fallback_location OmicsSequenceStore#fallback_location}
	FallbackLocation *string `field:"optional" json:"fallbackLocation" yaml:"fallbackLocation"`
	// Server-side encryption (SSE) settings for a store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/omics_sequence_store#sse_config OmicsSequenceStore#sse_config}
	SseConfig *OmicsSequenceStoreSseConfig `field:"optional" json:"sseConfig" yaml:"sseConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/omics_sequence_store#tags OmicsSequenceStore#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

