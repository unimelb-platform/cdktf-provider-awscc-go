package batchjobdefinition

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BatchJobDefinitionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_definition#type BatchJobDefinition#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_definition#container_properties BatchJobDefinition#container_properties}.
	ContainerProperties *BatchJobDefinitionContainerProperties `field:"optional" json:"containerProperties" yaml:"containerProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_definition#eks_properties BatchJobDefinition#eks_properties}.
	EksProperties *BatchJobDefinitionEksProperties `field:"optional" json:"eksProperties" yaml:"eksProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_definition#job_definition_name BatchJobDefinition#job_definition_name}.
	JobDefinitionName *string `field:"optional" json:"jobDefinitionName" yaml:"jobDefinitionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_definition#node_properties BatchJobDefinition#node_properties}.
	NodeProperties *BatchJobDefinitionNodeProperties `field:"optional" json:"nodeProperties" yaml:"nodeProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_definition#parameters BatchJobDefinition#parameters}.
	Parameters *string `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_definition#platform_capabilities BatchJobDefinition#platform_capabilities}.
	PlatformCapabilities *[]*string `field:"optional" json:"platformCapabilities" yaml:"platformCapabilities"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_definition#propagate_tags BatchJobDefinition#propagate_tags}.
	PropagateTags interface{} `field:"optional" json:"propagateTags" yaml:"propagateTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_definition#retry_strategy BatchJobDefinition#retry_strategy}.
	RetryStrategy *BatchJobDefinitionRetryStrategy `field:"optional" json:"retryStrategy" yaml:"retryStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_definition#scheduling_priority BatchJobDefinition#scheduling_priority}.
	SchedulingPriority *float64 `field:"optional" json:"schedulingPriority" yaml:"schedulingPriority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_definition#tags BatchJobDefinition#tags}.
	Tags *string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_definition#timeout BatchJobDefinition#timeout}.
	Timeout *BatchJobDefinitionTimeout `field:"optional" json:"timeout" yaml:"timeout"`
}

