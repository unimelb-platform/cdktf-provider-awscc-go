package evidentlyexperiment


type EvidentlyExperimentTreatments struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/evidently_experiment#feature EvidentlyExperiment#feature}.
	Feature *string `field:"required" json:"feature" yaml:"feature"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/evidently_experiment#treatment_name EvidentlyExperiment#treatment_name}.
	TreatmentName *string `field:"required" json:"treatmentName" yaml:"treatmentName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/evidently_experiment#variation EvidentlyExperiment#variation}.
	Variation *string `field:"required" json:"variation" yaml:"variation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/evidently_experiment#description EvidentlyExperiment#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

