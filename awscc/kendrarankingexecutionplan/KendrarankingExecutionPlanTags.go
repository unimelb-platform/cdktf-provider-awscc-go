package kendrarankingexecutionplan


type KendrarankingExecutionPlanTags struct {
	// A string used to identify this tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kendraranking_execution_plan#key KendrarankingExecutionPlan#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// A string containing the value for the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kendraranking_execution_plan#value KendrarankingExecutionPlan#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

