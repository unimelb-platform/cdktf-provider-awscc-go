package ecsprimarytaskset

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EcsPrimaryTaskSetConfig struct {
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
	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts the service to create the task set in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_primary_task_set#cluster EcsPrimaryTaskSet#cluster}
	Cluster *string `field:"required" json:"cluster" yaml:"cluster"`
	// The short name or full Amazon Resource Name (ARN) of the service to create the task set in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_primary_task_set#service EcsPrimaryTaskSet#service}
	Service *string `field:"required" json:"service" yaml:"service"`
	// The ID or full Amazon Resource Name (ARN) of the task set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_primary_task_set#task_set_id EcsPrimaryTaskSet#task_set_id}
	TaskSetId *string `field:"required" json:"taskSetId" yaml:"taskSetId"`
}

