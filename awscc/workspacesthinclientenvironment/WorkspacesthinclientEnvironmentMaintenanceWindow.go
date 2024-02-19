package workspacesthinclientenvironment


type WorkspacesthinclientEnvironmentMaintenanceWindow struct {
	// The type of maintenance window.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/workspacesthinclient_environment#type WorkspacesthinclientEnvironment#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The desired time zone maintenance window.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/workspacesthinclient_environment#apply_time_of WorkspacesthinclientEnvironment#apply_time_of}
	ApplyTimeOf *string `field:"optional" json:"applyTimeOf" yaml:"applyTimeOf"`
	// The date of maintenance window.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/workspacesthinclient_environment#days_of_the_week WorkspacesthinclientEnvironment#days_of_the_week}
	DaysOfTheWeek *[]*string `field:"optional" json:"daysOfTheWeek" yaml:"daysOfTheWeek"`
	// The hour end time of maintenance window.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/workspacesthinclient_environment#end_time_hour WorkspacesthinclientEnvironment#end_time_hour}
	EndTimeHour *float64 `field:"optional" json:"endTimeHour" yaml:"endTimeHour"`
	// The minute end time of maintenance window.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/workspacesthinclient_environment#end_time_minute WorkspacesthinclientEnvironment#end_time_minute}
	EndTimeMinute *float64 `field:"optional" json:"endTimeMinute" yaml:"endTimeMinute"`
	// The hour start time of maintenance window.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/workspacesthinclient_environment#start_time_hour WorkspacesthinclientEnvironment#start_time_hour}
	StartTimeHour *float64 `field:"optional" json:"startTimeHour" yaml:"startTimeHour"`
	// The minute start time of maintenance window.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/workspacesthinclient_environment#start_time_minute WorkspacesthinclientEnvironment#start_time_minute}
	StartTimeMinute *float64 `field:"optional" json:"startTimeMinute" yaml:"startTimeMinute"`
}

