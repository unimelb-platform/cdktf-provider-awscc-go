package iotanalyticspipeline

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IotanalyticsPipelineConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iotanalytics_pipeline#pipeline_activities IotanalyticsPipeline#pipeline_activities}.
	PipelineActivities interface{} `field:"required" json:"pipelineActivities" yaml:"pipelineActivities"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iotanalytics_pipeline#pipeline_name IotanalyticsPipeline#pipeline_name}.
	PipelineName *string `field:"optional" json:"pipelineName" yaml:"pipelineName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iotanalytics_pipeline#tags IotanalyticsPipeline#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

