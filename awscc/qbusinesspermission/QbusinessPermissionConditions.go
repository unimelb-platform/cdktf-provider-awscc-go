package qbusinesspermission


type QbusinessPermissionConditions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/qbusiness_permission#condition_key QbusinessPermission#condition_key}.
	ConditionKey *string `field:"optional" json:"conditionKey" yaml:"conditionKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/qbusiness_permission#condition_operator QbusinessPermission#condition_operator}.
	ConditionOperator *string `field:"optional" json:"conditionOperator" yaml:"conditionOperator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/qbusiness_permission#condition_values QbusinessPermission#condition_values}.
	ConditionValues *[]*string `field:"optional" json:"conditionValues" yaml:"conditionValues"`
}

