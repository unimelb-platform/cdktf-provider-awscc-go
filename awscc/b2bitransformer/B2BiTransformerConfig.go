package b2bitransformer

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type B2BiTransformerConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/b2bi_transformer#edi_type B2BiTransformer#edi_type}.
	EdiType *B2BiTransformerEdiType `field:"required" json:"ediType" yaml:"ediType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/b2bi_transformer#file_format B2BiTransformer#file_format}.
	FileFormat *string `field:"required" json:"fileFormat" yaml:"fileFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/b2bi_transformer#mapping_template B2BiTransformer#mapping_template}.
	MappingTemplate *string `field:"required" json:"mappingTemplate" yaml:"mappingTemplate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/b2bi_transformer#name B2BiTransformer#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/b2bi_transformer#status B2BiTransformer#status}.
	Status *string `field:"required" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/b2bi_transformer#modified_at B2BiTransformer#modified_at}.
	ModifiedAt *string `field:"optional" json:"modifiedAt" yaml:"modifiedAt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/b2bi_transformer#sample_document B2BiTransformer#sample_document}.
	SampleDocument *string `field:"optional" json:"sampleDocument" yaml:"sampleDocument"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/b2bi_transformer#tags B2BiTransformer#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

