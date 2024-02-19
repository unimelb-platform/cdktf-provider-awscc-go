package ssmincidentsresponseplan


type SsmincidentsResponsePlanActionsSsmAutomationParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ssmincidents_response_plan#key SsmincidentsResponsePlan#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ssmincidents_response_plan#values SsmincidentsResponsePlan#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

