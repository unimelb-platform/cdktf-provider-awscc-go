package cleanroomsmembership


type CleanroomsMembershipDefaultResultConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cleanrooms_membership#output_configuration CleanroomsMembership#output_configuration}.
	OutputConfiguration *CleanroomsMembershipDefaultResultConfigurationOutputConfiguration `field:"required" json:"outputConfiguration" yaml:"outputConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cleanrooms_membership#role_arn CleanroomsMembership#role_arn}.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

