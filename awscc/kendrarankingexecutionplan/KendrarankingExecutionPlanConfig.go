package kendrarankingexecutionplan

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KendrarankingExecutionPlanConfig struct {
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
	// Name of kendra ranking rescore execution plan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kendraranking_execution_plan#name KendrarankingExecutionPlan#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Capacity units.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kendraranking_execution_plan#capacity_units KendrarankingExecutionPlan#capacity_units}
	CapacityUnits *KendrarankingExecutionPlanCapacityUnits `field:"optional" json:"capacityUnits" yaml:"capacityUnits"`
	// A description for the execution plan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kendraranking_execution_plan#description KendrarankingExecutionPlan#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Tags for labeling the execution plan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kendraranking_execution_plan#tags KendrarankingExecutionPlan#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

