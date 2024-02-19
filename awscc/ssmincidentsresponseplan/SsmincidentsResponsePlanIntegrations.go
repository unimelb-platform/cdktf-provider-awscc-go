package ssmincidentsresponseplan


type SsmincidentsResponsePlanIntegrations struct {
	// The pagerDuty configuration to use when starting the incident.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ssmincidents_response_plan#pager_duty_configuration SsmincidentsResponsePlan#pager_duty_configuration}
	PagerDutyConfiguration *SsmincidentsResponsePlanIntegrationsPagerDutyConfiguration `field:"optional" json:"pagerDutyConfiguration" yaml:"pagerDutyConfiguration"`
}

