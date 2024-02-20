package personalizesolution


type PersonalizeSolutionSolutionConfigAutoMlConfig struct {
	// The metric to optimize.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/personalize_solution#metric_name PersonalizeSolution#metric_name}
	MetricName *string `field:"optional" json:"metricName" yaml:"metricName"`
	// The list of candidate recipes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/personalize_solution#recipe_list PersonalizeSolution#recipe_list}
	RecipeList *[]*string `field:"optional" json:"recipeList" yaml:"recipeList"`
}

