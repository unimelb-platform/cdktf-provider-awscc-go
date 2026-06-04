package applicationsignalsservicelevelobjective


type ApplicationsignalsServiceLevelObjectiveExclusionWindowsRecurrenceRule struct {
	// A cron or rate expression denoting how often to repeat this exclusion window.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/applicationsignals_service_level_objective#expression ApplicationsignalsServiceLevelObjective#expression}
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
}

