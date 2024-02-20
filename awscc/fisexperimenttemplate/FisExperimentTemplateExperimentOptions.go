package fisexperimenttemplate


type FisExperimentTemplateExperimentOptions struct {
	// The account targeting setting for the experiment template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/fis_experiment_template#account_targeting FisExperimentTemplate#account_targeting}
	AccountTargeting *string `field:"optional" json:"accountTargeting" yaml:"accountTargeting"`
	// The target resolution failure mode for the experiment template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/fis_experiment_template#empty_target_resolution_mode FisExperimentTemplate#empty_target_resolution_mode}
	EmptyTargetResolutionMode *string `field:"optional" json:"emptyTargetResolutionMode" yaml:"emptyTargetResolutionMode"`
}

