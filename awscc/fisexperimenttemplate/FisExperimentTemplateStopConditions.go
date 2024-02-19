package fisexperimenttemplate


type FisExperimentTemplateStopConditions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/fis_experiment_template#source FisExperimentTemplate#source}.
	Source *string `field:"required" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/fis_experiment_template#value FisExperimentTemplate#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

