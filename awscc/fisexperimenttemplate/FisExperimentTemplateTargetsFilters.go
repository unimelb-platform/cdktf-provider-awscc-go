package fisexperimenttemplate


type FisExperimentTemplateTargetsFilters struct {
	// The attribute path for the filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/fis_experiment_template#path FisExperimentTemplate#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// The attribute values for the filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/fis_experiment_template#values FisExperimentTemplate#values}
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

