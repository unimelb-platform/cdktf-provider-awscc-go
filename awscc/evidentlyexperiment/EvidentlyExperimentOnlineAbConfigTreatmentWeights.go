package evidentlyexperiment


type EvidentlyExperimentOnlineAbConfigTreatmentWeights struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/evidently_experiment#split_weight EvidentlyExperiment#split_weight}.
	SplitWeight *float64 `field:"required" json:"splitWeight" yaml:"splitWeight"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/evidently_experiment#treatment EvidentlyExperiment#treatment}.
	Treatment *string `field:"required" json:"treatment" yaml:"treatment"`
}

