package connecthoursofoperation


type ConnectHoursOfOperationConfigEndTime struct {
	// The hours.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_hours_of_operation#hours ConnectHoursOfOperation#hours}
	Hours *float64 `field:"required" json:"hours" yaml:"hours"`
	// The minutes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_hours_of_operation#minutes ConnectHoursOfOperation#minutes}
	Minutes *float64 `field:"required" json:"minutes" yaml:"minutes"`
}

