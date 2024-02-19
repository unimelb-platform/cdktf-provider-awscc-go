package ssmcontactsplan


type SsmcontactsPlanStages struct {
	// The time to wait until beginning the next stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ssmcontacts_plan#duration_in_minutes SsmcontactsPlan#duration_in_minutes}
	DurationInMinutes *float64 `field:"required" json:"durationInMinutes" yaml:"durationInMinutes"`
	// The contacts or contact methods that the escalation plan or engagement plan is engaging.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ssmcontacts_plan#targets SsmcontactsPlan#targets}
	Targets interface{} `field:"optional" json:"targets" yaml:"targets"`
}

