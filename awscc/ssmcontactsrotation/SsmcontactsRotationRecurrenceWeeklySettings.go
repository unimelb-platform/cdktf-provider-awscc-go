package ssmcontactsrotation


type SsmcontactsRotationRecurrenceWeeklySettings struct {
	// The day of the week when weekly recurring on-call shift rotations begin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ssmcontacts_rotation#day_of_week SsmcontactsRotation#day_of_week}
	DayOfWeek *string `field:"required" json:"dayOfWeek" yaml:"dayOfWeek"`
	// Details about when an on-call rotation shift begins or ends. Time of the day in format HH:MM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ssmcontacts_rotation#hand_off_time SsmcontactsRotation#hand_off_time}
	HandOffTime *string `field:"required" json:"handOffTime" yaml:"handOffTime"`
}

