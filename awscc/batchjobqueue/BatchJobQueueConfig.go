package batchjobqueue

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BatchJobQueueConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_queue#compute_environment_order BatchJobQueue#compute_environment_order}.
	ComputeEnvironmentOrder interface{} `field:"required" json:"computeEnvironmentOrder" yaml:"computeEnvironmentOrder"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_queue#priority BatchJobQueue#priority}.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_queue#job_queue_name BatchJobQueue#job_queue_name}.
	JobQueueName *string `field:"optional" json:"jobQueueName" yaml:"jobQueueName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_queue#scheduling_policy_arn BatchJobQueue#scheduling_policy_arn}.
	SchedulingPolicyArn *string `field:"optional" json:"schedulingPolicyArn" yaml:"schedulingPolicyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_queue#state BatchJobQueue#state}.
	State *string `field:"optional" json:"state" yaml:"state"`
	// A key-value pair to associate with a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_queue#tags BatchJobQueue#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

