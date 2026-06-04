package evidentlyproject


type EvidentlyProjectAppConfigResource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/evidently_project#application_id EvidentlyProject#application_id}.
	ApplicationId *string `field:"optional" json:"applicationId" yaml:"applicationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/evidently_project#environment_id EvidentlyProject#environment_id}.
	EnvironmentId *string `field:"optional" json:"environmentId" yaml:"environmentId"`
}

