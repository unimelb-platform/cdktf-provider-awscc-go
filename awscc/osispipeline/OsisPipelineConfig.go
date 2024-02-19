package osispipeline

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OsisPipelineConfig struct {
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
	// The maximum pipeline capacity, in Ingestion Compute Units (ICUs).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/osis_pipeline#max_units OsisPipeline#max_units}
	MaxUnits *float64 `field:"required" json:"maxUnits" yaml:"maxUnits"`
	// The minimum pipeline capacity, in Ingestion Compute Units (ICUs).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/osis_pipeline#min_units OsisPipeline#min_units}
	MinUnits *float64 `field:"required" json:"minUnits" yaml:"minUnits"`
	// The Data Prepper pipeline configuration in YAML format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/osis_pipeline#pipeline_configuration_body OsisPipeline#pipeline_configuration_body}
	PipelineConfigurationBody *string `field:"required" json:"pipelineConfigurationBody" yaml:"pipelineConfigurationBody"`
	// Name of the OpenSearch Ingestion Service pipeline to create.
	//
	// Pipeline names are unique across the pipelines owned by an account within an AWS Region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/osis_pipeline#pipeline_name OsisPipeline#pipeline_name}
	PipelineName *string `field:"required" json:"pipelineName" yaml:"pipelineName"`
	// Key-value pairs to configure buffering.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/osis_pipeline#buffer_options OsisPipeline#buffer_options}
	BufferOptions *OsisPipelineBufferOptions `field:"optional" json:"bufferOptions" yaml:"bufferOptions"`
	// Key-value pairs to configure encryption at rest.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/osis_pipeline#encryption_at_rest_options OsisPipeline#encryption_at_rest_options}
	EncryptionAtRestOptions *OsisPipelineEncryptionAtRestOptions `field:"optional" json:"encryptionAtRestOptions" yaml:"encryptionAtRestOptions"`
	// Key-value pairs to configure log publishing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/osis_pipeline#log_publishing_options OsisPipeline#log_publishing_options}
	LogPublishingOptions *OsisPipelineLogPublishingOptions `field:"optional" json:"logPublishingOptions" yaml:"logPublishingOptions"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/osis_pipeline#tags OsisPipeline#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Container for the values required to configure VPC access for the pipeline.
	//
	// If you don't specify these values, OpenSearch Ingestion Service creates the pipeline with a public endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/osis_pipeline#vpc_options OsisPipeline#vpc_options}
	VpcOptions *OsisPipelineVpcOptions `field:"optional" json:"vpcOptions" yaml:"vpcOptions"`
}

