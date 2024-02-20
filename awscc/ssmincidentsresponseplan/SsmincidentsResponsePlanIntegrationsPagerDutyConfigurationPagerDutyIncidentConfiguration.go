package ssmincidentsresponseplan


type SsmincidentsResponsePlanIntegrationsPagerDutyConfigurationPagerDutyIncidentConfiguration struct {
	// The pagerDuty serviceId.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ssmincidents_response_plan#service_id SsmincidentsResponsePlan#service_id}
	ServiceId *string `field:"required" json:"serviceId" yaml:"serviceId"`
}

