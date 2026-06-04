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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/b2bi_transformer#name B2BiTransformer#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/b2bi_transformer#status B2BiTransformer#status}.
	Status *string `field:"required" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/b2bi_transformer#edi_type B2BiTransformer#edi_type}.
	EdiType *B2BiTransformerEdiType `field:"optional" json:"ediType" yaml:"ediType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/b2bi_transformer#file_format B2BiTransformer#file_format}.
	FileFormat *string `field:"optional" json:"fileFormat" yaml:"fileFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/b2bi_transformer#input_conversion B2BiTransformer#input_conversion}.
	InputConversion *B2BiTransformerInputConversion `field:"optional" json:"inputConversion" yaml:"inputConversion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/b2bi_transformer#mapping B2BiTransformer#mapping}.
	Mapping *B2BiTransformerMapping `field:"optional" json:"mapping" yaml:"mapping"`
	// This shape is deprecated: This is a legacy trait. Please use input-conversion or output-conversion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/b2bi_transformer#mapping_template B2BiTransformer#mapping_template}
	MappingTemplate *string `field:"optional" json:"mappingTemplate" yaml:"mappingTemplate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/b2bi_transformer#output_conversion B2BiTransformer#output_conversion}.
	OutputConversion *B2BiTransformerOutputConversion `field:"optional" json:"outputConversion" yaml:"outputConversion"`
	// This shape is deprecated: This is a legacy trait. Please use input-conversion or output-conversion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/b2bi_transformer#sample_document B2BiTransformer#sample_document}
	SampleDocument *string `field:"optional" json:"sampleDocument" yaml:"sampleDocument"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/b2bi_transformer#sample_documents B2BiTransformer#sample_documents}.
	SampleDocuments *B2BiTransformerSampleDocuments `field:"optional" json:"sampleDocuments" yaml:"sampleDocuments"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/b2bi_transformer#tags B2BiTransformer#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

