package personalizesolution


type PersonalizeSolutionSolutionConfigHpoConfigAlgorithmHyperParameterRanges struct {
	// The categorical hyperparameters and their ranges.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/personalize_solution#categorical_hyper_parameter_ranges PersonalizeSolution#categorical_hyper_parameter_ranges}
	CategoricalHyperParameterRanges interface{} `field:"optional" json:"categoricalHyperParameterRanges" yaml:"categoricalHyperParameterRanges"`
	// The continuous hyperparameters and their ranges.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/personalize_solution#continuous_hyper_parameter_ranges PersonalizeSolution#continuous_hyper_parameter_ranges}
	ContinuousHyperParameterRanges interface{} `field:"optional" json:"continuousHyperParameterRanges" yaml:"continuousHyperParameterRanges"`
	// The integer hyperparameters and their ranges.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/personalize_solution#integer_hyper_parameter_ranges PersonalizeSolution#integer_hyper_parameter_ranges}
	IntegerHyperParameterRanges interface{} `field:"optional" json:"integerHyperParameterRanges" yaml:"integerHyperParameterRanges"`
}

