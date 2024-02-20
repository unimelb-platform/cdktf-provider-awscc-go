package batchschedulingpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BatchSchedulingPolicyConfig struct {
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
	// Fair Share Policy for the Job Queue.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_scheduling_policy#fairshare_policy BatchSchedulingPolicy#fairshare_policy}
	FairsharePolicy *BatchSchedulingPolicyFairsharePolicy `field:"optional" json:"fairsharePolicy" yaml:"fairsharePolicy"`
	// Name of Scheduling Policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_scheduling_policy#name BatchSchedulingPolicy#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A key-value pair to associate with a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_scheduling_policy#tags BatchSchedulingPolicy#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

