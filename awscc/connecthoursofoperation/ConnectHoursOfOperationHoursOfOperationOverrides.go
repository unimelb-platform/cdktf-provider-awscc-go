package connecthoursofoperation


type ConnectHoursOfOperationHoursOfOperationOverrides struct {
	// The date from which the hours of operation override would be effective.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/connect_hours_of_operation#effective_from ConnectHoursOfOperation#effective_from}
	EffectiveFrom *string `field:"optional" json:"effectiveFrom" yaml:"effectiveFrom"`
	// The date till which the hours of operation override would be effective.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/connect_hours_of_operation#effective_till ConnectHoursOfOperation#effective_till}
	EffectiveTill *string `field:"optional" json:"effectiveTill" yaml:"effectiveTill"`
	// The Resource Identifier for the hours of operation override.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/connect_hours_of_operation#hours_of_operation_override_id ConnectHoursOfOperation#hours_of_operation_override_id}
	HoursOfOperationOverrideId *string `field:"optional" json:"hoursOfOperationOverrideId" yaml:"hoursOfOperationOverrideId"`
	// Configuration information for the hours of operation override: day, start time, and end time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/connect_hours_of_operation#override_config ConnectHoursOfOperation#override_config}
	OverrideConfig interface{} `field:"optional" json:"overrideConfig" yaml:"overrideConfig"`
	// The description of the hours of operation override.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/connect_hours_of_operation#override_description ConnectHoursOfOperation#override_description}
	OverrideDescription *string `field:"optional" json:"overrideDescription" yaml:"overrideDescription"`
	// The name of the hours of operation override.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/connect_hours_of_operation#override_name ConnectHoursOfOperation#override_name}
	OverrideName *string `field:"optional" json:"overrideName" yaml:"overrideName"`
}

