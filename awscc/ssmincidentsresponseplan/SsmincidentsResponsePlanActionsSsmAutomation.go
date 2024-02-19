package ssmincidentsresponseplan


type SsmincidentsResponsePlanActionsSsmAutomation struct {
	// The document name to use when starting the SSM automation document.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ssmincidents_response_plan#document_name SsmincidentsResponsePlan#document_name}
	DocumentName *string `field:"required" json:"documentName" yaml:"documentName"`
	// The role ARN to use when starting the SSM automation document.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ssmincidents_response_plan#role_arn SsmincidentsResponsePlan#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The version of the document to use when starting the SSM automation document.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ssmincidents_response_plan#document_version SsmincidentsResponsePlan#document_version}
	DocumentVersion *string `field:"optional" json:"documentVersion" yaml:"documentVersion"`
	// The parameters with dynamic values to set when starting the SSM automation document.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ssmincidents_response_plan#dynamic_parameters SsmincidentsResponsePlan#dynamic_parameters}
	DynamicParameters interface{} `field:"optional" json:"dynamicParameters" yaml:"dynamicParameters"`
	// The parameters to set when starting the SSM automation document.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ssmincidents_response_plan#parameters SsmincidentsResponsePlan#parameters}
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// The account type to use when starting the SSM automation document.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ssmincidents_response_plan#target_account SsmincidentsResponsePlan#target_account}
	TargetAccount *string `field:"optional" json:"targetAccount" yaml:"targetAccount"`
}

