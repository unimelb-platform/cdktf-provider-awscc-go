package ssmincidentsresponseplan


type SsmincidentsResponsePlanActions struct {
	// The configuration to use when starting the SSM automation document.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ssmincidents_response_plan#ssm_automation SsmincidentsResponsePlan#ssm_automation}
	SsmAutomation *SsmincidentsResponsePlanActionsSsmAutomation `field:"optional" json:"ssmAutomation" yaml:"ssmAutomation"`
}

