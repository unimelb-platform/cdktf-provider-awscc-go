package configconformancepack

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConfigConformancePackConfig struct {
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
	// Name of the conformance pack which will be assigned as the unique identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/config_conformance_pack#conformance_pack_name ConfigConformancePack#conformance_pack_name}
	ConformancePackName *string `field:"required" json:"conformancePackName" yaml:"conformancePackName"`
	// A list of ConformancePackInputParameter objects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/config_conformance_pack#conformance_pack_input_parameters ConfigConformancePack#conformance_pack_input_parameters}
	ConformancePackInputParameters interface{} `field:"optional" json:"conformancePackInputParameters" yaml:"conformancePackInputParameters"`
	// AWS Config stores intermediate files while processing conformance pack template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/config_conformance_pack#delivery_s3_bucket ConfigConformancePack#delivery_s3_bucket}
	DeliveryS3Bucket *string `field:"optional" json:"deliveryS3Bucket" yaml:"deliveryS3Bucket"`
	// The prefix for delivery S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/config_conformance_pack#delivery_s3_key_prefix ConfigConformancePack#delivery_s3_key_prefix}
	DeliveryS3KeyPrefix *string `field:"optional" json:"deliveryS3KeyPrefix" yaml:"deliveryS3KeyPrefix"`
	// A string containing full conformance pack template body.
	//
	// You can only specify one of the template body or template S3Uri fields.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/config_conformance_pack#template_body ConfigConformancePack#template_body}
	TemplateBody *string `field:"optional" json:"templateBody" yaml:"templateBody"`
	// Location of file containing the template body which points to the conformance pack template that is located in an Amazon S3 bucket.
	//
	// You can only specify one of the template body or template S3Uri fields.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/config_conformance_pack#template_s3_uri ConfigConformancePack#template_s3_uri}
	TemplateS3Uri *string `field:"optional" json:"templateS3Uri" yaml:"templateS3Uri"`
	// The TemplateSSMDocumentDetails object contains the name of the SSM document and the version of the SSM document.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/config_conformance_pack#template_ssm_document_details ConfigConformancePack#template_ssm_document_details}
	TemplateSsmDocumentDetails *ConfigConformancePackTemplateSsmDocumentDetails `field:"optional" json:"templateSsmDocumentDetails" yaml:"templateSsmDocumentDetails"`
}

